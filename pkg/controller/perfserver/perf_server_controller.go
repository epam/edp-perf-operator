package perfserver

import (
	"context"
	"github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/v2/pkg/client/perf"
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfserver/chain"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"time"
)

func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcilePerfServer{
		client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
	}
}

func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("perfserver-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	p := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			oldObject := e.ObjectOld.(*v1alpha1.PerfServer)
			newObject := e.ObjectNew.(*v1alpha1.PerfServer)
			if oldObject.Spec != newObject.Spec {
				return true
			}
			return false
		},
	}

	if err = c.Watch(&source.Kind{Type: &v1alpha1.PerfServer{}}, &handler.EnqueueRequestForObject{}, p); err != nil {
		return err
	}

	return nil
}

var (
	_   reconcile.Reconciler = &ReconcilePerfServer{}
	log                      = logf.Log.WithName("controller_perf_server")
)

type ReconcilePerfServer struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcilePerfServer) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	rl := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	rl.Info("Reconciling PerfServer")

	i := &v1alpha1.PerfServer{}
	if err := r.client.Get(context.TODO(), request.NamespacedName, i); err != nil {
		if k8serrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	defer r.updateStatus(i)

	pc, err := r.newPerfRestClient(i.Spec.ApiUrl, i.Spec.CredentialName, i.Namespace)
	if err != nil {
		return reconcile.Result{}, err
	}

	if err := chain.CreateDefChain(r.client, r.scheme, pc).ServeRequest(i); err != nil {
		i.Status.DetailedMessage = err.Error()
		log.Error(err, "couldn't handle PERF server CR")
		return reconcile.Result{RequeueAfter: 5 * time.Minute}, nil
	}

	rl.Info("Reconciling PerfServer has been finished")
	return reconcile.Result{}, nil
}

func (r ReconcilePerfServer) updateStatus(server *v1alpha1.PerfServer) {
	server.Status.LastTimeUpdated = time.Now()
	if err := r.client.Status().Update(context.TODO(), server); err != nil {
		_ = r.client.Update(context.TODO(), server)
	}
}

func (r ReconcilePerfServer) newPerfRestClient(url, secretName, namespace string) (*perf.PerfClientAdapter, error) {
	credentials, err := perf.GetPerfCredentials(r.client, secretName, namespace)
	if err != nil {
		return nil, err
	}

	perfClient, err := perf.NewRestClient(url, credentials.Username, credentials.Password, credentials.LuminateToken)
	if err != nil {
		return nil, err
	}
	return perfClient, nil
}
