package sheets

import (
	"context"
	"fmt"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"google.golang.org/api/sheets/v4"
)

const (
	widthCoeficient  = 5
	heightCoeficient = 2.1
)

func (c *sheetsClient) setSizes(ctx context.Context, excelT *models.ExcelTable) error {
	spreadSheet, err := c.sheetsS.Spreadsheets.Get(c.sheetID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("не удалось получить информацию о таблице: %w", err)
	}

	var sheetID int64
	for _, sheet := range spreadSheet.Sheets {
		if sheet.Properties.Title == c.sheetName {
			sheetID = sheet.Properties.SheetId
			break
		}
	}
	if sheetID == 0 {
		return fmt.Errorf("лист %s не найден", c.sheetName)
	}
	var requests []*sheets.Request

	for _, col := range excelT.Columns {
		colIndex := columnLetterToIndex(col.ID)
		requests = append(requests, &sheets.Request{
			UpdateDimensionProperties: &sheets.UpdateDimensionPropertiesRequest{
				Range: &sheets.DimensionRange{
					SheetId:    sheetID,
					Dimension:  "COLUMNS",
					StartIndex: colIndex,
					EndIndex:   colIndex + 1,
				},
				Properties: &sheets.DimensionProperties{
					PixelSize: int64(col.Width) * widthCoeficient,
				},
				Fields: "pixelSize",
			},
		})
	}
	for _, row := range excelT.Rows {
		requests = append(requests, &sheets.Request{
			UpdateDimensionProperties: &sheets.UpdateDimensionPropertiesRequest{
				Range: &sheets.DimensionRange{
					SheetId:    sheetID,
					Dimension:  "ROWS",
					StartIndex: int64(row.ID),
					EndIndex:   int64(row.ID + 1),
				},
				Properties: &sheets.DimensionProperties{
					PixelSize: int64(row.Width * heightCoeficient),
				},
				Fields: "pixelSize",
			},
		})
	}

	if len(requests) == 0 {
		return nil
	}
	return c.doRequest(ctx, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	})
}
