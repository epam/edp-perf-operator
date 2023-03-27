package chain

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
			return h.createEdpComponent(server, perfIconPath, os.ReadFile)
		}

		return fmt.Errorf("failed to get client: %w", err)
	}

	log.Info("EDP component already exists. skip creating...", keyName, server.Name)

	return nil
}

func (h PutEdpComponent) createEdpComponent(server *perfApi.PerfServer, iconPath string, fileOpener func(string) ([]byte, error)) error {
	icon, err := getIcon(iconPath, fileOpener)
	if err != nil {
		return err
	}

	comp := &componentApi.EDPComponent{
		ObjectMeta: metav1.ObjectMeta{
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

func getIcon(path string, fileReaderFunc func(string) ([]byte, error)) (*string, error) {
	content, err := fileReaderFunc(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(content)

	return &encoded, nil
}
