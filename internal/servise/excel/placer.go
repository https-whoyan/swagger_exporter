package excel

import (
	"bytes"
	"strconv"

	"github.com/https-whoyan/swagger_exporter/internal/models"
	"github.com/xuri/excelize/v2"
)

func GetBuffer(rows []*models.Row) (*bytes.Buffer, error) {
	f := excelize.NewFile()
	sheetName := "Swagger Export"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}
	f.SetActiveSheet(index)

	for i, header := range []string{
		"Полный путь",
		"HTTP метод",
		"Микросервис",
		"Разрешенные роли",
		"Query параметры",
		"Body запроса",
		"Response запроса",
	} {
		cell := string(rune('A'+i)) + "1"
		err = f.SetCellValue(sheetName, cell, header)
		if err != nil {
			return nil, err
		}
	}

	for i, row := range rows {
		rowIndex := i + 2
		err = f.SetCellValue(sheetName, getCeil(rowIndex, "A"), row.Path)
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue(sheetName, getCeil(rowIndex, "B"), row.HttpMethod)
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue(sheetName, getCeil(rowIndex, "C"), row.Microservice)
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue(sheetName, getCeil(rowIndex, "D"), row.AllowedRoles)
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue(sheetName, getCeil(rowIndex, "E"), row.QueryParams)
		if err != nil {
			return nil, err
		}
		if len(row.Body) != 0 {
			err = f.SetCellValue(sheetName, getCeil(rowIndex, "F"), row.Body)
			if err != nil {
				return nil, err
			}
		}
		if len(row.Response) != 0 {
			err = f.SetCellValue(sheetName, getCeil(rowIndex, "G"), row.Response)
			if err != nil {
				return nil, err
			}
		}
	}

	var buf bytes.Buffer
	if err = f.Write(&buf); err != nil {
		return nil, err
	}
	return &buf, nil
}

func getCeil(rowIndex int, column string) string {
	return column + strconv.Itoa(rowIndex)
}
