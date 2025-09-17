package v3

import (
	"strings"

	"github.com/https-whoyan/swagger_exporter/internal/models"
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

// Используем локальное множество посещённых $ref, чтобы избежать бесконечной рекурсии.
func resolveSchema(schema mapStringInterfaceSlice, components mapStringInterfaceSlice) *models.SchemaInfo {
	// Используем локальное множество посещённых $ref, чтобы избежать бесконечной рекурсии.
	return resolveSchemaWithSeen(schema, components, make(map[string]bool))
}

func resolveSchemaWithSeen(schema mapStringInterfaceSlice, components mapStringInterfaceSlice, seen map[string]bool) *models.SchemaInfo {
	// Обработка $ref
	if ref, ok := schema[refKeyV3].(string); ok {
		refName := strings.TrimPrefix(ref, componentsSchemasKeyV3)

		// Если уже раскрывали этот ref в текущем стеке — вернуть ссылку без дальнейшего раскрытия.
		if seen[refName] {
			return &models.SchemaInfo{Ref: refName}
		}

		definition, found := components[refName].(mapStringInterfaceSlice)
		if !found {
			// Нет определения — вернуть как ссылку.
			return &models.SchemaInfo{Ref: refName}
		}

		seen[refName] = true
		defer delete(seen, refName)

		return resolveSchemaWithSeen(definition, components, seen)
	}

	// Тип схемы
	schemaType, _ := schema[typeKeyV3].(string)

	switch schemaType {
	case arrayTypeKeyV3:
		schemaInfo := &models.SchemaInfo{Type: arrayTypeKeyV3}
		if items, found := schema[itemsKeyV3]; found {
			switch typedItems := items.(type) {
			case mapStringInterfaceSlice:
				schemaInfo.Items = resolveSchemaWithSeen(typedItems, components, seen)
			case string:
				// Случай строкового $ref в items
				refName := strings.TrimPrefix(typedItems, componentsSchemasKeyV3)
				if defn, ok := components[refName].(mapStringInterfaceSlice); ok {
					if seen[refName] {
						schemaInfo.Items = &models.SchemaInfo{Ref: refName}
					} else {
						seen[refName] = true
						defer delete(seen, refName)
						schemaInfo.Items = resolveSchemaWithSeen(defn, components, seen)
					}
				} else {
					schemaInfo.Items = &models.SchemaInfo{Ref: refName}
				}
			}
		}
		return schemaInfo

	case objectTypeKeyV3:
		schemaInfo := &models.SchemaInfo{Type: objectTypeKeyV3, Properties: make(map[string]*models.SchemaInfo)}

		// additionalProperties — карта (map) или может быть булевым (true/false). Обрабатываем карту.
		if addProps, found := schema[additionalPropsKeyV3].(mapStringInterfaceSlice); found {
			schemaInfo.Type = mapTypeKeyV3
			schemaInfo.Items = resolveSchemaWithSeen(addProps, components, seen)
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
			schemaInfo.Properties[key] = resolveSchemaWithSeen(propSchema, components, seen)
		}
		return schemaInfo

	default:
		return &models.SchemaInfo{Type: schemaType}
	}
}
