package handler

import (
	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
)

type PerfDataSourceJenkinsHandler interface {
	ServeRequest(server *perfApi.PerfDataSourceJenkins) error
}
