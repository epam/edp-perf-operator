package apis

import (
	"github.com/epmd-edp/perf-operator/pkg/apis/edp/v1alpha1"
)

func init() {
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
