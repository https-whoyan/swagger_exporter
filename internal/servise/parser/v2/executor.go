package v2

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/https-whoyan/swagger_exporter/internal/models"
	"github.com/https-whoyan/swagger_exporter/internal/servise/parser"
)

type v2Executor struct{}

func NewV2Executor() parser.Parser {
	return &v2Executor{}
}

func (v2Executor) GetJsons(data []byte) ([]*models.JsonInfo, error) {
	var err error
	var swagger mapStringInterfaceSlice
	if err = json.Unmarshal(data, &swagger); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	paths, ok := swagger["paths"].(mapStringInterfaceSlice)
	if !ok {
		return nil, fmt.Errorf("отсутствует поле 'paths' в JSON")
	}
	definitions, _ := swagger["definitions"].(mapStringInterfaceSlice)

	var endpoints []*models.JsonInfo
	for path, methods := range paths {
		methodMap, methodsOk := methods.(mapStringInterfaceSlice)
		if !methodsOk {
			continue
		}
		for method, details := range methodMap {
			detailMap, detailsOk := details.(mapStringInterfaceSlice)
			if !detailsOk {
				continue
			}
			endpoints = append(endpoints, &models.JsonInfo{
				FullPath:     path,
				Method:       strings.ToUpper(method),
				Definition:   safeGetStr(detailMap[descriptionKey]),
				QueryParams:  extractQueryParams(detailMap, definitions),
				RequestBody:  extractRequestBody(detailMap, definitions),
				ResponseBody: extractResponseBody(detailMap, definitions),
			})
		}
	}

	log.Printf("swagger_exporter: parsing successful, find endpoints N: %d\n", len(endpoints))
	return endpoints, nil
}
