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

func resolveSchema(schema mapStringInterfaceSlise, definitions mapStringInterfaceSlise) *models.SchemaInfo {
	if ref, ok := schema[refKey].(string); ok {
		refName := strings.TrimPrefix(ref, "#/definitions/")
		definition, definitionOk := definitions[refName].(mapStringInterfaceSlise)
		if definitionOk {
			return resolveSchema(definition, definitions)
		}
	}

	result := &models.SchemaInfo{
		Type:       safeGetStr(schema[typeKey]),
		Properties: make(map[string]models.SchemaDetail),
	}
	if result.Type == arrayTypeKey {
		if items, ok := schema[itemsKey].(mapStringInterfaceSlise); ok {
			result.Properties[itemsKey] = models.SchemaDetail{
				Type:  safeGetStr(items[typeKey]),
				Items: resolveSchema(items, definitions),
			}
		}
	}
	if properties, ok := schema[propertiesKey].(mapStringInterfaceSlise); ok {
		for key, value := range properties {
			propMap, propOk := value.(mapStringInterfaceSlise)
			if !propOk {
				continue
			}
			result.Properties[key] = models.SchemaDetail{
				Type: safeGetStr(propMap[typeKey]),
			}
		}
	}

	return result
}
