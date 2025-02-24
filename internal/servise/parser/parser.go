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

	var swagger map[string]interface{}
	if err = json.Unmarshal(data, &swagger); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	basePath, _ := swagger["basePath"].(string)
	paths, ok := swagger["paths"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("отсутствует поле 'paths' в JSON")
	}
	definitions, _ := swagger["definitions"].(map[string]interface{})

	var endpoints []*models.JsonInfo
	for path, methods := range paths {
		methodMap, ok := methods.(map[string]interface{})
		if !ok {
			continue
		}

		for method, details := range methodMap {
			detailMap, ok := details.(map[string]interface{})
			if !ok {
				continue
			}
			queryParams := extractQueryParams(detailMap)
			requestBody := extractRequestBody(detailMap, definitions)
			responseBody := extractResponseBody(detailMap, definitions)
			endpoints = append(endpoints, &models.JsonInfo{
				FullPath:     basePath + path,
				Method:       strings.ToUpper(method),
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
