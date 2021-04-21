package luminate

import (
	"encoding/json"
	"github.com/epam/edp-perf-operator/v2/pkg/util/common"
	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

type LuminateClient interface {
	GetApiToken(clientId, secret string) (*string, error)
}

type LuminateClientAdapter struct {
	client resty.Client
}

var log = ctrl.Log.WithName("luminate_client")

func NewLuminateRestClient(url string) LuminateClientAdapter {
	cl := resty.New().
		SetHostURL(url)
	return LuminateClientAdapter{client: *cl}
}

func (c LuminateClientAdapter) GetApiToken(clientId, secret string) (*string, error) {
	rl := log.WithValues("clientId", clientId)
	rl.Info("getting Luminate API token")

	resp, err := c.client.R().
		SetBasicAuth(clientId, secret).
		Post("/v1/oauth/token")
	if err != nil || resp.IsError() {
		return nil, errors.Wrapf(err, "Couldn't get Luminate API token for %v client.", clientId)
	}

	at := &struct {
		AccessToken string `json:"access_token"`
	}{}

	if err = json.Unmarshal([]byte(resp.String()), at); err != nil {
		return nil, errors.Wrapf(err, "Couldn't parse Luminate API token for %v client.", clientId)
	}
	rl.Info("Luminate API token has been received.")
	return common.GetStringP(at.AccessToken), nil
}
