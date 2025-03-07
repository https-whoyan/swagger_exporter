package parser

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"strings"
)

const (
	refKey         = "$ref"
	typeKey        = "type"
	arrayTypeKey   = "array"
	objectTypeKey  = "object"
	itemsKey       = "items"
	propertiesKey  = "properties"
	descriptionKey = "description"
	definitionsKey = "#/definitions/"
)

type mapStringInterfaceSlise = map[string]interface{}

func resolveSchema(schema mapStringInterfaceSlise, definitions mapStringInterfaceSlise) *models.SchemaInfo {
	if ref, ok := schema[refKey].(string); ok {
		refName := strings.TrimPrefix(ref, definitionsKey)
		definition, found := definitions[refName].(mapStringInterfaceSlise)
		if found {
			return resolveSchema(definition, definitions)
		}
		return &models.SchemaInfo{Ref: refName}
	}
	schemaType, _ := schema[typeKey].(string)
	switch schemaType {
	case arrayTypeKey:
		schemaInfo := &models.SchemaInfo{Type: arrayTypeKey}
		if items, found := schema[itemsKey]; found {
			switch typpedItems := items.(type) {
			case mapStringInterfaceSlise:
				schemaInfo.Items = resolveSchema(typpedItems, definitions)
			case string:
				refName := strings.TrimPrefix(typpedItems, definitionsKey)
				schemaInfo.Items = resolveSchema(definitions[refName].(mapStringInterfaceSlise), definitions)
			}
		}
		return schemaInfo
	case objectTypeKey:
		schemaInfo := &models.SchemaInfo{Type: objectTypeKey, Properties: make(map[string]*models.SchemaInfo)}
		properties, found := schema[propertiesKey].(mapStringInterfaceSlise)
		if !found {
			return schemaInfo
		}
		for key, value := range properties {
			propSchema, ok := value.(mapStringInterfaceSlise)
			if !ok {
				continue
			}
			schemaInfo.Properties[key] = resolveSchema(propSchema, definitions)
		}
		return schemaInfo
	default:
		return &models.SchemaInfo{Type: schemaType}
	}
}
