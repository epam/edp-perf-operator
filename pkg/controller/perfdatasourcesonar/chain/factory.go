package chain

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfdatasourcesonar/chain/handler"
)

var log = ctrl.Log.WithName("perf_data_source_handler")

func CreateDefChain(c client.Client, scheme *runtime.Scheme, perfClient perf.PerfClient) handler.PerfDataSourceSonarHandler {
	return PutOwnerReference{
		client: c,
		scheme: scheme,
		next: PutDataSource{
			client:     c,
			perfClient: perfClient,
		},
	}
}

func nextServeOrNil(next handler.PerfDataSourceSonarHandler, ds *perfApi.PerfDataSourceSonar) error {
	if next != nil {
		if err := next.ServeRequest(ds); err != nil {
			return fmt.Errorf("failed to serve request: %w", err)
		}

		return nil
	}

	log.Info("handling of perf Sonar data source has been finished", "name", ds.Name)

	return nil
}
