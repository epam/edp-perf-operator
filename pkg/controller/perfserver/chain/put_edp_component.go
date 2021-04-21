package chain

import (
	"bufio"
	"context"
	"encoding/base64"
	edpCompApi "github.com/epam/edp-component-operator/pkg/apis/v1/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type PutEdpComponent struct {
	client client.Client
	scheme *runtime.Scheme
}

const (
	perfEdpComponentType = "perf"
	perfIconPath         = "/usr/local/configs/img/perf.svg"
)

func (h PutEdpComponent) ServeRequest(server *v1alpha1.PerfServer) error {
	log.Info("start creating EDP component", "name", server.Name)
	if err := h.putEdpComponent(server); err != nil {
		return err
	}
	log.Info("EDP component was created", "name", server.Name)
	return nil
}

func (h PutEdpComponent) putEdpComponent(server *v1alpha1.PerfServer) error {
	comp := &edpCompApi.EDPComponent{}
	err := h.client.Get(context.TODO(), types.NamespacedName{
		Name:      server.Name,
		Namespace: server.Namespace,
	}, comp)
	if err != nil {
		if errors.IsNotFound(err) {
			return h.createEdpComponent(server)
		}
		return err
	}
	log.Info("EDP component already exists. skip creating...", "name", server.Name)
	return nil
}

func (h PutEdpComponent) createEdpComponent(server *v1alpha1.PerfServer) error {
	icon, err := getIcon()
	if err != nil {
		return err
	}

	comp := &edpCompApi.EDPComponent{
		ObjectMeta: metav1.ObjectMeta{
			Name:      server.Name,
			Namespace: server.Namespace,
		},
		Spec: edpCompApi.EDPComponentSpec{
			Type:    perfEdpComponentType,
			Url:     server.Spec.RootUrl,
			Icon:    *icon,
			Visible: true,
		},
	}

	if err := controllerutil.SetControllerReference(server, comp, h.scheme); err != nil {
		return err
	}

	if err := h.client.Create(context.TODO(), comp); err != nil {
		return err
	}
	log.Info("EDP component has been created", "name", server.Name)
	return nil
}

func getIcon() (*string, error) {
	f, err := os.Open(perfIconPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	encoded := base64.StdEncoding.EncodeToString(content)
	return &encoded, nil
}
