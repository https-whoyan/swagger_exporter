package excel

import (
	"fmt"
)

func setHeader(
	cfg *config,
) (err error) {
	for column, header := range headersMap {
		cell := getCell(1, column)
		err = cfg.f.SetCellValue(cfg.sheetName, cell, header)
		if err != nil {
			return err
		}
		err = cfg.f.SetCellStyle(cfg.sheetName, cell, cell, headerStyleInt)
		if err != nil {
			return err
		}
	}
	return nil
}

func setCellsInfo(
	cfg *config,
) (err error) {
	sheetName := cfg.sheetName
	f := cfg.f
	rows := cfg.rows
	for i, row := range rows {
		rowIndex := i + 2

		// HttpMethod
		const httpMethodColumn = "B"
		cell := getCell(rowIndex, httpMethodColumn)
		err = parseHttpMethodAndSetStyle(cfg, cell, row.HttpMethod)
		if err != nil {
			return err
		}

		for _, column := range columnsArr {
			err = f.SetCellValue(
				sheetName,
				getCell(rowIndex, column),
				rowValueByColumnFns[column](row),
			)
			if err != nil {
				return err
			}
		}
		// Special
		for _, jsonColumn := range jsonColumns {
			cell = getCell(rowIndex, jsonColumn)
			err = setCeilStyle(
				cfg,
				cell,
				jsonStyleInt,
			)
			if err != nil {
				return err
			}
		}
		for _, centerColumn := range centerStylesColumns {
			cell = getCell(rowIndex, centerColumn)
			err = setCeilStyle(
				cfg,
				cell,
				centerStyleInt,
			)
			if err != nil {
				return err
			}
		}
		for _, boldStyleColumn := range boldStylesColumns {
			cell = getCell(rowIndex, boldStyleColumn)
			err = setCeilStyle(
				cfg,
				cell,
				boldTextStyle,
			)
		}
	}

	return nil
}

func setTable(
	cfg *config,
) error {
	tableStyle.Name = cfg.sheetName
	tableStyle.Range = fmt.Sprintf("A1:H%d", cfg.rowsN+1)
	err := cfg.f.AddTable(cfg.sheetName, tableStyle)
	if err != nil {
		return err
	}

	return nil
}

func formatSheets(cfg *config) error {
	if cfg.f.SheetCount != 1 {
		return cfg.f.DeleteSheet("Sheet1")
	}
	return nil
}
