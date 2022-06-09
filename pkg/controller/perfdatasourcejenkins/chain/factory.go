package chain

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfdatasourcejenkins/chain/handler"
)

var log = ctrl.Log.WithName("perf_data_source_handler")

func CreateDefChain(client client.Client, scheme *runtime.Scheme, perfClient perf.PerfClient) handler.PerfDataSourceJenkinsHandler {
	return PutOwnerReference{
		client: client,
		scheme: scheme,
		next: PutDataSource{
			client:     client,
			perfClient: perfClient,
		},
	}
}

func nextServeOrNil(next handler.PerfDataSourceJenkinsHandler, ds *perfApi.PerfDataSourceJenkins) error {
	if next != nil {
		return next.ServeRequest(ds)
	}
	log.Info("handling of perf Jenkins data source has been finished", "name", ds.Name)
	return nil
}
