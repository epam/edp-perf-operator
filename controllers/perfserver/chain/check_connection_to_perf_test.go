package chain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	perfApi "github.com/epam/edp-perf-operator/v2/api/edp/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf/mock"
)

func TestCheckConnectionToPerf_ShouldBeExecutedSuccessfully(t *testing.T) {
	ps := &perfApi.PerfServer{}

	objs := []runtime.Object{
		ps,
	}
	s := scheme.Scheme
	s.AddKnownTypes(v1.SchemeGroupVersion, ps)

	mPerfCl := new(mock.MockPerfClient)
	perf := CheckConnectionToPerf{
		client:     fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
		perfClient: mPerfCl,
	}

	mPerfCl.On("Connected").Return(true, nil)

	psr := &perfApi.PerfServer{}
	err := perf.ServeRequest(psr)
	assert.NoError(t, err)
	assert.Equal(t, true, psr.Status.Available)
}

func TestCheckConnectionToPerf_ShouldBeExecutedWithError(t *testing.T) {
	ps := &perfApi.PerfServer{}

	objs := []runtime.Object{
		ps,
	}
	s := scheme.Scheme
	s.AddKnownTypes(v1.SchemeGroupVersion, ps)

	mPerfCl := new(mock.MockPerfClient)
	perf := CheckConnectionToPerf{
		client:     fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
		perfClient: mPerfCl,
	}

	mPerfCl.On("Connected").Return(false, fmt.Errorf("failed"))

	psr := &perfApi.PerfServer{
		Status: perfApi.PerfServerStatus{},
	}
	err := perf.ServeRequest(psr)
	assert.Error(t, err)
}

func TestCheckConnectionToPerf_ShouldNotBeUpdated(t *testing.T) {
	mPerfCl := new(mock.MockPerfClient)
	perf := CheckConnectionToPerf{
		client:     fake.NewClientBuilder().WithRuntimeObjects([]runtime.Object{}...).Build(),
		perfClient: mPerfCl,
	}

	mPerfCl.On("Connected").Return(true, nil)

	psr := &perfApi.PerfServer{
		Status: perfApi.PerfServerStatus{},
	}
	err := perf.ServeRequest(psr)
	assert.NoError(t, err)
}
