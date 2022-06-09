package handler

import (
	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
)

type PerfDataSourceSonarHandler interface {
	ServeRequest(server *perfApi.PerfDataSourceSonar) error
}
