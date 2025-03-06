package sheets

import (
	"context"
	"fmt"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"google.golang.org/api/sheets/v4"
)

func (c *sheetsClient) setStyles(ctx context.Context, excelT *models.ExcelTable) error {
	cells := *excelT.Cells
	var requests []*sheets.Request

	for _, ceil := range cells {
		row, col := parsePosition(ceil.Position.String())
		requests = append(requests, &sheets.Request{
			UpdateCells: &sheets.UpdateCellsRequest{
				Range: &sheets.GridRange{
					SheetId:          c.sheetIDInt,
					StartRowIndex:    int64(row),
					EndRowIndex:      int64(row + 1),
					StartColumnIndex: int64(col),
					EndColumnIndex:   int64(col + 1),
				},
				Rows: []*sheets.RowData{
					{Values: []*sheets.CellData{
						{UserEnteredFormat: ceil.CeilStyle.SheetStyle()},
					}},
				},
				Fields: "userEnteredFormat",
			},
		})
	}
	if len(requests) > 0 {
		err := c.doRequest(ctx, &sheets.BatchUpdateSpreadsheetRequest{
			Requests: requests,
		})
		if err != nil {
			return fmt.Errorf("ошибка при применении стилей: %w", err)
		}
	}
	return nil
}
