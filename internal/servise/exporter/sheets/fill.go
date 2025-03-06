package sheets

import (
	"context"
	"fmt"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"google.golang.org/api/sheets/v4"
)

func (c *sheetsClient) setInfo(ctx context.Context, excelT *models.ExcelTable) error {
	cells := *excelT.Cells
	var values [][]interface{}
	resp, err := c.sheetsS.Spreadsheets.Values.Get(c.sheetID, c.sheetName+"!A:Z").Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("ошибка при получении текущих данных: %w", err)
	}
	if len(resp.Values) > 0 {
		values = make([][]interface{}, len(resp.Values))
		for i := range resp.Values {
			values[i] = make([]interface{}, len(resp.Values[i]))
			copy(values[i], resp.Values[i])
		}
	}
	for ceilPosition, ceil := range cells {
		row, col := parsePosition(ceilPosition.String())
		if models.IsSkippedColumn(col) && row != 0 {
			continue
		}
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
	_, err = c.sheetsS.Spreadsheets.Values.Update(
		c.sheetID, valueRange.Range, valueRange,
	).ValueInputOption("RAW").Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("ошибка при обновлении значений: %w", err)
	}
	return nil
}
