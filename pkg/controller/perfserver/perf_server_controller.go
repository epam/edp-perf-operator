package perfserver

import (
	"context"
	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfserver/chain"
	"github.com/go-logr/logr"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

type ReconcilePerfServer struct {
	Client client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

func (r *ReconcilePerfServer) SetupWithManager(mgr ctrl.Manager) error {
	p := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			oldObject := e.ObjectOld.(*perfApi.PerfServer)
			newObject := e.ObjectNew.(*perfApi.PerfServer)
			if oldObject.Spec != newObject.Spec {
				return true
			}
			return false
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&perfApi.PerfServer{}, builder.WithPredicates(p)).
		Complete(r)
}

func (r *ReconcilePerfServer) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := r.Log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	log.Info("Reconciling PerfServer")

	i := &perfApi.PerfServer{}
	if err := r.Client.Get(ctx, request.NamespacedName, i); err != nil {
		if k8serrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	defer r.updateStatus(ctx, i)

	pc, err := r.newPerfRestClient(i.Spec.ApiUrl, i.Spec.CredentialName, i.Namespace)
	if err != nil {
		return reconcile.Result{}, err
	}

	if err := chain.CreateDefChain(r.Client, r.Scheme, pc).ServeRequest(i); err != nil {
		i.Status.DetailedMessage = err.Error()
		log.Error(err, "couldn't handle PERF server CR")
		return reconcile.Result{RequeueAfter: 5 * time.Minute}, nil
	}

	log.Info("Reconciling PerfServer has been finished")
	return reconcile.Result{}, nil
}

func (r ReconcilePerfServer) updateStatus(ctx context.Context, server *perfApi.PerfServer) {
	server.Status.LastTimeUpdated = time.Now()
	if err := r.Client.Status().Update(ctx, server); err != nil {
		_ = r.Client.Update(ctx, server)
	}
}

func (r ReconcilePerfServer) newPerfRestClient(url, secretName, namespace string) (*perf.PerfClientAdapter, error) {
	credentials, err := perf.GetPerfCredentials(r.Client, secretName, namespace)
	if err != nil {
		return nil, err
	}

	perfClient, err := perf.NewRestClient(url, credentials.Username, credentials.Password, credentials.LuminateToken)
	if err != nil {
		return nil, err
	}
	return perfClient, nil
}
