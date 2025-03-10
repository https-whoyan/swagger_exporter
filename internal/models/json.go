package models

//go:generate easyjson -all

type JsonInfo struct {
	FullPath     string               `json:"full_path"`
	Definition   string               `json:"definition"`
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
	Type       string                 `json:"type,omitempty"`       // object, array, string и т. д.
	Properties map[string]*SchemaInfo `json:"properties,omitempty"` // Если это объект
	Items      *SchemaInfo            `json:"items,omitempty"`      // Если это массив
	Ref        string                 `json:"$ref,omitempty"`       // Если это ссылка на другую структуру
}
