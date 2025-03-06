package models

type Row struct {
	Path         string `json:"path"`
	HttpMethod   string `json:"http_method"`
	Definition   string `json:"definition"`
	Microservice string `json:"microservice"`
	AllowedRoles string `json:"allowed_roles"`
	QueryParams  []byte `json:"query_params"`
	Body         []byte `json:"body"`
	Response     []byte `json:"response"`
}

type Rows struct {
	mp map[string]*Row
}

func NewRows() *Rows {
	return &Rows{mp: make(map[string]*Row)}
}

func (r Rows) Add(row *Row) {
	r.mp[row.Path+":"+row.HttpMethod] = row
}

func (r Rows) Get(path string, method string) *Row {
	return r.mp[path+":"+method]
}

func (r *Rows) Len() int {
	return len(r.mp)
}
