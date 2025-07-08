package v3

import (
	"encoding/json"
	"fmt"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"github.com/https-whoyan/swagger_exporter/internal/servise/parser"
	"log"
	"strings"
)

type v3Executor struct{}

func NewV3Executor() parser.Parser {
	return &v3Executor{}
}

func (v3Executor) GetJsons(data []byte) ([]*models.JsonInfo, error) {
	var swagger map[string]interface{}
	if err := json.Unmarshal(data, &swagger); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	paths, ok := swagger["paths"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("отсутствует поле 'paths' в JSON")
	}

	components := make(map[string]interface{})
	if c, ok := swagger["components"].(map[string]interface{}); ok {
		if schemas, ok := c["schemas"].(map[string]interface{}); ok {
			components = schemas
		}
	}

	var endpoints []*models.JsonInfo
	for path, methods := range paths {
		methodMap, isMethodMap := methods.(map[string]interface{})
		if !isMethodMap {
			continue
		}
		for method, details := range methodMap {
			detailMap, isDetailMap := details.(map[string]interface{})
			if !isDetailMap {
				continue
			}
			endpoints = append(endpoints, &models.JsonInfo{
				FullPath:     path,
				Method:       strings.ToUpper(method),
				Definition:   safeGetStr(detailMap["description"]),
				QueryParams:  extractQueryParams(detailMap, components),
				RequestBody:  extractRequestBody(detailMap, components),
				ResponseBody: extractResponseBody(detailMap, components),
			})
		}
	}

	log.Printf("swagger_exporter: OpenAPI v3 parsing successful, found endpoints: %d\n", len(endpoints))
	return endpoints, nil
}
