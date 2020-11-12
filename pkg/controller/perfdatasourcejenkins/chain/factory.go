package chain

import (
	"github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/v2/pkg/client/perf"
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfdatasourcejenkins/chain/handler"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("perf_data_source_handler")

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

func nextServeOrNil(next handler.PerfDataSourceJenkinsHandler, ds *v1alpha1.PerfDataSourceJenkins) error {
	if next != nil {
		return next.ServeRequest(ds)
	}
	log.Info("handling of perf Jenkins data source has been finished", "name", ds.Name)
	return nil
}
