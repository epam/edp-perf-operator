package chain

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
	"github.com/epam/edp-perf-operator/v2/controllers/perfdatasourcesonar/chain/handler"
	"github.com/epam/edp-perf-operator/v2/pkg/util/cluster"
	"github.com/epam/edp-perf-operator/v2/pkg/util/consts"
)

type PutOwnerReference struct {
	client client.Client
	scheme *runtime.Scheme
	next   handler.PerfDataSourceSonarHandler
}

func (h PutOwnerReference) ServeRequest(ds *perfApi.PerfDataSourceSonar) error {
	log.Info("put owner reference for Sonar data source", keyName, ds.Name)

	if err := h.setPerfOwnerRef(ds); err != nil {
		return err
	}

	log.Info("owner ref for perf Sonar data source has been added", keyName, ds.Name)

	return nextServeOrNil(h.next, ds)
}

func (h PutOwnerReference) setPerfOwnerRef(ds *perfApi.PerfDataSourceSonar) error {
	log.Info("try to set owner ref for perf Sonar data source", keyName, ds.Name)

	if ow := cluster.GetOwnerReference(consts.CodebaseKind, ds.GetOwnerReferences()); ow != nil {
		log.Info("PerfDataSourceSonar already has owner ref", "data source", ds.Name, "owner name", ow.Name)

		return nil
	}

	c, err := cluster.GetCodebase(h.client, ds.Spec.CodebaseName, ds.Namespace)
	if err != nil {
		return fmt.Errorf("couldn't get %v Codebase from cluster: %w", ds.Spec.CodebaseName, err)
	}

	if err = controllerutil.SetControllerReference(c, ds, h.scheme); err != nil {
		return fmt.Errorf("couldn't set owner ref for %v PerfDataSourceSonar: %w", ds.Name, err)
	}

	if err = h.client.Update(context.TODO(), ds); err != nil {
		return fmt.Errorf("an error has been occurred while updating perf Sonar data source's owner %v: %w", ds.Name, err)
	}

	return nil
}
