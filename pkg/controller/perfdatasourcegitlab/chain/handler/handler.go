package handler

import (
	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
)

type PerfDataSourceGitLabHandler interface {
	ServeRequest(server *perfApi.PerfDataSourceGitLab) error
}
