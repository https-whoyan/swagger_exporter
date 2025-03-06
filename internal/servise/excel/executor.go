package excel

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"github.com/xuri/excelize/v2"
	"log"
)

func GetExcelTable(rows []*models.Row) (*models.ExcelTable, error) {
	if len(rows) == 0 {
		return nil, nil
	}
	f := excelize.NewFile()
	sheetName := getSheetName(rows)
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}
	f.SetActiveSheet(index)

	cfg := &config{
		f:         f,
		sheetName: sheetName,
		rows:      rows,
		rowsN:     len(rows),
	}

	for _, fn := range []func(cfg *config) error{
		initStyles,
		formatSheets,
		setTable,
		setAllRowsStyles,
		setColumnsWidth,
		setRowsHeight,
		setHeader,
		imitationTableStyle,
		setCellsInfo,
	} {
		err = fn(cfg)
		if err != nil {
			return nil, err
		}
	}

	// Cells
	ceils, err := getCells(cfg)
	if err != nil {
		return nil, err
	}
	// rows, columns
	excelRows, err := getRowsHeight(cfg)
	if err != nil {
		return nil, err
	}
	excelColumns, err := getColumnsWidth(cfg)
	if err != nil {
		return nil, err
	}
	outTable := &models.ExcelTable{
		Cells:        ceils,
		ExcelRows:    excelRows,
		ExcelColumns: excelColumns,
	}
	log.Printf(
		"swagger_exporter: will be places %d rows and %d unique ceils.",
		len(outTable.ExcelRows),
		outTable.Cells.Len(),
	)
	return outTable, nil
}
