package mock

import (
	"github.com/epam/edp-perf-operator/v2/pkg/model/command"
	"github.com/epam/edp-perf-operator/v2/pkg/model/dto"
	"github.com/stretchr/testify/mock"
)

type MockPerfClient struct {
	mock.Mock
}

func (m MockPerfClient) Connected() (bool, error) {
	args := m.Called()
	return args.Get(0).(bool), args.Error(1)
}

func (m MockPerfClient) GetProject(name string) (ds *dto.PerfProject, err error) {
	panic("implement me")
}

func (m MockPerfClient) ProjectExists(name string) (bool, error) {
	args := m.Called(name)
	return args.Get(0).(bool), args.Error(1)
}

func (m MockPerfClient) GetProjectDataSource(projectName, dsType string) (*dto.DataSource, error) {
	args := m.Called(projectName, dsType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.DataSource), args.Error(1)
}

func (m MockPerfClient) CreateDataSource(projectName string, command command.DataSourceCommand) error {
	args := m.Called(projectName, command)
	return args.Error(0)
}

func (m MockPerfClient) ActivateDataSource(projectName string, dataSourceId int) error {
	args := m.Called(projectName, dataSourceId)
	return args.Error(0)
}

func (m MockPerfClient) UpdateDataSource(command command.DataSourceCommand) error {
	args := m.Called(command)
	return args.Error(0)
}
