package perfdatasourcejenkins

import (
	"context"
	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfdatasourcejenkins/chain"
	"github.com/epam/edp-perf-operator/v2/pkg/util/cluster"
	"github.com/epam/edp-perf-operator/v2/pkg/util/common"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

var _ reconcile.Reconciler = &ReconcilePerfDataSourceJenkins{}

func NewReconcilePerfDataSourceJenkins(client client.Client, scheme *runtime.Scheme, log logr.Logger) *ReconcilePerfDataSourceJenkins {
	return &ReconcilePerfDataSourceJenkins{
		client: client,
		scheme: scheme,
		log:    log.WithName("perf-data-source-jenkins"),
	}
}

type ReconcilePerfDataSourceJenkins struct {
	client client.Client
	scheme *runtime.Scheme
	log    logr.Logger
}

func (r *ReconcilePerfDataSourceJenkins) SetupWithManager(mgr ctrl.Manager) error {
	p := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			oldJn := e.ObjectOld.(*perfApi.PerfDataSourceJenkins).Spec.Config.JobNames
			newJn := e.ObjectNew.(*perfApi.PerfDataSourceJenkins).Spec.Config.JobNames
			return dataSourceUpdated(oldJn, newJn)
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&perfApi.PerfDataSourceJenkins{}, builder.WithPredicates(p)).
		Complete(r)
}

func dataSourceUpdated(old, new []string) bool {
	common.SortArray(old)
	common.SortArray(new)
	return !reflect.DeepEqual(old, new)
}

func (r *ReconcilePerfDataSourceJenkins) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := r.log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	log.V(2).Info("Reconciling PerfDataSourceJenkins")

	i := &perfApi.PerfDataSourceJenkins{}
	if err := r.client.Get(ctx, request.NamespacedName, i); err != nil {
		if k8serrors.IsNotFound(err) {
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

	log.Info("Reconciling PerfDataSourceJenkins has been finished")
	return reconcile.Result{}, nil
}

func (r ReconcilePerfDataSourceJenkins) updateStatus(ctx context.Context, ds *perfApi.PerfDataSourceJenkins) {
	if err := r.client.Status().Update(ctx, ds); err != nil {
		_ = r.client.Update(ctx, ds)
	}
}

func (r ReconcilePerfDataSourceJenkins) newPerfRestClient(url, secretName, namespace string) (*perf.PerfClientAdapter, error) {
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
