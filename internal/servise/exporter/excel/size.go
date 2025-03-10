package excel

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"github.com/xuri/excelize/v2"
)

func (s *excelService) setColumnsWidth(f *excelize.File, t *models.ExcelTable) error {
	for _, col := range t.ExcelColumns {
		err := f.SetColWidth(s.sheetName, col.ID, col.ID, col.Width)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *excelService) setRowsHeight(f *excelize.File, t *models.ExcelTable) error {
	for _, col := range t.ExcelRows {
		err := f.SetRowHeight(s.sheetName, col.ID+1, col.Width)
		if err != nil {
			return err
		}
	}
	return nil
}
