package excel

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
)

var (
	columnsArr = mapKeys(headersMap)
	headersMap = map[string]string{
		"A": "Полный путь",
		"B": "HTTP метод",
		"C": "Definition",
		"D": "Микросервис",
		"E": "Разрешенные роли",
		"F": "Query параметры",
		"G": "Body запроса",
		"H": "Response запроса",
	}
)

var widthMap = map[string]float64{
	"A": 45,  // Полный путь
	"B": 17,  // метод
	"C": 43,  // definition
	"D": 19,  // микросервис
	"E": 37,  // разрешенные роли
	"F": 64,  // args
	"G": 81,  // body
	"H": 108, // response
}

var maxColumn = "H"

var rowValueByColumnFns = map[string]func(r *models.Row) interface{}{
	"A": func(r *models.Row) interface{} {
		return r.Path
	},
	"B": func(r *models.Row) interface{} {
		return r.HttpMethod
	},
	"C": func(r *models.Row) interface{} {
		return r.Definition
	},
	"D": func(r *models.Row) interface{} {
		return r.Microservice
	},
	"E": func(r *models.Row) interface{} {
		return r.AllowedRoles
	},
	"F": func(r *models.Row) interface{} {
		return r.QueryParams
	},
	"G": func(r *models.Row) interface{} {
		formattedBody := formatJSON(r.Body)
		return formattedBody
	},
	"H": func(r *models.Row) interface{} {
		formattedResponse := formatJSON(r.Response)
		return formattedResponse
	},
}

func mapKeys[K comparable, V any](mp map[K]V) []K {
	out := make([]K, 0, len(mp))
	for k := range mp {
		out = append(out, k)
	}
	return out
}

func setColumnsWidth(cfg *config) error {
	for _, col := range columnsArr {
		err := cfg.f.SetColWidth(cfg.sheetName, col, col, widthMap[col])
		if err != nil {
			return err
		}
	}
	return nil
}

func getColumnsWidth(cfg *config) ([]*models.ExcelColumn, error) {
	var out []*models.ExcelColumn
	for _, col := range columnsArr {
		colWidth, err := cfg.f.GetColWidth(cfg.sheetName, col)
		if err != nil {
			return nil, err
		}
		out = append(out, &models.ExcelColumn{
			ID:    col,
			Width: int(colWidth),
		})
	}
	return out, nil
}
