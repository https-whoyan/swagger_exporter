package v2

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"strings"
)

const (
	refKey               = "$ref"
	typeKey              = "type"
	arrayTypeKey         = "array"
	mapTypeKey           = "map"
	objectTypeKey        = "object"
	itemsKey             = "items"
	propertiesKey        = "properties"
	additionalProperties = "additionalProperties"
	descriptionKey       = "description"
	definitionsKey       = "#/definitions/"
)

func resolveSchema(schema mapStringInterfaceSlice, definitions mapStringInterfaceSlice) *models.SchemaInfo {
	if ref, ok := schema[refKey].(string); ok {
		refName := strings.TrimPrefix(ref, definitionsKey)
		definition, found := definitions[refName].(mapStringInterfaceSlice)
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
			switch typedItems := items.(type) {
			case mapStringInterfaceSlice:
				schemaInfo.Items = resolveSchema(typedItems, definitions)
			case string:
				refName := strings.TrimPrefix(typedItems, definitionsKey)
				schemaInfo.Items = resolveSchema(definitions[refName].(mapStringInterfaceSlice), definitions)
			}
		}
		return schemaInfo
	case objectTypeKey:
		schemaInfo := &models.SchemaInfo{Type: objectTypeKey, Properties: make(map[string]*models.SchemaInfo)}
		if addProps, found := schema[additionalProperties].(mapStringInterfaceSlice); found {
			schemaInfo.Type = mapTypeKey
			schemaInfo.Items = resolveSchema(addProps, definitions)
			return schemaInfo
		}
		properties, found := schema[propertiesKey].(mapStringInterfaceSlice)
		if !found {
			return schemaInfo
		}
		for key, value := range properties {
			propSchema, ok := value.(mapStringInterfaceSlice)
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
