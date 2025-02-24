package models

type Row struct {
	Path         string `json:"path"`
	HttpMethod   string `json:"http_method"`
	Microservice string `json:"microservice"`
	AllowedRoles string `json:"allowed_roles"`
	QueryParams  []byte `json:"query_params"`
	Body         []byte `json:"body"`
	Response     []byte `json:"response"`
}
