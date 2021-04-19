package chain

import (
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfserver/chain/handler"
	"github.com/pkg/errors"
)

type PutPerfProject struct {
	next       handler.PerfServerHandler
	perfClient perf.PerfClient
}

func (h PutPerfProject) ServeRequest(server *v1alpha1.PerfServer) error {
	log.Info("put PERF project", "name", server.Spec.ProjectName)
	if err := h.tryToCreatePerfProject(server); err != nil {
		return err
	}
	log.Info("PERF project has been created ", "name", server.Spec.ProjectName)
	return nextServeOrNil(h.next, server)
}

func (h PutPerfProject) tryToCreatePerfProject(ps *v1alpha1.PerfServer) error {
	exists, err := h.perfClient.ProjectExists(ps.Spec.ProjectName)
	if err != nil {
		return err
	}
	if exists {
		log.Info("PERF project already exists. skip creating", "name", ps.Spec.ProjectName)
		return nil
	}
	return errors.Errorf("%v project wasn't replicated from UPSA to PERF", ps.Spec.ProjectName)
}
