package controller

import (
	"github.com/epmd-edp/perf-operator/pkg/controller/perfdatasource"
)

func init() {
	AddToManagerFuncs = append(AddToManagerFuncs, perfdatasource.Add)
}
