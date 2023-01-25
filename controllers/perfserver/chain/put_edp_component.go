package chain

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	componentApi "github.com/epam/edp-component-operator/api/v1"
	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
)

type PutEdpComponent struct {
	client client.Client
	scheme *runtime.Scheme
}

const (
	perfEdpComponentType = "perf"
	perfIconPath         = "/usr/local/configs/img/perf.svg"
	keyName              = "name"
)

func (h PutEdpComponent) ServeRequest(server *perfApi.PerfServer) error {
	log.Info("start creating EDP component", keyName, server.Name)

	if err := h.putEdpComponent(server); err != nil {
		return err
	}

	log.Info("EDP component was created", keyName, server.Name)

	return nil
}

func (h PutEdpComponent) putEdpComponent(server *perfApi.PerfServer) error {
	if err := h.client.Get(context.TODO(), types.NamespacedName{
		Name:      server.Name,
		Namespace: server.Namespace,
	}, &componentApi.EDPComponent{}); err != nil {
		if errors.IsNotFound(err) {
			return h.createEdpComponent(server)
		}

		return fmt.Errorf("failed toget client: %w", err)
	}

	log.Info("EDP component already exists. skip creating...", keyName, server.Name)

	return nil
}

func (h PutEdpComponent) createEdpComponent(server *perfApi.PerfServer) error {
	icon, err := getIcon()
	if err != nil {
		return err
	}

	comp := &componentApi.EDPComponent{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      server.Name,
			Namespace: server.Namespace,
		},
		Spec: componentApi.EDPComponentSpec{
			Type:    perfEdpComponentType,
			Url:     server.Spec.RootUrl,
			Icon:    *icon,
			Visible: true,
		},
	}

	if err = controllerutil.SetControllerReference(server, comp, h.scheme); err != nil {
		return fmt.Errorf("failed to set controller reference: %w", err)
	}

	if err = h.client.Create(context.TODO(), comp); err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	log.Info("EDP component has been created", keyName, server.Name)

	return nil
}

func getIcon() (*string, error) {
	f, err := os.Open(perfIconPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer func(f *os.File) {
		if defErr := f.Close(); defErr != nil {
			log.Error(defErr, "failed to close file")
		}
	}(f)

	reader := bufio.NewReader(f)

	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read content: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(content)

	return &encoded, nil
}
