package perfdatasourcegitlab

import (
	"context"
	"reflect"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfdatasourcegitlab/chain"
	"github.com/epam/edp-perf-operator/v2/pkg/util/cluster"
	"github.com/epam/edp-perf-operator/v2/pkg/util/common"
)

func NewReconcilePerfDataSourceGitLab(client client.Client, scheme *runtime.Scheme, log logr.Logger) *ReconcilePerfDataSourceGitLab {
	return &ReconcilePerfDataSourceGitLab{
		client: client,
		scheme: scheme,
		log:    log.WithName("perf-data-source-gitlab"),
	}
}

type ReconcilePerfDataSourceGitLab struct {
	client client.Client
	scheme *runtime.Scheme
	log    logr.Logger
}

func (r *ReconcilePerfDataSourceGitLab) SetupWithManager(mgr ctrl.Manager) error {
	p := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			oldPds := e.ObjectOld.(*perfApi.PerfDataSourceGitLab).Spec.Config.Branches
			newPds := e.ObjectNew.(*perfApi.PerfDataSourceGitLab).Spec.Config.Branches
			return dataSourceUpdated(oldPds, newPds)
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&perfApi.PerfDataSourceGitLab{}, builder.WithPredicates(p)).
		Complete(r)
}

func dataSourceUpdated(old, new []string) bool {
	common.SortArray(old)
	common.SortArray(new)
	return !reflect.DeepEqual(old, new)
}

func (r *ReconcilePerfDataSourceGitLab) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := r.log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	log.V(2).Info("Reconciling PerfDataSourceGitLab")

	i := &perfApi.PerfDataSourceGitLab{}
	if err := r.client.Get(ctx, request.NamespacedName, i); err != nil {
		if k8sErrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	defer r.updateStatus(ctx, i)

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

	log.Info("Reconciling PerfDataSourceGitLab has been finished")
	return reconcile.Result{}, nil
}

func (r ReconcilePerfDataSourceGitLab) updateStatus(ctx context.Context, ds *perfApi.PerfDataSourceGitLab) {
	if err := r.client.Status().Update(ctx, ds); err != nil {
		_ = r.client.Update(ctx, ds)
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
