package chain

import (
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfdatasourcesonar/chain/handler"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var log = ctrl.Log.WithName("perf_data_source_handler")

func CreateDefChain(client client.Client, scheme *runtime.Scheme, perfClient perf.PerfClient) handler.PerfDataSourceSonarHandler {
	return PutOwnerReference{
		client: client,
		scheme: scheme,
		next: PutDataSource{
			client:     client,
			perfClient: perfClient,
		},
	}
}

func nextServeOrNil(next handler.PerfDataSourceSonarHandler, ds *v1alpha1.PerfDataSourceSonar) error {
	if next != nil {
		return next.ServeRequest(ds)
	}
	log.Info("handling of perf Sonar data source has been finished", "name", ds.Name)
	return nil
}
