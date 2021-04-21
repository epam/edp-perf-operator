package chain

import (
	"errors"
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const fakeName = "fake-name"

func TestPutPerfProject_ProjectExistsShouldBeExecutedSuccessfully(t *testing.T) {
	mPerfCl := new(mock.MockPerfClient)
	project := PutPerfProject{
		perfClient: mPerfCl,
	}

	mPerfCl.On("ProjectExists", fakeName).Return(true, nil)

	psr := &v1alpha1.PerfServer{
		Spec: v1alpha1.PerfServerSpec{
			ProjectName: fakeName,
		},
	}
	assert.NoError(t, project.ServeRequest(psr))
}

func TestPutPerfProject_ProjectDoesntExistShouldBeExecutedSuccessfully(t *testing.T) {
	mPerfCl := new(mock.MockPerfClient)
	project := PutPerfProject{
		perfClient: mPerfCl,
	}

	mPerfCl.On("ProjectExists", fakeName).Return(false, nil)

	psr := &v1alpha1.PerfServer{
		Spec: v1alpha1.PerfServerSpec{
			ProjectName: fakeName,
		},
	}
	assert.Error(t, project.ServeRequest(psr))
}

func TestPutPerfProject_ThrowErrorDuringProjectExistsCall(t *testing.T) {
	mPerfCl := new(mock.MockPerfClient)
	project := PutPerfProject{
		perfClient: mPerfCl,
	}

	mPerfCl.On("ProjectExists", fakeName).Return(false, errors.New("failed"))

	psr := &v1alpha1.PerfServer{
		Spec: v1alpha1.PerfServerSpec{
			ProjectName: fakeName,
		},
	}
	assert.Error(t, project.ServeRequest(psr))
}
