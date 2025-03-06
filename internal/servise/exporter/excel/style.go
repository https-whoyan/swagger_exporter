package excel

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"github.com/xuri/excelize/v2"
)

func (s *excelService) setStyles(f *excelize.File, t *models.ExcelTable) error {
	cells := *t.Cells
	for ceilPosition, ceil := range cells {
		excelizeStyle, err := ceil.CeilStyle.ExcelizeStyle()
		if err != nil {
			return err
		}
		excelizeStyleInt, err := f.NewStyle(excelizeStyle)
		if err != nil {
			return err
		}
		err = f.SetCellStyle(s.sheetName, ceilPosition.Str(), ceilPosition.Str(), excelizeStyleInt)
		if err != nil {
			return err
		}
	}
	return nil
}
