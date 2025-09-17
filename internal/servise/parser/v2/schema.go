package v2

import (
	"strings"

	"github.com/https-whoyan/swagger_exporter/internal/models"
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
	return resolveSchemaWithSeen(schema, definitions, make(map[string]bool))
}

func resolveSchemaWithSeen(schema mapStringInterfaceSlice, definitions mapStringInterfaceSlice, seen map[string]bool) *models.SchemaInfo {
	if ref, ok := schema[refKey].(string); ok {
		refName := strings.TrimPrefix(ref, definitionsKey)
		if seen[refName] {
			return &models.SchemaInfo{Ref: refName}
		}
		definition, found := definitions[refName].(mapStringInterfaceSlice)
		if !found {
			return &models.SchemaInfo{Ref: refName}
		}
		seen[refName] = true
		defer delete(seen, refName)

		return resolveSchemaWithSeen(definition, definitions, seen)
	}

	schemaType, _ := schema[typeKey].(string)
	switch schemaType {
	case arrayTypeKey:
		schemaInfo := &models.SchemaInfo{Type: arrayTypeKey}
		if items, found := schema[itemsKey]; found {
			switch typedItems := items.(type) {
			case mapStringInterfaceSlice:
				schemaInfo.Items = resolveSchemaWithSeen(typedItems, definitions, seen)
			case string:
				// Случай, когда items — строка с $ref-ом
				refName := strings.TrimPrefix(typedItems, definitionsKey)
				if defn, ok := definitions[refName].(mapStringInterfaceSlice); ok {
					if seen[refName] {
						schemaInfo.Items = &models.SchemaInfo{Ref: refName}
					} else {
						seen[refName] = true
						defer delete(seen, refName)
						schemaInfo.Items = resolveSchemaWithSeen(defn, definitions, seen)
					}
				} else {
					schemaInfo.Items = &models.SchemaInfo{Ref: refName}
				}
			}
		}
		return schemaInfo

	case objectTypeKey:
		schemaInfo := &models.SchemaInfo{Type: objectTypeKey, Properties: make(map[string]*models.SchemaInfo)}

		// Случай map: additionalProperties
		if addProps, found := schema[additionalProperties].(mapStringInterfaceSlice); found {
			schemaInfo.Type = mapTypeKey
			schemaInfo.Items = resolveSchemaWithSeen(addProps, definitions, seen)
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
			schemaInfo.Properties[key] = resolveSchemaWithSeen(propSchema, definitions, seen)
		}
		return schemaInfo

	default:
		return &models.SchemaInfo{Type: schemaType}
	}
}

// ... existing code ...
