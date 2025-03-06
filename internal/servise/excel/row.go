package excel

import "github.com/https-whoyan/swagger_exporter/internal/models"

const (
	headerRowHeight = 25
	infoRowHeight   = 40
)

func setRowsHeight(cfg *config) error {
	// Header
	err := cfg.f.SetRowHeight(cfg.sheetName, 1, headerRowHeight)
	if err != nil {
		return err
	}

	for i := 1; i <= cfg.rowsN; i++ {
		rowNum := i + 1
		err = cfg.f.SetRowHeight(cfg.sheetName, rowNum, infoRowHeight)
		if err != nil {
			return err
		}
	}
	return nil
}

func getRowsHeight(cfg *config) ([]*models.ExcelRow, error) {
	var out []*models.ExcelRow
	for i := 1; i <= cfg.rowsN+1; i++ {
		rowHeight, err := cfg.f.GetRowHeight(cfg.sheetName, i)
		if err != nil {
			return nil, err
		}
		out = append(out, &models.ExcelRow{
			ID:    i - 1,
			Width: rowHeight,
		})
	}
	return out, nil
}
