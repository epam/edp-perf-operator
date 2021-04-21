package chain

import (
	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

const (
	fakeName      = "fake-name"
	fakeNamespace = "fake-namespace"
)

func TestPutOwnerReference_PerfDataSourceContainsPerfServerOwnerReference(t *testing.T) {
	pds := &v1alpha1.PerfDataSourceSonar{
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
	pds := &v1alpha1.PerfDataSourceSonar{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
		Spec: v1alpha1.PerfDataSourceSonarSpec{
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

	cl := fake.NewClientBuilder().
		WithScheme(scheme.Scheme).
		WithObjects(pds, c).
		Build()

	ch := PutOwnerReference{
		scheme: scheme.Scheme,
		client: cl,
	}
	assert.NoError(t, ch.ServeRequest(pds))
}

func TestPutOwnerReference_PerfServerShouldNotBeFound(t *testing.T) {
	pds := &v1alpha1.PerfDataSourceSonar{
		ObjectMeta: v1.ObjectMeta{
			Namespace: fakeNamespace,
		},
		Spec: v1alpha1.PerfDataSourceSonarSpec{
			PerfServerName: fakeName,
		},
	}

	ps := &v1alpha1.PerfServer{}

	ch := PutOwnerReference{
		scheme: scheme.Scheme,
		client: fake.NewFakeClient(pds, ps),
	}
	assert.Error(t, ch.ServeRequest(pds))
}
