package sheets

import (
	"context"
	"fmt"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"google.golang.org/api/sheets/v4"
)

func (c *sheetsClient) deleteSheet(ctx context.Context, _ *models.ExcelTable) error {
	spreadsheet, err := c.sheetsS.Spreadsheets.Get(c.sheetID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("не удалось получить информацию о таблице: %w", err)
	}
	var sheetID int64
	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.Title == c.sheetName {
			sheetID = sheet.Properties.SheetId
			break
		}
	}
	if sheetID == 0 {
		return nil
	}
	return c.doRequest(ctx, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{DeleteSheet: &sheets.DeleteSheetRequest{SheetId: sheetID}},
		},
	})
}

func (c *sheetsClient) createSheet(ctx context.Context, _ *models.ExcelTable) error {
	return c.doRequest(ctx, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{AddSheet: &sheets.AddSheetRequest{
				Properties: &sheets.SheetProperties{Title: c.sheetName},
			}},
		},
	})
}

func (c *sheetsClient) findSheetID(ctx context.Context, _ *models.ExcelTable) error {
	spreadsheet, err := c.sheetsS.Spreadsheets.Get(c.sheetID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("не удалось получить информацию о таблице: %w", err)
	}
	var sheetID int64
	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.Title == c.sheetName {
			sheetID = sheet.Properties.SheetId
			break
		}
	}
	c.sheetIDInt = sheetID
	if sheetID == 0 {
		// create new
		err = c.createSheet(ctx, nil)
		if err != nil {
			return err
		}
		// find
		err = c.findSheetID(ctx, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
