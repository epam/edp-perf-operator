package controller

import (
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfdatasourcegitlab"
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfdatasourcejenkins"
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfdatasourcesonar"
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfserver"
)

func init() {
	AddToManagerFuncs = append(AddToManagerFuncs, perfserver.Add, perfdatasourcejenkins.Add,
		perfdatasourcesonar.Add, perfdatasourcegitlab.Add)
}
