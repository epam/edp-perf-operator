package perfserver

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	perfApi "github.com/epam/edp-perf-operator/v2/api/edp/v1"
	"github.com/epam/edp-perf-operator/v2/controllers/perfserver/chain"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
)

const (
	duration5Minutes = 5 * time.Minute
)

func NewReconcilePerfServer(c client.Client, scheme *runtime.Scheme, log logr.Logger) *ReconcilePerfServer {
	return &ReconcilePerfServer{
		client: c,
		scheme: scheme,
		log:    log.WithName("perf-server"),
	}
}

type ReconcilePerfServer struct {
	client client.Client
	scheme *runtime.Scheme
	log    logr.Logger
}

func (r *ReconcilePerfServer) SetupWithManager(mgr ctrl.Manager) error {
	p := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			oldObject, ok := e.ObjectOld.(*perfApi.PerfServer)
			if !ok {
				return false
			}

			newObject, ok := e.ObjectNew.(*perfApi.PerfServer)
			if !ok {
				return false
			}

			return oldObject.Spec != newObject.Spec
		},
	}

	if err := ctrl.NewControllerManagedBy(mgr).
		For(&perfApi.PerfServer{}, builder.WithPredicates(p)).
		Complete(r); err != nil {
		return fmt.Errorf("failed to build controller: %w", err)
	}

	return nil
}

//+kubebuilder:rbac:groups=v2.edp.epam.com,namespace=placeholder,resources=perfservers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=v2.edp.epam.com,namespace=placeholder,resources=perfservers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=v2.edp.epam.com,namespace=placeholder,resources=perfservers/finalizers,verbs=update

func (r *ReconcilePerfServer) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := r.log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	log.Info("Reconciling PerfServer")

	i := &perfApi.PerfServer{}
	if err := r.client.Get(ctx, request.NamespacedName, i); err != nil {
		if k8sErrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, fmt.Errorf("failed to get perf server: %w", err)
	}

	defer r.updateStatus(ctx, i)

	pc, err := r.newPerfRestClient(i.Spec.ApiUrl, i.Spec.CredentialName, i.Namespace)
	if err != nil {
		return reconcile.Result{}, err
	}

	if err = chain.CreateDefChain(r.client, r.scheme, pc).ServeRequest(i); err != nil {
		i.Status.DetailedMessage = err.Error()
		log.Error(err, "couldn't handle PERF server CR")

		return reconcile.Result{RequeueAfter: duration5Minutes}, nil
	}

	log.Info("Reconciling PerfServer has been finished")

	return reconcile.Result{}, nil
}

func (r ReconcilePerfServer) updateStatus(ctx context.Context, server *perfApi.PerfServer) {
	server.Status.LastTimeUpdated = metaV1.Now()
	if err := r.client.Status().Update(ctx, server); err != nil {
		_ = r.client.Update(ctx, server)
	}
}

func (r ReconcilePerfServer) newPerfRestClient(url, secretName, namespace string) (*perf.PerfClientAdapter, error) {
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
