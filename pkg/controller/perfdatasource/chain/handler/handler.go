package handler

import "github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"

type PerfDataSourceHandler interface {
	ServeRequest(server *v1alpha1.PerfDataSource) error
}
