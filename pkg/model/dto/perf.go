package dto

type PerfProject struct {
	Id            int           `json:"id"`
	Name          string        `json:"nodeName"`
	HasDataSource bool          `json:"hasDataSource"`
	Children      []PerfProject `json:"children"`
}

type DataSource struct {
	Id     int                    `json:"id"`
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Active bool                   `json:"active"`
	Config map[string]interface{} `json:"config"`
}

type PerfCredentials struct {
	Username      string
	Password      string
	LuminateToken string
}
