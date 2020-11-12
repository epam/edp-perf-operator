package chain

import (
	"context"
	"github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfdatasourcejenkins/chain/handler"
	"github.com/epmd-edp/perf-operator/v2/pkg/util/cluster"
	"github.com/epmd-edp/perf-operator/v2/pkg/util/consts"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type PutOwnerReference struct {
	client client.Client
	scheme *runtime.Scheme
	next   handler.PerfDataSourceJenkinsHandler
}

func (h PutOwnerReference) ServeRequest(ds *v1alpha1.PerfDataSourceJenkins) error {
	log.Info("put owner reference for Jenkins data source", "name", ds.Name)
	if err := h.setPerfOwnerRef(ds); err != nil {
		return err
	}
	log.Info("owner ref for perf Jenkins data source has been added", "name", ds.Name)
	return nextServeOrNil(h.next, ds)
}

func (h PutOwnerReference) setPerfOwnerRef(ds *v1alpha1.PerfDataSourceJenkins) error {
	log.Info("try to set owner ref for perf Jenkins data source", "name", ds.Name)
	if ow := cluster.GetOwnerReference(consts.PerfServerKind, ds.GetOwnerReferences()); ow != nil {
		log.Info("PerfDataSourceJenkins already has owner ref",
			"data source", ds.Name, "owner name", ow.Name)
		return nil
	}

	ps, err := cluster.GetPerfServerCr(h.client, ds.Spec.PerfServerName, ds.Namespace)
	if err != nil {
		return errors.Wrapf(err, "couldn't get %v PerfServer from cluster", ds.Spec.PerfServerName)
	}

	if err := controllerutil.SetControllerReference(ps, ds, h.scheme); err != nil {
		return errors.Wrapf(err, "couldn't set owner ref for %v PerfDataSourceJenkins", ds.Name)
	}

	if err := h.client.Update(context.TODO(), ds); err != nil {
		return errors.Wrapf(err, "an error has been occurred while updating perf Jenkins data source's owner %v", ds.Name)
	}
	return nil
}
