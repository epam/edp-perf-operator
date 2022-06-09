package chain

import (
	"context"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfdatasourcegitlab/chain/handler"
	"github.com/epam/edp-perf-operator/v2/pkg/util/cluster"
	"github.com/epam/edp-perf-operator/v2/pkg/util/consts"
)

type PutOwnerReference struct {
	client client.Client
	scheme *runtime.Scheme
	next   handler.PerfDataSourceGitLabHandler
}

func (h PutOwnerReference) ServeRequest(ds *perfApi.PerfDataSourceGitLab) error {
	log.Info("put owner reference for GitLab data source", "name", ds.Name)
	if err := h.setPerfOwnerRef(ds); err != nil {
		return err
	}
	log.Info("owner ref for perf GitLab data source has been added", "name", ds.Name)
	return nextServeOrNil(h.next, ds)
}

func (h PutOwnerReference) setPerfOwnerRef(ds *perfApi.PerfDataSourceGitLab) error {
	log.Info("try to set owner ref for perf GitLab data source", "name", ds.Name)
	if ow := cluster.GetOwnerReference(consts.CodebaseKind, ds.GetOwnerReferences()); ow != nil {
		log.Info("PerfDataSourceGitLab already has owner ref",
			"data source", ds.Name, "owner name", ow.Name)
		return nil
	}

	c, err := cluster.GetCodebase(h.client, ds.Spec.CodebaseName, ds.Namespace)
	if err != nil {
		return errors.Wrapf(err, "couldn't get %v Codebase from cluster", ds.Spec.CodebaseName)
	}

	if err := controllerutil.SetControllerReference(c, ds, h.scheme); err != nil {
		return errors.Wrapf(err, "couldn't set owner ref for %v PerfDataSourceGitLab", ds.Name)
	}

	if err := h.client.Update(context.TODO(), ds); err != nil {
		return errors.Wrapf(err, "an error has been occurred while updating perf GitLab data source's owner %v", ds.Name)
	}
	return nil
}
