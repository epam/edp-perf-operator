package chain

import (
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfserver/chain/handler"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var log = ctrl.Log.WithName("perf_server_handler")

func CreateDefChain(client client.Client, scheme *runtime.Scheme, perfClient perf.PerfClient) handler.PerfServerHandler {
	return CheckConnectionToPerf{
		next: PutPerfProject{
			next: PutEdpComponent{
				client: client,
				scheme: scheme,
			},
			perfClient: perfClient,
		},
		client:     client,
		perfClient: perfClient,
	}
}

func nextServeOrNil(next handler.PerfServerHandler, server *v1alpha1.PerfServer) error {
	if next != nil {
		return next.ServeRequest(server)
	}
	log.Info("handling of PerfServer has been finished", "name", server.Name)
	return nil
}
