package cluster

import (
	"context"
	"fmt"
	"os"
	"strconv"

	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1"

	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
)

const (
	watchNamespaceEnvVar   = "WATCH_NAMESPACE"
	debugModeEnvVar        = "DEBUG_MODE"
	inClusterNamespacePath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
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

func GetOwnerReference(ownerKind string, ors []metaV1.OwnerReference) *metaV1.OwnerReference {
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

func GetPerfServerCr(c client.Client, name, namespace string) (*perfApi.PerfServer, error) {
	ps := &perfApi.PerfServer{}
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

// GetWatchNamespace returns the namespace the operator should be watching for changes
func GetWatchNamespace() (string, error) {
	ns, found := os.LookupEnv(watchNamespaceEnvVar)
	if !found {
		return "", fmt.Errorf("%s must be set", watchNamespaceEnvVar)
	}
	return ns, nil
}

// GetDebugMode returns the debug mode value
func GetDebugMode() (bool, error) {
	mode, found := os.LookupEnv(debugModeEnvVar)
	if !found {
		return false, nil
	}

	b, err := strconv.ParseBool(mode)
	if err != nil {
		return false, err
	}
	return b, nil
}

// Check whether the operator is running in cluster or locally
func RunningInCluster() bool {
	_, err := os.Stat(inClusterNamespacePath)
	return !os.IsNotExist(err)
}
