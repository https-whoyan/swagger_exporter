package excel

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
)

func getCells(cfg *config) (*models.Cells, error) {
	outCells := models.NewCells()
	for rowID := range cfg.rowsN + 2 {
		if rowID == 0 { // Header
			continue
		}
		byRow := models.NewCells()
		for _, columnID := range columnsArr {
			cell := getCell(rowID, columnID)
			// Style
			ceilStyleInt, err := cfg.f.GetCellStyle(cfg.sheetName, cell)
			if err != nil {
				return nil, err
			}
			ceilStyle, err := cfg.f.GetStyle(ceilStyleInt)
			ceilStyle.Font.Size = fontSize
			ceilStyle.Font.Family = fontFamily
			if err != nil {
				return nil, err
			}
			modelsStyle, err := models.FromExelizeCeilStyle(ceilStyle)
			if err != nil {
				return nil, err
			}
			// Value
			ceilValueStr, err := cfg.f.GetCellValue(cfg.sheetName, cell)
			if err != nil {
				return nil, err
			}
			outCells.Add(
				models.NewCeilPosition(cell),
				ceilValueStr,
				modelsStyle,
			)
			byRow.Add(
				models.NewCeilPosition(cell),
				ceilValueStr,
				modelsStyle,
			)
		}
	}
	return &outCells, nil
}
