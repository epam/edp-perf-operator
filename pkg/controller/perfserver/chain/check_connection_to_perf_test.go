package chain

import (
	"errors"
	"github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/v2/pkg/client/perf/mock"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

func TestCheckConnectionToPerf_ShouldBeExecutedSuccessfully(t *testing.T) {
	ps := &v1alpha1.PerfServer{}

	objs := []runtime.Object{
		ps,
	}
	s := scheme.Scheme
	s.AddKnownTypes(v1.SchemeGroupVersion, ps)

	mPerfCl := new(mock.MockPerfClient)
	perf := CheckConnectionToPerf{
		client:     fake.NewFakeClient(objs...),
		perfClient: mPerfCl,
	}

	mPerfCl.On("Connected").Return(true, nil)

	psr := &v1alpha1.PerfServer{}
	err := perf.ServeRequest(psr)
	assert.NoError(t, err)
	assert.Equal(t, true, psr.Status.Available)
}

func TestCheckConnectionToPerf_ShouldBeExecutedWithError(t *testing.T) {
	ps := &v1alpha1.PerfServer{}

	objs := []runtime.Object{
		ps,
	}
	s := scheme.Scheme
	s.AddKnownTypes(v1.SchemeGroupVersion, ps)

	mPerfCl := new(mock.MockPerfClient)
	perf := CheckConnectionToPerf{
		client:     fake.NewFakeClient(objs...),
		perfClient: mPerfCl,
	}

	mPerfCl.On("Connected").Return(false, errors.New("failed"))

	psr := &v1alpha1.PerfServer{
		Status: v1alpha1.PerfServerStatus{},
	}
	err := perf.ServeRequest(psr)
	assert.Error(t, err)
}

func TestCheckConnectionToPerf_ShouldNotBeUpdated(t *testing.T) {
	mPerfCl := new(mock.MockPerfClient)
	perf := CheckConnectionToPerf{
		client:     fake.NewFakeClient([]runtime.Object{}...),
		perfClient: mPerfCl,
	}

	mPerfCl.On("Connected").Return(true, nil)

	psr := &v1alpha1.PerfServer{
		Status: v1alpha1.PerfServerStatus{},
	}
	err := perf.ServeRequest(psr)
	assert.NoError(t, err)
}
