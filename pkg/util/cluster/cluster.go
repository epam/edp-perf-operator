package cluster

import (
	"context"
	codebaseApi "github.com/epmd-edp/codebase-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"
	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetSecret(client client.Client, name, namespace string) (*coreV1.Secret, error) {
	s := &coreV1.Secret{}
	err := client.Get(context.TODO(), types.NamespacedName{
		Name:      name,
		Namespace: namespace,
	}, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func GetOwnerReference(ownerKind string, ors []metav1.OwnerReference) *metav1.OwnerReference {
	if len(ors) == 0 {
		return nil
	}
	for _, o := range ors {
		if o.Kind == ownerKind {
			return &o
		}
	}
	return nil
}

func GetPerfServerCr(c client.Client, name, namespace string) (*v1alpha1.PerfServer, error) {
	ps := &v1alpha1.PerfServer{}
	if err := c.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, ps); err != nil {
		return nil, err
	}
	return ps, nil
}

func GetConfigMap(client client.Client, name, namespace string) (*v1.ConfigMap, error) {
	cm := &v1.ConfigMap{}
	err := client.Get(context.TODO(), types.NamespacedName{
		Name:      name,
		Namespace: namespace,
	}, cm)
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func GetCodebase(client client.Client, name, namespace string) (*codebaseApi.Codebase, error) {
	i := &codebaseApi.Codebase{}
	if err := client.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, i); err != nil {
		return nil, err
	}
	return i, nil
}
