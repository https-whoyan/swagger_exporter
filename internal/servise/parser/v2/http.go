package v2

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"strings"
)

const (
	parametersKey = "parameters"
	requiredKey   = "required"
	queryKey      = "query"
	schemaKey     = "schema"
)

type (
	mapStringInterfaceSlice = map[string]interface{}
)

func extractQueryParams(detailMap mapStringInterfaceSlice, definitions mapStringInterfaceSlice) map[string]models.ParamInfo {
	params := make(map[string]models.ParamInfo)
	parameters, ok := detailMap[parametersKey].([]interface{})
	if !ok {
		return params
	}
	for _, param := range parameters {
		paramMap, isObject := param.(mapStringInterfaceSlice)
		if !isObject {
			continue
		}
		if ref, found := paramMap[refKey].(string); found {
			refName := strings.TrimPrefix(ref, "#/definitions/")
			if definition, exists := definitions[refName].(mapStringInterfaceSlice); exists {
				paramMap = definition
			} else {
				continue
			}
		}
		if safeGetStr(paramMap["in"]) != queryKey {
			continue
		}
		name := safeGetStr(paramMap["name"])
		required, _ := paramMap[requiredKey].(bool)
		params[name] = models.ParamInfo{
			Type:        safeGetStr(paramMap["type"]),
			Description: safeGetStr(paramMap["description"]),
			Required:    required,
		}
	}
	return params
}

func extractRequestBody(detailMap, definitions mapStringInterfaceSlice) *models.SchemaInfo {
	parameters, ok := detailMap[parametersKey].([]interface{})
	if !ok {
		return nil
	}
	var lastBodyParam mapStringInterfaceSlice
	for _, param := range parameters {
		paramMap, isObject := param.(mapStringInterfaceSlice)
		if !isObject {
			continue
		}
		if safeGetStr(paramMap["in"]) == "body" {
			lastBodyParam = paramMap
		}
	}
	if lastBodyParam == nil {
		return nil
	}
	if ref, exists := lastBodyParam[refKey].(string); exists {
		refName := strings.TrimPrefix(ref, "#/definitions/")
		if definition, found := definitions[refName].(mapStringInterfaceSlice); found {
			return resolveSchema(definition, definitions)
		}
		return nil
	}
	schema, schemaOk := lastBodyParam[schemaKey].(mapStringInterfaceSlice)
	if !schemaOk {
		return nil
	}
	return resolveSchema(schema, definitions)
}

func extractResponseBody(detailMap, definitions mapStringInterfaceSlice) *models.SchemaInfo {
	responses, ok := detailMap["responses"].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}

	resp, ok := responses["200"].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}
	if ref, exists := resp[refKey].(string); exists {
		refName := strings.TrimPrefix(ref, "#/definitions/")
		if definition, found := definitions[refName].(mapStringInterfaceSlice); found {
			return resolveSchema(definition, definitions)
		}
		return nil
	}
	schema, ok := resp[schemaKey].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}

	return resolveSchema(schema, definitions)
}
