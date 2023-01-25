package perfdatasourcejenkins

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/go-logr/logr"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
	"github.com/epam/edp-perf-operator/v2/controllers/perfdatasourcejenkins/chain"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/util/cluster"
	"github.com/epam/edp-perf-operator/v2/pkg/util/common"
)

var _ reconcile.Reconciler = &ReconcilePerfDataSourceJenkins{}

func NewReconcilePerfDataSourceJenkins(c client.Client, scheme *runtime.Scheme, log logr.Logger) *ReconcilePerfDataSourceJenkins {
	return &ReconcilePerfDataSourceJenkins{
		client: c,
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
			oldDataSource, ok := e.ObjectOld.(*perfApi.PerfDataSourceJenkins)
			if !ok {
				return false
			}

			newDataSource, ok := e.ObjectNew.(*perfApi.PerfDataSourceJenkins)
			if !ok {
				return false
			}

			oldJn := oldDataSource.Spec.Config.JobNames
			newJn := newDataSource.Spec.Config.JobNames

			return dataSourceUpdated(oldJn, newJn)
		},
	}

	if err := ctrl.NewControllerManagedBy(mgr).
		For(&perfApi.PerfDataSourceJenkins{}, builder.WithPredicates(p)).
		Complete(r); err != nil {
		return fmt.Errorf("failed to build controller manager: %w", err)
	}

	return nil
}

func dataSourceUpdated(oldDataSource, newDataSource []string) bool {
	common.SortArray(oldDataSource)
	common.SortArray(newDataSource)

	return !reflect.DeepEqual(oldDataSource, newDataSource)
}

//+kubebuilder:rbac:groups=v2.edp.epam.com,namespace=placeholder,resources=perfdatasourcejenkinses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=v2.edp.epam.com,namespace=placeholder,resources=perfdatasourcejenkinses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=v2.edp.epam.com,namespace=placeholder,resources=perfdatasourcejenkinses/finalizers,verbs=update

func (r *ReconcilePerfDataSourceJenkins) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	logger := r.log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	logger.V(2).Info("Reconciling PerfDataSourceJenkins")

	i := &perfApi.PerfDataSourceJenkins{}
	if err := r.client.Get(ctx, request.NamespacedName, i); err != nil {
		if k8sErrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, fmt.Errorf("failed to get perf data source jenkins %s: %w", request.Namespace, err)
	}

	defer r.updateStatus(ctx, i)

	ps, err := cluster.GetPerfServerCr(r.client, i.Spec.PerfServerName, i.Namespace)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to get %v PerfServer from cluster: %w", i.Spec.PerfServerName, err)
	}

	if !ps.Status.Available {
		logger.Info("Perf instance is unavailable. skip creating/updating data source in PERF", "name", ps.Name)

		return reconcile.Result{RequeueAfter: 2 * time.Minute}, nil
	}

	pc, err := r.newPerfRestClient(ps.Spec.ApiUrl, ps.Spec.CredentialName, ps.Namespace)
	if err != nil {
		return reconcile.Result{}, err
	}

	if err = chain.CreateDefChain(r.client, r.scheme, pc).ServeRequest(i); err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to create default chain: %w", err)
	}

	logger.Info("Reconciling PerfDataSourceJenkins has been finished")

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
		return nil, fmt.Errorf("failed to get perf credentials: %w", err)
	}

	perfClient, err := perf.NewRestClient(url, credentials.Username, credentials.Password, credentials.LuminateToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return perfClient, nil
}
