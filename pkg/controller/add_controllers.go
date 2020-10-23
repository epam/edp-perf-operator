package controller

import (
	"github.com/epmd-edp/perf-operator/pkg/controller/perfdatasource"
	"github.com/epmd-edp/perf-operator/pkg/controller/perfserver"
)

func init() {
	AddToManagerFuncs = append(AddToManagerFuncs, perfserver.Add, perfdatasource.Add)
}
