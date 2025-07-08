package v3

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"strings"
)

const (
	refKeyV3               = "$ref"
	typeKeyV3              = "type"
	arrayTypeKeyV3         = "array"
	mapTypeKeyV3           = "map"
	objectTypeKeyV3        = "object"
	itemsKeyV3             = "items"
	propertiesKeyV3        = "properties"
	additionalPropsKeyV3   = "additionalProperties"
	componentsSchemasKeyV3 = "#/components/schemas/"
)

type mapStringInterfaceSlice = map[string]interface{}

func resolveSchema(schema mapStringInterfaceSlice, components mapStringInterfaceSlice) *models.SchemaInfo {
	if ref, ok := schema[refKeyV3].(string); ok {
		refName := strings.TrimPrefix(ref, componentsSchemasKeyV3)
		if definition, found := components[refName].(mapStringInterfaceSlice); found {
			return resolveSchema(definition, components)
		}
		return &models.SchemaInfo{Ref: refName}
	}

	schemaType, _ := schema[typeKeyV3].(string)

	switch schemaType {
	case arrayTypeKeyV3:
		schemaInfo := &models.SchemaInfo{Type: arrayTypeKeyV3}
		if items, found := schema[itemsKeyV3]; found {
			switch typedItems := items.(type) {
			case mapStringInterfaceSlice:
				schemaInfo.Items = resolveSchema(typedItems, components)
			case string:
				refName := strings.TrimPrefix(typedItems, componentsSchemasKeyV3)
				if resolved, ok := components[refName].(mapStringInterfaceSlice); ok {
					schemaInfo.Items = resolveSchema(resolved, components)
				}
			}
		}
		return schemaInfo

	case objectTypeKeyV3:
		schemaInfo := &models.SchemaInfo{Type: objectTypeKeyV3, Properties: make(map[string]*models.SchemaInfo)}
		if addProps, found := schema[additionalPropsKeyV3].(mapStringInterfaceSlice); found {
			schemaInfo.Type = mapTypeKeyV3
			schemaInfo.Items = resolveSchema(addProps, components)
			return schemaInfo
		}
		properties, found := schema[propertiesKeyV3].(mapStringInterfaceSlice)
		if !found {
			return schemaInfo
		}
		for key, value := range properties {
			propSchema, ok := value.(mapStringInterfaceSlice)
			if !ok {
				continue
			}
			schemaInfo.Properties[key] = resolveSchema(propSchema, components)
		}
		return schemaInfo

	default:
		return &models.SchemaInfo{Type: schemaType}
	}
}
