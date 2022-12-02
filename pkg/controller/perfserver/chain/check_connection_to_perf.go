package chain

import (
	"context"
	"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/controller/perfserver/chain/handler"
)

type CheckConnectionToPerf struct {
	next       handler.PerfServerHandler
	client     client.Client
	perfClient perf.PerfClient
}

func (h CheckConnectionToPerf) ServeRequest(server *perfApi.PerfServer) error {
	log.Info("start checking connection to PERF", "url", server.Spec.RootUrl)

	connected, err := h.perfClient.Connected()
	if err != nil {
		server.Status.Available = connected
		wrapedErr := fmt.Errorf("failed to connect to PERF instance with %v url: %w", server.Spec.RootUrl, err)
		server.Status.DetailedMessage = wrapedErr.Error()

		return wrapedErr
	}

	server.Status.Available = connected
	server.Status.DetailedMessage = "connected"

	h.updateStatus(server)

	log.Info("connection to PERF has been established", "url", server.Spec.RootUrl)

	return nextServeOrNil(h.next, server)
}

func (h CheckConnectionToPerf) updateStatus(server *perfApi.PerfServer) {
	server.Status.LastTimeUpdated = metaV1.Now()
	if err := h.client.Status().Update(context.TODO(), server); err != nil {
		_ = h.client.Update(context.TODO(), server)
	}
}
