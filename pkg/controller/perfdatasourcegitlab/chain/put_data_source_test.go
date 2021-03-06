package chain

import (
	"errors"
	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf/mock"
	"github.com/epam/edp-perf-operator/v2/pkg/model/command"
	"github.com/epam/edp-perf-operator/v2/pkg/model/dto"
	"github.com/stretchr/testify/assert"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

const (
	gitlabDsType     = "GITLAB"
	fakeCodebaseName = "stub-val"
)

func init() {
	utilruntime.Must(perfApi.AddToScheme(scheme.Scheme))
	utilruntime.Must(codebaseApi.AddToScheme(scheme.Scheme))
}

func TestPutDataSource_ShouldUpdateGitLabDataSourceWithoutActivating(t *testing.T) {
	pds := &v1alpha1.PerfDataSourceGitLab{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
		Spec: v1alpha1.PerfDataSourceGitLabSpec{
			PerfServerName: fakeName,
			Type:           gitlabDsType,
			Config: v1alpha1.DataSourceGitLabConfig{
				Repositories: []string{"repo1"},
				Branches:     []string{"master"},
				Url:          fakeName,
			},
		},
	}

	ps := &v1alpha1.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: v1alpha1.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	sec := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      gitLabSecretName,
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("fake"),
			"password": []byte("fake"),
		},
	}

	mPerfCl := new(mock.MockPerfClient)
	ch := PutDataSource{
		client:     fake.NewFakeClient(pds, ps, sec),
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, gitlabDsType).
		Return(&dto.DataSource{
			Active: true,
			Type:   gitlabDsType,
			Config: map[string]interface{}{
				"repositories": []interface{}{"repo2"},
				"branches":     []interface{}{"develop"},
			},
		}, nil)

	mPerfCl.On("UpdateDataSource", command.DataSourceCommand{
		Type: gitlabDsType,
		Config: command.DataSourceGitlabConfig{
			Repositories:   []string{"repo2", "repo1"},
			Url:            fakeName,
			InstanceId:     fakeName,
			WithMembership: false,
			AllPublic:      false,
			AllBranches:    false,
			Branches:       []string{"develop", "master"},
			Username:       "fake",
			Password:       "fake",
		},
	}).Return(nil)

	assert.NoError(t, ch.ServeRequest(pds))
	assert.Equal(t, "created", pds.Status.Status)
}

func TestPutDataSource_ShouldUpdateGitLabDataSourceWithActivating(t *testing.T) {
	pds := &v1alpha1.PerfDataSourceGitLab{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
		Spec: v1alpha1.PerfDataSourceGitLabSpec{
			PerfServerName: fakeName,
			Type:           gitlabDsType,
			Config: v1alpha1.DataSourceGitLabConfig{
				Repositories: []string{"repo1"},
				Branches:     []string{"master"},
				Url:          fakeName,
			},
		},
	}

	ps := &v1alpha1.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: v1alpha1.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	sec := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      gitLabSecretName,
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("fake"),
			"password": []byte("fake"),
		},
	}

	cl := fake.NewClientBuilder().
		WithScheme(scheme.Scheme).
		WithObjects(pds, ps, sec).
		Build()

	mPerfCl := new(mock.MockPerfClient)
	ch := PutDataSource{
		client:     cl,
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, gitlabDsType).
		Return(&dto.DataSource{
			Active: false,
			Type:   gitlabDsType,
			Config: map[string]interface{}{
				"repositories": []interface{}{"repo2"},
				"branches":     []interface{}{"develop"},
			},
		}, nil)

	mPerfCl.On("UpdateDataSource", command.DataSourceCommand{
		Type: gitlabDsType,
		Config: command.DataSourceGitlabConfig{
			Repositories:   []string{"repo2", "repo1"},
			Url:            fakeName,
			InstanceId:     fakeName,
			WithMembership: false,
			AllPublic:      false,
			AllBranches:    false,
			Branches:       []string{"develop", "master"},
			Username:       "fake",
			Password:       "fake",
		},
	}).Return(nil)

	mPerfCl.On("ActivateDataSource", fakeName, 0).Return(nil)

	assert.NoError(t, ch.ServeRequest(pds))
	assert.Equal(t, "created", pds.Status.Status)
}

func TestPutDataSource_ShouldCreateGitLabDataSource(t *testing.T) {
	pds := &v1alpha1.PerfDataSourceGitLab{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
		Spec: v1alpha1.PerfDataSourceGitLabSpec{
			PerfServerName: fakeName,
			Type:           gitlabDsType,
			Config: v1alpha1.DataSourceGitLabConfig{
				Repositories: []string{"repo1"},
				Branches:     []string{"master"},
				Url:          fakeName,
			},
		},
	}

	ps := &v1alpha1.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: v1alpha1.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	sec := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      gitLabSecretName,
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("fake"),
			"password": []byte("fake"),
		},
	}

	cl := fake.NewClientBuilder().
		WithScheme(scheme.Scheme).
		WithObjects(pds, ps, sec).
		Build()

	mPerfCl := new(mock.MockPerfClient)
	ch := PutDataSource{
		client:     cl,
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, gitlabDsType).Return(nil, nil)

	mPerfCl.On("CreateDataSource", fakeName, command.DataSourceCommand{
		Type: gitlabDsType,
		Config: command.DataSourceGitlabConfig{
			Repositories: []string{"repo1"},
			Branches:     []string{"master"},
			Url:          fakeName,
			InstanceId:   fakeName,
			Username:     "fake",
			Password:     "fake",
		},
	}).Return(nil)

	assert.NoError(t, ch.ServeRequest(pds))
	assert.Equal(t, "created", pds.Status.Status)
}

func TestPutDataSource_ShouldNotFindDataSourceInPERF(t *testing.T) {
	ps := &v1alpha1.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: v1alpha1.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	objs := []runtime.Object{
		ps,
	}

	mPerfCl := new(mock.MockPerfClient)
	ch := PutDataSource{
		client:     fake.NewFakeClient(objs...),
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, "").Return(nil, errors.New("failed"))

	pds := &v1alpha1.PerfDataSourceGitLab{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
	}

	assert.Error(t, ch.ServeRequest(pds))
	assert.Equal(t, "error", pds.Status.Status)
}

func TestPutDataSource_ShouldNotActivateDataSource(t *testing.T) {
	pds := &v1alpha1.PerfDataSourceGitLab{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
		Spec: v1alpha1.PerfDataSourceGitLabSpec{
			Type: gitlabDsType,
			Config: v1alpha1.DataSourceGitLabConfig{
				Repositories: []string{"repo1"},
				Branches:     []string{"master"},
				Url:          fakeName,
			},
		},
	}

	ps := &v1alpha1.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: v1alpha1.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	sec := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      gitLabSecretName,
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("fake"),
			"password": []byte("fake"),
		},
	}

	objs := []runtime.Object{
		pds, ps, sec,
	}

	mPerfCl := new(mock.MockPerfClient)
	ch := PutDataSource{
		client:     fake.NewFakeClient(objs...),
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, gitlabDsType).
		Return(&dto.DataSource{
			Active: false,
			Type:   gitlabDsType,
			Config: map[string]interface{}{
				"repositories": []interface{}{"repo2"},
				"branches":     []interface{}{"develop"},
			},
		}, nil)

	mPerfCl.On("UpdateDataSource", command.DataSourceCommand{
		Type: gitlabDsType,
		Config: command.DataSourceGitlabConfig{
			Repositories:   []string{"repo2", "repo1"},
			Url:            fakeName,
			InstanceId:     fakeName,
			WithMembership: false,
			AllPublic:      false,
			AllBranches:    false,
			Branches:       []string{"develop", "master"},
			Username:       "fake",
			Password:       "fake",
		},
	}).Return(nil)

	mPerfCl.On("ActivateDataSource", fakeName, 0).Return(errors.New("failed"))

	assert.Error(t, ch.ServeRequest(pds))
	assert.Equal(t, "error", pds.Status.Status)
}

func TestPutDataSource_ShouldNotUpdateDataSourceBecauseOfMissingNewParameters(t *testing.T) {
	pds := &v1alpha1.PerfDataSourceGitLab{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
		Spec: v1alpha1.PerfDataSourceGitLabSpec{
			PerfServerName: fakeName,
			Type:           gitlabDsType,
			Config: v1alpha1.DataSourceGitLabConfig{
				Repositories: []string{"repo1"},
				Branches:     []string{"master"},
				Url:          fakeName,
			},
		},
	}

	ps := &v1alpha1.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: v1alpha1.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	sec := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      gitLabSecretName,
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("fake"),
			"password": []byte("fake"),
		},
	}

	cl := fake.NewClientBuilder().
		WithScheme(scheme.Scheme).
		WithObjects(pds, ps, sec).
		Build()

	mPerfCl := new(mock.MockPerfClient)
	ch := PutDataSource{
		client:     cl,
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, gitlabDsType).
		Return(&dto.DataSource{
			Active: true,
			Type:   gitlabDsType,
			Config: map[string]interface{}{
				"repositories": []interface{}{"repo2"},
				"branches":     []interface{}{"develop"},
			},
		}, nil)

	mPerfCl.On("UpdateDataSource", command.DataSourceCommand{
		Type: gitlabDsType,
		Config: command.DataSourceGitlabConfig{
			Repositories:   []string{"repo2", "repo1"},
			Url:            fakeName,
			InstanceId:     fakeName,
			WithMembership: false,
			AllPublic:      false,
			AllBranches:    false,
			Branches:       []string{"develop", "master"},
			Username:       "fake",
			Password:       "fake",
		},
	}).Return(nil)

	assert.NoError(t, ch.ServeRequest(pds))
}
