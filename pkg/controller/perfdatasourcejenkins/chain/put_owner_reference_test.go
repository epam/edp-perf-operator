package chain

import (
	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

const (
	fakeName      = "fake-name"
	fakeNamespace = "fake-namespace"
)

func TestPutOwnerReference_PerfDataSourceContainsPerfServerOwnerReference(t *testing.T) {
	pds := &v1alpha1.PerfDataSourceJenkins{
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
	pds := &v1alpha1.PerfDataSourceJenkins{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: v1alpha1.PerfDataSourceJenkinsSpec{
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
	pds := &v1alpha1.PerfDataSourceJenkins{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
		},
		Spec: v1alpha1.PerfDataSourceJenkinsSpec{
			PerfServerName: fakeName,
		},
	}

	ps := &v1alpha1.PerfServer{}

	objs := []runtime.Object{
		pds, ps,
	}

	ch := PutOwnerReference{
		scheme: scheme.Scheme,
		client: fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
	}
	assert.Error(t, ch.ServeRequest(pds))
}
