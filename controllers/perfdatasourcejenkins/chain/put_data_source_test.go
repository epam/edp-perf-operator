package chain

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/api/v1"
	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf/mock"
	"github.com/epam/edp-perf-operator/v2/pkg/model/command"
	"github.com/epam/edp-perf-operator/v2/pkg/model/dto"
)

const (
	jenkinsDsType    = "JENKINS"
	fakeCodebaseName = "stub-val"
)

// nolint
func init() {
	utilruntime.Must(perfApi.AddToScheme(scheme.Scheme))
	utilruntime.Must(codebaseApi.AddToScheme(scheme.Scheme))
}

func TestPutDataSource_ShouldUpdateJenkinsDataSourceWithoutActivating(t *testing.T) {
	pds := &perfApi.PerfDataSourceJenkins{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
		Spec: perfApi.PerfDataSourceJenkinsSpec{
			PerfServerName: fakeName,
			Type:           jenkinsDsType,
			Config: perfApi.DataSourceJenkinsConfig{
				JobNames: []string{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName),
					fmt.Sprintf("/%v/%v-Build-%v", fakeCodebaseName, strings.ToUpper(fakeName), fakeCodebaseName)},
				Url: fakeName,
			},
		},
	}

	ps := &perfApi.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: perfApi.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	sec := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      jenkinsDataSourceSecretName,
			Namespace: fakeNamespace,
		},
		Data: map[string][]byte{
			"username": []byte("fake"),
			"password": []byte("fake"),
		},
	}

	mPerfCl := new(mock.MockPerfClient)
	ch := PutDataSource{
		client:     fake.NewClientBuilder().WithRuntimeObjects(pds, ps, sec).Build(),
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, jenkinsDsType).
		Return(&dto.DataSource{
			Active: true,
			Type:   jenkinsDsType,
			Config: map[string]interface{}{
				"jobNames": []interface{}{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName)},
			},
		}, nil)

	mPerfCl.On("UpdateDataSource", command.DataSourceCommand{
		Type: jenkinsDsType,
		Config: command.DataSourceJenkinsConfig{
			JobNames: []string{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName),
				fmt.Sprintf("/%v/%v-Build-%v", fakeCodebaseName, strings.ToUpper(fakeName), fakeCodebaseName)},
			Url:      fakeName,
			Username: "fake",
			Password: "fake",
		},
	}).Return(nil)

	assert.NoError(t, ch.ServeRequest(pds))
	assert.Equal(t, "created", pds.Status.Status)
}

func TestPutDataSource_ShouldUpdateJenkinsDataSourceWithActivating(t *testing.T) {
	pds := &perfApi.PerfDataSourceJenkins{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
		Spec: perfApi.PerfDataSourceJenkinsSpec{
			PerfServerName: fakeName,
			Type:           jenkinsDsType,
			Config: perfApi.DataSourceJenkinsConfig{
				JobNames: []string{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName),
					fmt.Sprintf("/%v/%v-Build-%v", fakeCodebaseName, strings.ToUpper(fakeName), fakeCodebaseName)},
				Url: fakeName,
			},
		},
	}

	ps := &perfApi.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: perfApi.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	sec := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      jenkinsDataSourceSecretName,
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

	mPerfCl.On("GetProjectDataSource", fakeName, jenkinsDsType).
		Return(&dto.DataSource{
			Active: false,
			Type:   jenkinsDsType,
			Config: map[string]interface{}{
				"jobNames": []interface{}{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName)},
			},
		}, nil)

	mPerfCl.On("UpdateDataSource", command.DataSourceCommand{
		Type: jenkinsDsType,
		Config: command.DataSourceJenkinsConfig{
			JobNames: []string{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName),
				fmt.Sprintf("/%v/%v-Build-%v", fakeCodebaseName, strings.ToUpper(fakeName), fakeCodebaseName)},
			Url:      fakeName,
			Username: "fake",
			Password: "fake",
		},
	}).Return(nil)

	mPerfCl.On("ActivateDataSource", fakeName, 0).Return(nil)

	assert.NoError(t, ch.ServeRequest(pds))
	assert.Equal(t, "created", pds.Status.Status)
}

func TestPutDataSource_ShouldCreateJenkinsDataSource(t *testing.T) {
	pds := &perfApi.PerfDataSourceJenkins{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
		Spec: perfApi.PerfDataSourceJenkinsSpec{
			PerfServerName: fakeName,
			Type:           jenkinsDsType,
			Config: perfApi.DataSourceJenkinsConfig{
				JobNames: []string{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName),
					fmt.Sprintf("/%v/%v-Build-%v", fakeCodebaseName, strings.ToUpper(fakeName), fakeCodebaseName)},
				Url: fakeName,
			},
		},
	}

	ps := &perfApi.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: perfApi.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	sec := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      jenkinsDataSourceSecretName,
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
		client:     fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, jenkinsDsType).Return(nil, nil)

	mPerfCl.On("CreateDataSource", fakeName, command.DataSourceCommand{
		Type: jenkinsDsType,
		Config: command.DataSourceJenkinsConfig{
			JobNames: []string{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName),
				fmt.Sprintf("/%v/%v-Build-%v", fakeCodebaseName, strings.ToUpper(fakeName), fakeCodebaseName)},
			Url:      fakeName,
			Username: "fake",
			Password: "fake",
		},
	}).Return(nil)

	assert.NoError(t, ch.ServeRequest(pds))
	assert.Equal(t, "created", pds.Status.Status)
}

func TestPutDataSource_ShouldNotFindDataSourceInPERF(t *testing.T) {
	ps := &perfApi.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: perfApi.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	objs := []runtime.Object{
		ps,
	}

	mPerfCl := new(mock.MockPerfClient)
	ch := PutDataSource{
		client:     fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, "").Return(nil, fmt.Errorf("failed"))

	pds := &perfApi.PerfDataSourceJenkins{
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
	pds := &perfApi.PerfDataSourceJenkins{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
		Spec: perfApi.PerfDataSourceJenkinsSpec{
			Type: jenkinsDsType,
			Config: perfApi.DataSourceJenkinsConfig{
				JobNames: []string{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName),
					fmt.Sprintf("/%v/%v-Build-%v", fakeCodebaseName, strings.ToUpper(fakeName), fakeCodebaseName)},
				Url: fakeName,
			},
		},
	}

	ps := &perfApi.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: perfApi.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	sec := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      jenkinsDataSourceSecretName,
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
		client:     fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, jenkinsDsType).
		Return(&dto.DataSource{
			Active: false,
			Type:   jenkinsDsType,
			Config: map[string]interface{}{
				"jobNames": []interface{}{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName)},
			},
		}, nil)

	mPerfCl.On("UpdateDataSource", command.DataSourceCommand{
		Type: jenkinsDsType,
		Config: command.DataSourceJenkinsConfig{
			JobNames: []string{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName),
				fmt.Sprintf("/%v/%v-Build-%v", fakeCodebaseName, strings.ToUpper(fakeName), fakeCodebaseName)},
			Url:      fakeName,
			Username: "fake",
			Password: "fake",
		},
	}).Return(nil)

	mPerfCl.On("ActivateDataSource", fakeName, 0).Return(fmt.Errorf("failed"))

	assert.Error(t, ch.ServeRequest(pds))
	assert.Equal(t, "error", pds.Status.Status)
}

func TestPutDataSource_ShouldNotUpdateDataSourceBecauseOfMissingNewParameters(t *testing.T) {
	pds := &perfApi.PerfDataSourceJenkins{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "PerfServer",
					Name: fakeName,
				},
			},
		},
		Spec: perfApi.PerfDataSourceJenkinsSpec{
			PerfServerName: fakeName,
			Type:           jenkinsDsType,
			Config: perfApi.DataSourceJenkinsConfig{
				JobNames: []string{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName)},
				Url:      fakeName,
			},
		},
	}

	ps := &perfApi.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: perfApi.PerfServerSpec{
			ProjectName: fakeName,
		},
	}

	sec := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      jenkinsDataSourceSecretName,
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
		client:     fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
		perfClient: mPerfCl,
	}

	mPerfCl.On("GetProjectDataSource", fakeName, jenkinsDsType).
		Return(&dto.DataSource{
			Active: true,
			Type:   jenkinsDsType,
			Config: map[string]interface{}{
				"jobNames": []interface{}{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName)},
			},
		}, nil)

	mPerfCl.On("UpdateDataSource", command.DataSourceCommand{
		Type: jenkinsDsType,
		Config: command.DataSourceJenkinsConfig{
			JobNames: []string{fmt.Sprintf("/%v/%v-Build-%v", fakeName, strings.ToUpper(fakeName), fakeName),
				fmt.Sprintf("/%v/%v-Build-%v", fakeCodebaseName, strings.ToUpper(fakeName), fakeCodebaseName)},
			Url:      fakeName,
			Username: "fake",
			Password: "fake",
		},
	}).Return(nil)

	assert.NoError(t, ch.ServeRequest(pds))
}
