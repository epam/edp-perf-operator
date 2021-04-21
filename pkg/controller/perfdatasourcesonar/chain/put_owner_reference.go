package chain

import (
	"context"
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfdatasourcesonar/chain/handler"
	"github.com/epam/edp-perf-operator/v2/pkg/util/cluster"
	"github.com/epam/edp-perf-operator/v2/pkg/util/consts"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type PutOwnerReference struct {
	client client.Client
	scheme *runtime.Scheme
	next   handler.PerfDataSourceSonarHandler
}

func (h PutOwnerReference) ServeRequest(ds *v1alpha1.PerfDataSourceSonar) error {
	log.Info("put owner reference for Sonar data source", "name", ds.Name)
	if err := h.setPerfOwnerRef(ds); err != nil {
		return err
	}
	log.Info("owner ref for perf Sonar data source has been added", "name", ds.Name)
	return nextServeOrNil(h.next, ds)
}

func (h PutOwnerReference) setPerfOwnerRef(ds *v1alpha1.PerfDataSourceSonar) error {
	log.Info("try to set owner ref for perf Sonar data source", "name", ds.Name)
	if ow := cluster.GetOwnerReference(consts.CodebaseKind, ds.GetOwnerReferences()); ow != nil {
		log.Info("PerfDataSourceSonar already has owner ref",
			"data source", ds.Name, "owner name", ow.Name)
		return nil
	}

	c, err := cluster.GetCodebase(h.client, ds.Spec.CodebaseName, ds.Namespace)
	if err != nil {
		return errors.Wrapf(err, "couldn't get %v Codebase from cluster", ds.Spec.CodebaseName)
	}

	if err := controllerutil.SetControllerReference(c, ds, h.scheme); err != nil {
		return errors.Wrapf(err, "couldn't set owner ref for %v PerfDataSourceSonar", ds.Name)
	}

	if err := h.client.Update(context.TODO(), ds); err != nil {
		return errors.Wrapf(err, "an error has been occurred while updating perf Sonar data source's owner %v", ds.Name)
	}
	return nil
}
