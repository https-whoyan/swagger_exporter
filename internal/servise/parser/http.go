package parser

import "github.com/https-whoyan/swagger_exporter/internal/models"

const (
	parametersKey = "parameters"
	requiredKey   = "required"
	queryKey      = "query"
	schemaKey     = "schema"
)

func extractQueryParams(detailMap map[string]interface{}) map[string]models.ParamInfo {
	params := make(map[string]models.ParamInfo)

	if parameters, ok := detailMap[parametersKey].([]interface{}); ok {
		for _, param := range parameters {
			paramMap, ok := param.(map[string]interface{})
			if !ok {
				continue
			}

			if paramMap["in"] == queryKey {
				name := paramMap["name"].(string)
				required, requiredOk := paramMap[requiredKey].(bool)
				if !requiredOk {
					required = false
				}
				params[name] = models.ParamInfo{
					Type:        safeGetStr(paramMap[typeKey]),
					Description: safeGetStr(paramMap[descriptionKey]),
					Required:    required,
				}
			}
		}
	}

	return params
}

func extractRequestBody(detailMap map[string]interface{}, definitions map[string]interface{}) *models.SchemaInfo {
	if parameters, ok := detailMap[parametersKey].([]interface{}); ok {
		for _, param := range parameters {
			paramMap, ok := param.(map[string]interface{})
			if !ok {
				continue
			}

			if paramMap["in"] == "body" {
				if schema, ok := paramMap[schemaKey].(map[string]interface{}); ok {
					return resolveSchema(schema, definitions)
				}
			}
		}
	}
	return nil
}

func extractResponseBody(detailMap map[string]interface{}, definitions map[string]interface{}) *models.SchemaInfo {
	if responses, ok := detailMap["responses"].(map[string]interface{}); ok {
		if resp, ok := responses["200"].(map[string]interface{}); ok {
			if schema, ok := resp[schemaKey].(map[string]interface{}); ok {
				return resolveSchema(schema, definitions)
			}
		}
	}
	return nil
}
