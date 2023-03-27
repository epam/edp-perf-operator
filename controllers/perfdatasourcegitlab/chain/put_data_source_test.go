package chain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/api/v1"
	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf/mock"
	"github.com/epam/edp-perf-operator/v2/pkg/model/command"
	"github.com/epam/edp-perf-operator/v2/pkg/model/dto"
)

const (
	gitlabDsType = "GITLAB"
)

func TestPutDataSource_ServeRequest(t *testing.T) {
	t.Parallel()

	scheme := runtime.NewScheme()

	require.NoError(t, corev1.AddToScheme(scheme))
	require.NoError(t, perfApi.AddToScheme(scheme))
	require.NoError(t, codebaseApi.AddToScheme(scheme))

	type parameters struct {
		dataSourceCommandRepositories []string
		dataSourceCommandBranches     []string

		dataSourceActive       bool
		dataSourceRepositories []string
		dataSourceBranches     []string

		activateDataSourceErr   error
		getProjectDataSourceErr error
	}

	tests := []struct {
		name           string
		getProjectFail bool
		status         string
		parameters     parameters
		wantErr        require.ErrorAssertionFunc
	}{
		{
			name:    "should update GitLabDataSource without activating",
			wantErr: require.NoError,
			status:  "created",
			parameters: parameters{
				dataSourceActive:       true,
				dataSourceRepositories: []string{"repo2"},
				dataSourceBranches:     []string{"develop"},

				dataSourceCommandRepositories: []string{"repo2", "repo1"},
				dataSourceCommandBranches:     []string{"develop", "master"},
			},
		},
		{
			name:    "should update GitLabDataSource with activating",
			wantErr: require.NoError,
			status:  "created",
			parameters: parameters{
				dataSourceRepositories: []string{"repo2"},
				dataSourceBranches:     []string{"develop"},

				dataSourceCommandRepositories: []string{"repo2", "repo1"},
				dataSourceCommandBranches:     []string{"develop", "master"},
			},
		},
		{
			name:    "should create GitLabDataSource",
			wantErr: require.NoError,
			status:  "created",
			parameters: parameters{
				dataSourceCommandRepositories: []string{"repo1"},
				dataSourceCommandBranches:     []string{"master"},
			},
		},
		{
			name:    "should not find DataSource in PERF",
			wantErr: require.Error,
			status:  "error",
			parameters: parameters{
				getProjectDataSourceErr: fmt.Errorf("failed"),
			},
		},
		{
			name: "should not activate DataSource",
			parameters: parameters{
				dataSourceRepositories: []string{"repo2"},
				dataSourceBranches:     []string{"develop"},

				dataSourceCommandRepositories: []string{"repo1"},
				dataSourceCommandBranches:     []string{"master"},
				activateDataSourceErr:         fmt.Errorf("failed"),
			},
			wantErr: require.Error,
			status:  "error",
		},
		{
			name:    "should not update DataSource because of missing new parameters",
			wantErr: require.NoError,
			parameters: parameters{
				dataSourceActive:       true,
				dataSourceRepositories: []string{"repo2"},
				dataSourceBranches:     []string{"develop"},

				dataSourceCommandRepositories: []string{"repo2", "repo1"},
				dataSourceCommandBranches:     []string{"develop", "master"},
			},
		},
		{
			name:           "should fail while getting project DataSource",
			getProjectFail: true,
			wantErr:        require.Error,
			status:         "error",
			parameters: parameters{
				getProjectDataSourceErr: fmt.Errorf("failed"),

				dataSourceCommandRepositories: []string{"repo2", "repo1"},
				dataSourceCommandBranches:     []string{"develop", "master"},
			},
		},
		{
			name:    "should run without errors, because there is nothing to update",
			wantErr: require.NoError,
			parameters: parameters{
				dataSourceActive:       true,
				dataSourceRepositories: []string{"repo1"},
				dataSourceBranches:     []string{"master"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			perfDataSourceGitLab := &perfApi.PerfDataSourceGitLab{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: fakeNamespace,
					OwnerReferences: []metav1.OwnerReference{
						{
							Kind: "PerfServer",
							Name: fakeName,
						},
					},
				},
				Spec: perfApi.PerfDataSourceGitLabSpec{
					PerfServerName: fakeName,
					Type:           gitlabDsType,
					Config: perfApi.DataSourceGitLabConfig{
						Repositories: []string{"repo1"},
						Branches:     []string{"master"},
						Url:          fakeName,
					},
				},
			}

			perfServer := &perfApi.PerfServer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fakeName,
					Namespace: fakeNamespace,
				},
				Spec: perfApi.PerfServerSpec{
					ProjectName: fakeName,
				},
			}

			gitLabUserSecret := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      gitLabSecretName,
					Namespace: fakeNamespace,
				},
				Data: map[string][]byte{
					"username": []byte("fake"),
					"password": []byte("fake"),
				},
			}

			objects := []runtime.Object{perfDataSourceGitLab, perfServer, gitLabUserSecret}

			perfClient := new(mock.MockPerfClient)

			var dataSource *dto.DataSource = nil

			if tt.parameters.dataSourceRepositories != nil {
				dataSource = &dto.DataSource{
					Active: tt.parameters.dataSourceActive,
					Type:   gitlabDsType,
					Config: map[string]any{
						"repositories": tt.parameters.dataSourceRepositories,
						"branches":     tt.parameters.dataSourceBranches,
					},
				}
			}
			dataSourceCommand := command.DataSourceCommand{
				Type: gitlabDsType,
				Config: command.DataSourceGitlabConfig{
					Url:          fakeName,
					InstanceId:   fakeName,
					Username:     "fake",
					Password:     "fake",
					Repositories: tt.parameters.dataSourceCommandRepositories,
					Branches:     tt.parameters.dataSourceCommandBranches,
				},
			}

			perfClient.On("CreateDataSource", fakeName, dataSourceCommand).Return(nil)

			perfClient.On("GetProjectDataSource", fakeName, gitlabDsType).
				Return(dataSource, tt.parameters.getProjectDataSourceErr)
			perfClient.On("GetProjectDataSource", fakeName, "").
				Return(dataSource, tt.parameters.getProjectDataSourceErr)

			perfClient.On("UpdateDataSource", dataSourceCommand).Return(nil)

			perfClient.On("ActivateDataSource", fakeName, 0).Return(tt.parameters.activateDataSourceErr)

			putDataSource := PutDataSource{
				client: fake.NewClientBuilder().
					WithScheme(scheme).
					WithRuntimeObjects(objects...).
					Build(),
				perfClient: perfClient,
			}

			tt.wantErr(t, putDataSource.ServeRequest(perfDataSourceGitLab))

			if tt.status != "" {
				require.Equal(t, tt.status, perfDataSourceGitLab.Status.Status)
			}
		})
	}
}
