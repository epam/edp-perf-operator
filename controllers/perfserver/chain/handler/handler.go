package handler

import (
	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
)

type PerfServerHandler interface {
	ServeRequest(server *perfApi.PerfServer) error
}
