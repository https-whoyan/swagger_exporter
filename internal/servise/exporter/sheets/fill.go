package sheets

import (
	"context"
	"fmt"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"google.golang.org/api/sheets/v4"
)

func (c *sheetsClient) setInfo(ctx context.Context, excelT *models.ExcelTable) error {
	// For each row
	n := len(excelT.ExcelRows)
	for rowID := 1; rowID <= n; rowID++ {
		pathCeil := excelT.Cells.Get(
			models.CeilPosition(fmt.Sprintf("A%d", rowID)),
		)
		if pathCeil == nil {
			continue
		}
		ceilPath := pathCeil.CeilValue.(string)
		method := excelT.Cells.Get(
			models.CeilPosition(fmt.Sprintf("B%d", rowID)),
		).CeilValue.(string)
		parsedData := c.parsedData.Get(ceilPath, method)
		if parsedData == nil {
			continue
		}
		// Replace
		excelT.Cells.ReplaceValue(
			models.CeilPosition(fmt.Sprintf("E%d", rowID)),
			parsedData.AllowedRoles,
		)
	}
	cells := *excelT.Cells
	var values [][]interface{}
	for ceilPosition, ceil := range cells {
		row, col := parsePosition(ceilPosition.String())
		for len(values) <= row {
			values = append(values, []interface{}{})
		}
		for len(values[row]) <= col {
			values[row] = append(values[row], "")
		}
		values[row][col] = ceil.CeilValue
	}
	valueRange := &sheets.ValueRange{
		Range:  c.sheetName + "!A1",
		Values: values,
	}
	_, err := c.sheetsS.Spreadsheets.Values.Update(
		c.sheetID, valueRange.Range, valueRange,
	).ValueInputOption("RAW").Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("ошибка при обновлении значений: %w", err)
	}
	return nil
}
