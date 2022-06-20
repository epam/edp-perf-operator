package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	edpApi "github.com/epam/edp-component-operator/pkg/apis/v1/v1"

	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
)

const (
	fakeNamespace = "fake-namespace"
)

func TestPutEdpComponent_EdpComponentAlreadyExists(t *testing.T) {
	edpComp := &edpApi.EDPComponent{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
	}

	objs := []runtime.Object{
		edpComp,
	}
	s := scheme.Scheme
	s.AddKnownTypes(v1.SchemeGroupVersion, edpComp)

	ch := PutEdpComponent{
		scheme: s,
		client: fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
	}

	psr := &perfApi.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
	}

	assert.NoError(t, ch.ServeRequest(psr))
}

func TestPutEdpComponent_SchemeDoesntContainEdpComponent(t *testing.T) {
	s := scheme.Scheme
	s.AddKnownTypes(v1.SchemeGroupVersion)

	ch := PutEdpComponent{
		scheme: s,
		client: fake.NewClientBuilder().WithRuntimeObjects([]runtime.Object{}...).Build(),
	}

	psr := &perfApi.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
	}

	assert.Error(t, ch.ServeRequest(psr))
}

func TestPutEdpComponent_IconDoesntExist(t *testing.T) {
	edpComp := &edpApi.EDPComponent{}

	psr := &perfApi.PerfServer{
		ObjectMeta: v1.ObjectMeta{
			Name:      fakeName,
			Namespace: fakeNamespace,
		},
	}

	objs := []runtime.Object{
		edpComp, psr,
	}
	s := scheme.Scheme
	s.AddKnownTypes(v1.SchemeGroupVersion, edpComp, psr)

	ch := PutEdpComponent{
		scheme: s,
		client: fake.NewClientBuilder().WithRuntimeObjects(objs...).Build(),
	}

	assert.Error(t, ch.ServeRequest(psr))
}
