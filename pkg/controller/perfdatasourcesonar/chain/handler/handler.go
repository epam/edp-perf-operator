package handler

import "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"

type PerfDataSourceSonarHandler interface {
	ServeRequest(server *v1alpha1.PerfDataSourceSonar) error
}
