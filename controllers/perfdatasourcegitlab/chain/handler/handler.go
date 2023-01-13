package handler

import (
	perfApi "github.com/epam/edp-perf-operator/v2/api/edp/v1"
)

type PerfDataSourceGitLabHandler interface {
	ServeRequest(server *perfApi.PerfDataSourceGitLab) error
}
