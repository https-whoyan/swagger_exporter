package parser

import (
	"encoding/json"
	"fmt"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"io"
	"os"
	"strings"
)

func parseSwagger(file *os.File) (out []*models.JsonInfo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var swagger mapStringInterfaceSlise
	if err = json.Unmarshal(data, &swagger); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	paths, ok := swagger["paths"].(mapStringInterfaceSlise)
	if !ok {
		return nil, fmt.Errorf("отсутствует поле 'paths' в JSON")
	}
	definitions, _ := swagger["definitions"].(mapStringInterfaceSlise)

	var endpoints []*models.JsonInfo
	for path, methods := range paths {
		methodMap, methodsOk := methods.(mapStringInterfaceSlise)
		if !methodsOk {
			continue
		}
		for method, details := range methodMap {
			detailMap, detailsOk := details.(mapStringInterfaceSlise)
			if !detailsOk {
				continue
			}
			definition := safeGetStr(detailMap[descriptionKey])
			queryParams := extractQueryParams(detailMap)
			requestBody := extractRequestBody(detailMap, definitions)
			responseBody := extractResponseBody(detailMap, definitions)
			endpoints = append(endpoints, &models.JsonInfo{
				FullPath:     path,
				Method:       strings.ToUpper(method),
				Definition:   definition,
				QueryParams:  queryParams,
				RequestBody:  requestBody,
				ResponseBody: responseBody,
			})
		}
	}

	return endpoints, nil
}

func safeGetStr(val interface{}) string {
	if val == nil {
		return ""
	}
	str, _ := val.(string)
	return str
}
