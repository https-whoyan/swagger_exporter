package convertor

import (
	"encoding/json"
	"github.com/https-whoyan/swagger_exporter/internal/config"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"sort"
)

func Convert(cfg *config.Config, jsons []*models.JsonInfo) ([]*models.Row, error) {
	var out = make([]*models.Row, len(jsons))
	microservice := cfg.Microservice
	for i, iJson := range jsons {
		queryArgsJson, err := json.Marshal(iJson.QueryParams)
		if err != nil {
			return nil, err
		}
		bodyJson, err := json.Marshal(iJson.RequestBody)
		if err != nil {
			return nil, err
		}
		responseBodyJson, err := json.Marshal(iJson.ResponseBody)
		if err != nil {
			return nil, err
		}
		if string(queryArgsJson) == "null" {
			queryArgsJson = []byte{'{', '}'}
		}
		if string(bodyJson) == "null" {
			bodyJson = []byte{'{', '}'}
		}
		if string(responseBodyJson) == "null" {
			responseBodyJson = []byte{'{', '}'}
		}
		out[i] = &models.Row{
			Path:         iJson.FullPath,
			HttpMethod:   iJson.Method,
			Microservice: microservice,
			AllowedRoles: "",
			QueryParams:  queryArgsJson,
			Body:         bodyJson,
			Response:     responseBodyJson,
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Path < out[j].Path
	})
	return out, nil
}
