package handler

import "github.com/epmd-edp/perf-operator/pkg/apis/edp/v1alpha1"

type PerfServerHandler interface {
	ServeRequest(server *v1alpha1.PerfServer) error
}
