package chain

import (
	"fmt"

	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
	"github.com/epam/edp-perf-operator/v2/controllers/perfserver/chain/handler"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
)

type PutPerfProject struct {
	next       handler.PerfServerHandler
	perfClient perf.PerfClient
}

func (h PutPerfProject) ServeRequest(server *perfApi.PerfServer) error {
	log.Info("put PERF project", "name", server.Spec.ProjectName)

	if err := h.tryToCreatePerfProject(server); err != nil {
		return err
	}

	log.Info("PERF project has been created ", "name", server.Spec.ProjectName)

	return nextServeOrNil(h.next, server)
}

func (h PutPerfProject) tryToCreatePerfProject(ps *perfApi.PerfServer) error {
	exists, err := h.perfClient.ProjectExists(ps.Spec.ProjectName)
	if err != nil {
		return fmt.Errorf("failed to check if project exists: %w", err)
	}

	if exists {
		log.Info("PERF project already exists. skip creating", "name", ps.Spec.ProjectName)

		return nil
	}

	return fmt.Errorf("failed to replicate %v project from UPSA to PERF: %w", ps.Spec.ProjectName, err)
}
