package perfdatasourcegitlab

import (
	"context"
	v1alpha12 "github.com/epmd-edp/codebase-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/v2/pkg/client/perf"
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfdatasourcegitlab/chain"
	"github.com/epmd-edp/perf-operator/v2/pkg/util/cluster"
	"github.com/epmd-edp/perf-operator/v2/pkg/util/common"
	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"reflect"
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

var (
	log = logf.Log.WithName("controller_perf_data_source_gitlab")
)

func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	scheme := mgr.GetScheme()
	addKnownTypes(scheme)
	return &ReconcilePerfDataSourceGitLab{
		client: mgr.GetClient(),
		scheme: scheme,
	}
}

func addKnownTypes(scheme *runtime.Scheme) {
	schemeGroupVersion := schema.GroupVersion{Group: "v2.edp.epam.com", Version: "v1alpha1"}
	scheme.AddKnownTypes(schemeGroupVersion,
		&v1alpha12.Codebase{},
		&v1alpha12.CodebaseList{},
	)
	metav1.AddToGroupVersion(scheme, schemeGroupVersion)
}

func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("perfdatasourcegitlab-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	p := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			oldPds := e.ObjectOld.(*v1alpha1.PerfDataSourceGitLab).Spec.Config.Branches
			newPds := e.ObjectNew.(*v1alpha1.PerfDataSourceGitLab).Spec.Config.Branches
			return dataSourceUpdated(oldPds, newPds)
		},
	}

	if err = c.Watch(&source.Kind{Type: &v1alpha1.PerfDataSourceGitLab{}}, &handler.EnqueueRequestForObject{}, p); err != nil {
		return err
	}

	return nil
}

func dataSourceUpdated(old, new []string) bool {
	common.SortArray(old)
	common.SortArray(new)
	return !reflect.DeepEqual(old, new)
}

var _ reconcile.Reconciler = &ReconcilePerfDataSourceGitLab{}

type ReconcilePerfDataSourceGitLab struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcilePerfDataSourceGitLab) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	rl := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	rl.V(2).Info("Reconciling PerfDataSourceGitLab")

	i := &v1alpha1.PerfDataSourceGitLab{}
	if err := r.client.Get(context.TODO(), request.NamespacedName, i); err != nil {
		if k8serrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	defer r.updateStatus(i)

	ps, err := cluster.GetPerfServerCr(r.client, i.Spec.PerfServerName, i.Namespace)
	if err != nil {
		return reconcile.Result{}, errors.Wrapf(err, "couldn't get %v PerfServer from cluster", i.Spec.PerfServerName)
	}

	if !ps.Status.Available {
		log.Info("Perf instance is unavailable. skip creating/updating data source in PERF", "name", ps.Name)
		return reconcile.Result{RequeueAfter: 2 * time.Minute}, nil
	}

	pc, err := r.newPerfRestClient(ps.Spec.ApiUrl, ps.Spec.CredentialName, ps.Namespace)
	if err != nil {
		return reconcile.Result{}, err
	}

	if err := chain.CreateDefChain(r.client, r.scheme, pc).ServeRequest(i); err != nil {
		return reconcile.Result{}, err
	}

	rl.Info("Reconciling PerfDataSourceGitLab has been finished")
	return reconcile.Result{}, nil
}

func (r ReconcilePerfDataSourceGitLab) updateStatus(ds *v1alpha1.PerfDataSourceGitLab) {
	if err := r.client.Status().Update(context.TODO(), ds); err != nil {
		_ = r.client.Update(context.TODO(), ds)
	}
}

func (r ReconcilePerfDataSourceGitLab) newPerfRestClient(url, secretName, namespace string) (*perf.PerfClientAdapter, error) {
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
