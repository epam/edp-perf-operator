package handler

import (
	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
)

type PerfDataSourceSonarHandler interface {
	ServeRequest(server *perfApi.PerfDataSourceSonar) error
}
