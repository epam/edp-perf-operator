package chain

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfserver/chain/handler"
)

var log = ctrl.Log.WithName("perf_server_handler")

func CreateDefChain(c client.Client, scheme *runtime.Scheme, perfClient perf.PerfClient) handler.PerfServerHandler {
	return CheckConnectionToPerf{
		next: PutPerfProject{
			next: PutEdpComponent{
				client: c,
				scheme: scheme,
			},
			perfClient: perfClient,
		},
		client:     c,
		perfClient: perfClient,
	}
}

func nextServeOrNil(next handler.PerfServerHandler, server *perfApi.PerfServer) error {
	if next != nil {
		if err := next.ServeRequest(server); err != nil {
			return fmt.Errorf("failed to serve request: %w", err)
		}

		return nil
	}

	log.Info("handling of PerfServer has been finished", "name", server.Name)

	return nil
}
