package convertor

import (
	"encoding/json"
	"github.com/https-whoyan/swagger_exporter/internal/config"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"sort"
)

func Convert(cfg config.Config, jsons []*models.JsonInfo) ([]*models.Row, error) {
	var out = make([]*models.Row, len(jsons))
	microservice := cfg.Common().Microservice
	for i, iJson := range jsons {
		queryArgsJson, err := marshal(iJson.QueryParams)
		if err != nil {
			return nil, err
		}
		bodyJson, err := marshal(iJson.RequestBody)
		if err != nil {
			return nil, err
		}
		responseBodyJson, err := marshal(iJson.ResponseBody)
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
			Definition:   iJson.Definition,
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

func marshal(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "\t")
}
