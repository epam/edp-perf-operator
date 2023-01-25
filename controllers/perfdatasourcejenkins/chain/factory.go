package chain

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
	"github.com/epam/edp-perf-operator/v2/controllers/perfdatasourcejenkins/chain/handler"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
)

var log = ctrl.Log.WithName("perf_data_source_handler")

func CreateDefChain(c client.Client, scheme *runtime.Scheme, perfClient perf.PerfClient) handler.PerfDataSourceJenkinsHandler {
	return PutOwnerReference{
		client: c,
		scheme: scheme,
		next: PutDataSource{
			client:     c,
			perfClient: perfClient,
		},
	}
}

func nextServeOrNil(next handler.PerfDataSourceJenkinsHandler, ds *perfApi.PerfDataSourceJenkins) error {
	if next != nil {
		if err := next.ServeRequest(ds); err != nil {
			return fmt.Errorf("failed to serve request: %w", err)
		}

		return nil
	}

	log.Info("handling of perf Jenkins data source has been finished", "name", ds.Name)

	return nil
}
