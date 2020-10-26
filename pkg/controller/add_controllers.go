package controller

import (
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfdatasource"
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfserver"
)

func init() {
	AddToManagerFuncs = append(AddToManagerFuncs, perfserver.Add, perfdatasource.Add)
}
