package perf

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/resty.v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/epam/edp-perf-operator/v2/pkg/client/luminate"
	"github.com/epam/edp-perf-operator/v2/pkg/model/command"
	"github.com/epam/edp-perf-operator/v2/pkg/model/dto"
	"github.com/epam/edp-perf-operator/v2/pkg/util/cluster"
)

const (
	keyID                = "id"
	keyContentType       = "Content-Type"
	valueApplicationJson = "application/json"
)

type PerfClient interface {
	Connected() (bool, error)
	GetProject(name string) (ds *dto.PerfProject, err error)
	ProjectExists(name string) (bool, error)
	GetProjectDataSource(projectName, dsType string) (*dto.DataSource, error)
	CreateDataSource(projectName string, command command.DataSourceCommand) error
	ActivateDataSource(projectName string, dataSourceId int) error
	UpdateDataSource(command command.DataSourceCommand) error
}

type PerfClientAdapter struct {
	client resty.Client
}

var log = ctrl.Log.WithName("perf_client")

const luminatesecConfigMapName = "luminatesec-conf"

func NewRestClient(url, user, pwd, lumToken string) (*PerfClientAdapter, error) {
	rl := log.WithValues("url", url, "user", user)
	rl.Info("initializing new Perf REST client.")

	token, err := getAuthorizationToken(url, user, pwd, lumToken)
	if err != nil {
		return nil, err
	}

	cl := resty.New().
		SetHostURL(url).
		SetHeader("lum-api-token", lumToken).
		SetAuthToken(token)

	rl.Info("Perf REST client successfully has been created.")

	return &PerfClientAdapter{
		client: *cl,
	}, err
}

func GetPerfCredentials(c client.Client, secretName, namespace string) (*dto.PerfCredentials, error) {
	cm, err := cluster.GetConfigMap(c, luminatesecConfigMapName, namespace)
	if err != nil {
		return nil, fmt.Errorf("failed to get config map: %w", err)
	}

	lumClient := luminate.NewLuminateRestClient(cm.Data["apiUrl"])

	lumSecret, err := cluster.GetSecret(c, cm.Data["credentialName"], namespace)
	if err != nil {
		return nil, fmt.Errorf("failed to get secret: %w", err)
	}

	lumToken, err := lumClient.GetApiToken(string(lumSecret.Data["username"]), string(lumSecret.Data["password"]))
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	s, err := cluster.GetSecret(c, secretName, namespace)
	if err != nil {
		return nil, fmt.Errorf("failed to get secret: %w", err)
	}

	return &dto.PerfCredentials{
		Username:      string(s.Data["username"]),
		Password:      string(s.Data["password"]),
		LuminateToken: *lumToken,
	}, nil
}

func getAuthorizationToken(url, user, pwd, lumApiToken string) (string, error) {
	resp, err := resty.R().
		SetHeaders(map[string]string{
			keyContentType:  "application/x-www-form-urlencoded",
			"accept":        "text/plain",
			"lum-api-token": lumApiToken,
		}).
		SetFormData(map[string]string{
			"username":       user,
			"password":       pwd,
			"useExternalSSO": "false", // weird behaviour. at this moment should be false despite of using lumApiToken
		}).Post(url + "/api/v2/sso/token")
	if err != nil {
		return "", fmt.Errorf("failed to get PERF token for %v user: %w", user, err)
	}

	if resp.IsError() {
		return "", fmt.Errorf("failed to get PERF token for %v user. Status - %v: %w", user, resp.StatusCode(), err)
	}

	return resp.String(), nil
}

func (c *PerfClientAdapter) Connected() (bool, error) {
	log.Info("start checking connection to PERF", "url", c.client.HostURL)

	_, err := c.getProjects()
	if err != nil {
		return false, fmt.Errorf("failed to establish connection with PERF %v: %w", c.client.HostURL, err)
	}

	log.Info("connection to PERF was established.", "url", c.client.HostURL)

	return true, nil
}

func (c *PerfClientAdapter) GetProject(name string) (ds *dto.PerfProject, err error) {
	projects, err := c.getProjects()
	if err != nil {
		return nil, err
	}

	var (
		node     dto.PerfProject
		findNode func(projects []dto.PerfProject, name string)
	)

	findNode = func(projects []dto.PerfProject, name string) {
		for _, child := range projects {
			if strings.EqualFold(child.Name, name) {
				node = child
				break
			}

			findNode(child.Children, name)
		}
	}
	findNode(projects, name)

	return &node, nil
}

func (c *PerfClientAdapter) ProjectExists(name string) (bool, error) {
	log.Info("start checking project for existence", "name", name)

	project, err := c.GetProject(name)
	if err != nil {
		return false, err
	}

	return project != nil, nil
}

func (c *PerfClientAdapter) getProjects() ([]dto.PerfProject, error) {
	var pp []dto.PerfProject

	resp, err := c.client.R().
		SetHeader(keyContentType, valueApplicationJson).
		SetResult(&pp).
		Get("/api/v2/nodes")
	if err != nil {
		return nil, fmt.Errorf("failed to get projects from PERF: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get projects from PERF. Status - %v: %w", resp.StatusCode(), err)
	}

	return pp, nil
}

func (c *PerfClientAdapter) GetProjectDataSource(projectName, dsType string) (*dto.DataSource, error) {
	rlog := log.WithValues("projectName", projectName, "dsType", dsType)
	rlog.Info("start retrieving PERF datasource")

	project, err := c.GetProject(projectName)
	if err != nil {
		return nil, err
	}

	if project == nil {
		return nil, fmt.Errorf("failed to find PERF project %v: %w", projectName, err)
	}

	if !project.HasDataSource {
		rlog.Info("there're no datasources.")

		return nil, nil
	}

	dss, err := c.getProjectDataSources(project.Id)
	if err != nil {
		return nil, err
	}

	for _, ds := range dss {
		if strings.EqualFold(ds.Type, dsType) {
			rlog.Info("datasource has been found in PERF.")
			return &ds, nil
		}
	}

	rlog.Info("datasource has not been found in PERF.")

	return nil, nil
}

func (c *PerfClientAdapter) getProjectDataSources(projectId int) ([]dto.DataSource, error) {
	var ds []dto.DataSource

	resp, err := c.client.R().
		SetHeader(keyContentType, valueApplicationJson).
		SetResult(&ds).
		SetPathParams(map[string]string{
			keyID: strconv.Itoa(projectId),
		}).
		Get("/api/v2/nodes/{id}/datasets/datasources")
	if err != nil {
		return nil, fmt.Errorf("failed to get datasources for %v project: %w", projectId, err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to get datasources %v project. Status - %v: %w",
			projectId, resp.StatusCode(), err)
	}

	return ds, nil
}

func (c *PerfClientAdapter) CreateDataSource(projectName string, dsc command.DataSourceCommand) error {
	rlog := log.WithValues("project name", projectName, "datasource name", dsc.Name)
	rlog.Info("start creating datasource under project")

	project, err := c.GetProject(projectName)
	if err != nil {
		return err
	}

	if project == nil {
		return fmt.Errorf("PERF project %v wasn't found", projectName)
	}

	resp, err := c.client.R().
		SetHeader(keyContentType, valueApplicationJson).
		SetPathParams(map[string]string{
			keyID: strconv.Itoa(project.Id),
		}).
		SetBody(dsc).
		Post("/api/v2/datasources/node/{id}")
	if err != nil {
		return fmt.Errorf("failed to create %v datasource under %v project: %w", dsc.Name, projectName, err)
	}

	if resp.IsError() {
		return fmt.Errorf("failed to create %v datasource under %v project. Status - %v: %w",
			dsc.Name, projectName, resp.StatusCode(), err)
	}

	rlog.Info("datasource has been created.")

	return nil
}

func (c *PerfClientAdapter) ActivateDataSource(projectName string, dataSourceId int) error {
	rlog := log.WithValues("project name", projectName, "datasource id", dataSourceId)
	rlog.Info("try to activate data source")

	resp, err := c.client.R().
		SetHeader(keyContentType, valueApplicationJson).
		SetPathParams(map[string]string{
			keyID: strconv.Itoa(dataSourceId),
		}).
		Put("/api/v2/datasources/{id}/activation")
	if err != nil {
		return fmt.Errorf("failed to activate %v datasource under %v project: %w", dataSourceId, projectName, err)
	}

	if resp.IsError() {
		return fmt.Errorf("failed to activate %v datasource under %v project. Status - %v: %w",
			dataSourceId, projectName, resp.StatusCode(), err)
	}

	rlog.Info("data source has been activated")

	return nil
}

func (c *PerfClientAdapter) UpdateDataSource(dsc command.DataSourceCommand) error {
	log.Info("start updating PERF datasource", "name", dsc.Name)

	resp, err := c.client.R().
		SetHeader(keyContentType, valueApplicationJson).
		SetPathParams(map[string]string{
			keyID: strconv.Itoa(dsc.Id),
		}).
		SetBody(dsc).
		Put("/api/v2/datasources/{id}")
	if err != nil {
		return fmt.Errorf("failed to update %v datasource: %w", dsc.Name, err)
	}

	if resp.IsError() {
		return fmt.Errorf("failed to update %v datasource. Status - %v: %w", dsc.Name, resp.StatusCode(), err)
	}

	log.Info("PERF datasource has been update.", "name", dsc.Name)

	return nil
}
