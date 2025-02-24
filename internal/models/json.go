package models

type JsonInfo struct {
	FullPath     string               `json:"full_path"`
	Method       string               `json:"method"`
	QueryParams  map[string]ParamInfo `json:"query_params"`
	RequestBody  *SchemaInfo          `json:"request_body"`
	ResponseBody *SchemaInfo          `json:"response_body"`
}

type ParamInfo struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type SchemaInfo struct {
	Type       string                  `json:"type"`
	Properties map[string]SchemaDetail `json:"properties"`
}

type SchemaDetail struct {
	Type  string      `json:"type"`
	Items *SchemaInfo `json:"items,omitempty"`
}
