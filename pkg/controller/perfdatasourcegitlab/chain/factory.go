package chain

import (
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfdatasourcegitlab/chain/handler"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var log = ctrl.Log.WithName("perf_data_source_gitlab_handler")

func CreateDefChain(client client.Client, scheme *runtime.Scheme, perfClient perf.PerfClient) handler.PerfDataSourceGitLabHandler {
	return PutOwnerReference{
		client: client,
		scheme: scheme,
		next: PutDataSource{
			client:     client,
			perfClient: perfClient,
		},
	}
}

func nextServeOrNil(next handler.PerfDataSourceGitLabHandler, ds *v1alpha1.PerfDataSourceGitLab) error {
	if next != nil {
		return next.ServeRequest(ds)
	}
	log.Info("handling of perf GitLab data source has been finished", "name", ds.Name)
	return nil
}
