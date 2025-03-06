package sheets

import (
	"context"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"log"
)

func (c *sheetsClient) parseData(ctx context.Context, _ *models.ExcelTable) error {
	c.parsedData = models.NewRows()
	resp, err := c.sheetsS.Spreadsheets.Values.Get(c.sheetID, c.sheetName+"!A:Z").Context(ctx).Do()
	if err != nil {
		return err
	}
	parsedData := parseData(resp.Values)
	c.parsedData = models.NewRows()
	for i, r := range parsedData {
		if i == 0 {
			// Header
			continue
		}
		var parsedE models.Row
		err = r.Unmarshal(&parsedE)
		if err != nil {
			return err
		}
		c.parsedData.Add(&parsedE)
	}
	log.Printf("swagger_exporter: parsed from sheets data (withoud header) len: %d\n", c.parsedData.Len())
	return nil
}
