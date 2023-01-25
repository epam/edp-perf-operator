package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/api/v1"
	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
)

const (
	fakeName      = "fake-name"
	fakeNamespace = "fake-namespace"
)

func TestPutOwnerReference_PerfDataSourceContainsPerfServerOwnerReference(t *testing.T) {
	pds := &perfApi.PerfDataSourceJenkins{
		ObjectMeta: v1.ObjectMeta{
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "Codebase",
				},
			},
		},
	}
	ch := PutOwnerReference{}
	assert.NoError(t, ch.ServeRequest(pds))
}

func TestPutOwnerReference_ShouldSetOwnerReference(t *testing.T) {
	pds := &perfApi.PerfDataSourceJenkins{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: perfApi.PerfDataSourceJenkinsSpec{
			CodebaseName:   fakeName,
			PerfServerName: fakeName,
		},
	}

	c := &codebaseApi.Codebase{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
	}

	objs := []runtime.Object{
		pds, c,
	}

	ch := PutOwnerReference{
		scheme: scheme.Scheme,
		client: fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
	}
	assert.NoError(t, ch.ServeRequest(pds))
}

func TestPutOwnerReference_PerfServerShouldNotBeFound(t *testing.T) {
	pds := &perfApi.PerfDataSourceJenkins{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
		},
		Spec: perfApi.PerfDataSourceJenkinsSpec{
			PerfServerName: fakeName,
		},
	}

	ps := &perfApi.PerfServer{}

	objs := []runtime.Object{
		pds, ps,
	}

	ch := PutOwnerReference{
		scheme: scheme.Scheme,
		client: fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
	}
	assert.Error(t, ch.ServeRequest(pds))
}
