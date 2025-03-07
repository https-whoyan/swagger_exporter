package parser

import (
	"encoding/json"
	"fmt"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"io"
	"log"
	"os"
	"strings"
)

func parseSwagger(file *os.File) (out []*models.JsonInfo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic при парсинге: %v", r)
		}
	}()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

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

func safeGetStr(val interface{}) string {
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}
