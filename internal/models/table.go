package models

type ExcelRow struct {
	ID    int
	Width float64
}

type ExcelColumn struct {
	ID    string
	Width float64
}

type ExcelTable struct {
	Cells   *Cells
	Rows    []*ExcelRow
	Columns []*ExcelColumn
}

const DefaultSheetName = "Sheet"

var SkippedColumnsSetValue = []int{
	4, // - allowed roles
}

func IsSkippedColumn(column int) bool {
	for _, skippedColumn := range SkippedColumnsSetValue {
		if column == skippedColumn {
			return true
		}
	}
	return false
}
