package chain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf/mock"
)

const fakeName = "fake-name"

func TestPutPerfProject_ProjectExistsShouldBeExecutedSuccessfully(t *testing.T) {
	mPerfCl := new(mock.MockPerfClient)
	project := PutPerfProject{
		perfClient: mPerfCl,
	}

	mPerfCl.On("ProjectExists", fakeName).Return(true, nil)

	psr := &perfApi.PerfServer{
		Spec: perfApi.PerfServerSpec{
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

	psr := &perfApi.PerfServer{
		Spec: perfApi.PerfServerSpec{
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

	psr := &perfApi.PerfServer{
		Spec: perfApi.PerfServerSpec{
			ProjectName: fakeName,
		},
	}
	assert.Error(t, project.ServeRequest(psr))
}
