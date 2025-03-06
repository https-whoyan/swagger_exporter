package excel

import (
	"strconv"

	"github.com/https-whoyan/swagger_exporter/internal/models"

	"dario.cat/mergo"
)

const defaultSheetName = "OutSwagger"

var (
	falseBool = false
	trueBool  = true
)

func getCell(rowIndex int, column string) string {
	return column + strconv.Itoa(rowIndex)
}

func formatJSON(data []byte) string {
	return string(data)
}

func getSheetName(rows []*models.Row) string {
	if len(rows) == 0 {
		return defaultSheetName
	}
	microserviceName := rows[0].Microservice
	if len(microserviceName) == 0 {
		return defaultSheetName
	}
	return microserviceName
}

func setCeilStyle(cfg *config, cell string, additionStyle int) error {
	currStyle, err := cfg.f.GetCellStyle(cfg.sheetName, cell)
	if err != nil {
		return err
	}
	if currStyle == 0 {
		return cfg.f.SetCellStyle(cfg.sheetName, cell, cell, additionStyle)
	}
	existingStyle, err := cfg.f.GetStyle(currStyle)
	if err != nil {
		return err
	}
	newStyleData, err := cfg.f.GetStyle(additionStyle)
	if err != nil {
		return err
	}
	err = mergo.Merge(existingStyle, newStyleData, mergo.WithOverride)
	if err != nil {
		return err
	}

	outStyle, err := cfg.f.NewStyle(existingStyle)
	if err != nil {
		return err
	}
	return cfg.f.SetCellStyle(cfg.sheetName, cell, cell, outStyle)
}
