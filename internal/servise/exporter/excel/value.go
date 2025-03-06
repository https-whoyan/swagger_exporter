package excel

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"github.com/xuri/excelize/v2"
)

func (s *excelService) setCellsInfo(f *excelize.File, t *models.ExcelTable) error {
	cells := *t.Cells
	for ceilPosition, ceil := range cells {
		err := f.SetCellValue(s.sheetName, ceilPosition.Str(), ceil.CeilValue)
		if err != nil {
			return err
		}
	}
	return nil
}
