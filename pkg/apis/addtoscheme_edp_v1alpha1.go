package apis

import (
	"github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"
)

func init() {
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
