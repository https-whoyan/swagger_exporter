package v3

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"strings"
)

const (
	queryKeyV3       = "query"
	inKeyV3          = "in"
	nameKeyV3        = "name"
	requiredKeyV3    = "required"
	schemaKeyV3      = "schema"
	contentKeyV3     = "content"
	jsonMimeV3       = "application/json"
	responsesKeyV3   = "responses"
	requestBodyKeyV3 = "requestBody"
)

func extractQueryParams(operation mapStringInterfaceSlice, components mapStringInterfaceSlice) map[string]models.ParamInfo {
	params := make(map[string]models.ParamInfo)

	parameters, ok := operation["parameters"].([]interface{})
	if !ok {
		return params
	}
	for _, param := range parameters {
		paramMap, isMap := param.(mapStringInterfaceSlice)
		if !isMap {
			continue
		}
		if ref, found := paramMap["$ref"].(string); found {
			refName := strings.TrimPrefix(ref, "#/components/parameters/")
			if def, ok := components[refName].(mapStringInterfaceSlice); ok {
				paramMap = def
			} else {
				continue
			}
		}
		if paramMap[inKeyV3] != queryKeyV3 {
			continue
		}
		name := paramMap[nameKeyV3].(string)
		required, _ := paramMap[requiredKeyV3].(bool)
		schema, ok := paramMap[schemaKeyV3].(mapStringInterfaceSlice)
		var typ string
		if ok {
			typ, _ = schema["type"].(string)
		}
		params[name] = models.ParamInfo{
			Type:        typ,
			Description: safeGetStr(paramMap["description"]),
			Required:    required,
		}
	}
	return params
}

func extractRequestBody(operation, components mapStringInterfaceSlice) *models.SchemaInfo {
	reqBody, ok := operation[requestBodyKeyV3].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}
	if ref, ok := reqBody["$ref"].(string); ok {
		refName := strings.TrimPrefix(ref, "#/components/requestBodies/")
		if resolved, found := components[refName].(mapStringInterfaceSlice); found {
			reqBody = resolved
		} else {
			return nil
		}
	}

	content, ok := reqBody[contentKeyV3].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}

	jsonContent, ok := content[jsonMimeV3].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}

	schema, ok := jsonContent[schemaKeyV3].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}

	return resolveSchema(schema, components)
}

func extractResponseBody(operation, components mapStringInterfaceSlice) *models.SchemaInfo {
	responses, ok := operation[responsesKeyV3].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}
	resp, ok := responses["200"].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}
	content, ok := resp[contentKeyV3].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}
	jsonContent, ok := content[jsonMimeV3].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}
	schema, ok := jsonContent[schemaKeyV3].(mapStringInterfaceSlice)
	if !ok {
		return nil
	}
	return resolveSchema(schema, components)
}
