package models

type ExcelRow struct {
	ID    int
	Width int
}

type ExcelColumn struct {
	ID    string
	Width int
}

type ExcelTable struct {
	Cells   *Cells
	Rows    []*ExcelRow
	Columns []*ExcelColumn
}

const DefaultSheetName = "Sheet"
