package parser

import (
	"strings"
	
	"github.com/https-whoyan/swagger_exporter/internal/models"
)

const (
	refKey         = "$ref"
	typeKey        = "type"
	descriptionKey = "description"
	arrayTypeKey   = "array"
	itemsKey       = "items"
	propertiesKey  = "properties"
)

func resolveSchema(schema map[string]interface{}, definitions map[string]interface{}) *models.SchemaInfo {
	if ref, ok := schema[refKey].(string); ok {
		refName := strings.TrimPrefix(ref, "#/definitions/")
		if def, ok := definitions[refName].(map[string]interface{}); ok {
			return resolveSchema(def, definitions)
		}
	}

	result := &models.SchemaInfo{
		Type:       safeGetStr(schema[typeKey]),
		Properties: make(map[string]models.SchemaDetail),
	}

	if result.Type == arrayTypeKey {
		if items, ok := schema[itemsKey].(map[string]interface{}); ok {
			result.Properties[itemsKey] = models.SchemaDetail{
				Type:  safeGetStr(items[typeKey]),
				Items: resolveSchema(items, definitions),
			}
		}
	}

	if properties, ok := schema[propertiesKey].(map[string]interface{}); ok {
		for key, value := range properties {
			if propMap, ok := value.(map[string]interface{}); ok {
				result.Properties[key] = models.SchemaDetail{
					Type: safeGetStr(propMap[typeKey]),
				}
			}
		}
	}

	return result
}
