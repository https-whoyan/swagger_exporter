package parser

import "github.com/https-whoyan/swagger_exporter/internal/models"

const (
	parametersKey = "parameters"
	requiredKey   = "required"
	queryKey      = "query"
	schemaKey     = "schema"
)

type (
	interfaceSlise          = []interface{}
	mapStringInterfaceSlise = map[string]interface{}
)

func extractQueryParams(detailMap mapStringInterfaceSlise) map[string]models.ParamInfo {
	params := make(map[string]models.ParamInfo)

	parameters, ok := detailMap[parametersKey].(interfaceSlise)
	if !ok {
		return params
	}
	for _, param := range parameters {
		paramMap, paramsOk := param.(mapStringInterfaceSlise)
		if !paramsOk {
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

	return params
}

func extractRequestBody(
	detailMap, definitions mapStringInterfaceSlise,
) *models.SchemaInfo {
	parameters, parametersOk := detailMap[parametersKey].(interfaceSlise)
	if !parametersOk {
		return nil
	}
	for _, param := range parameters {
		paramMap, ok := param.(mapStringInterfaceSlise)
		if !ok {
			continue
		}
		if paramMap["in"] == "body" {
			schema, schemaOk := paramMap[schemaKey].(mapStringInterfaceSlise)
			if !schemaOk {
				continue
			}
			return resolveSchema(schema, definitions)
		}
	}
	return nil
}

func extractResponseBody(
	detailMap mapStringInterfaceSlise, definitions mapStringInterfaceSlise,
) *models.SchemaInfo {
	responses, ok := detailMap["responses"].(mapStringInterfaceSlise)
	if !ok {
		return nil
	}
	var (
		resp   mapStringInterfaceSlise
		schema mapStringInterfaceSlise
	)
	resp, ok = responses["200"].(mapStringInterfaceSlise)
	if !ok {
		return nil
	}
	schema, ok = resp[schemaKey].(mapStringInterfaceSlise)
	if !ok {
		return nil

	}
	return resolveSchema(schema, definitions)
}
