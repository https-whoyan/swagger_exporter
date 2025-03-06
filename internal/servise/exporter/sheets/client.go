package sheets

import (
	"context"
	"fmt"
	configPkg "github.com/https-whoyan/swagger_exporter/internal/config"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"github.com/https-whoyan/swagger_exporter/internal/servise/oauth"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"os"
)

type Client interface {
	Export(excelTable *models.ExcelTable) error
}
type sheetsClient struct {
	sheetName  string
	sheetIDInt int64
	sheetID    string
	parsedData *models.Rows
	sheetsS    *sheets.Service
}

func NewSheetsClient(cfg *configPkg.GoogleSheetsConfig) (Client, error) {
	ctx := context.Background()
	creedsFile, err := os.Open(cfg.GoogleSheetsCredsFile)
	if err != nil {
		return nil, err
	}
	httpCli, err := oauth.GetHTTPCli(ctx, creedsFile)
	if err != nil {
		return nil, err
	}
	sheetsS, err := sheets.NewService(ctx, option.WithHTTPClient(httpCli))
	if err != nil {
		return nil, err
	}
	sheetsName := models.DefaultSheetName
	if len(cfg.Microservice) != 0 {
		sheetsName = cfg.Microservice
	}
	return &sheetsClient{
		sheetName: sheetsName,
		sheetID:   cfg.SheetID,
		sheetsS:   sheetsS,
	}, nil
}

func (c *sheetsClient) Export(excelTable *models.ExcelTable) error {
	if c.sheetID == "" {
		return fmt.Errorf("sheetID is empty")
	}
	ctx := context.Background()
	for _, fn := range []func(ctx context.Context, excelT *models.ExcelTable) error{
		c.findSheetID,
		c.parseData,
		c.deleteSheet,
		c.createSheet,
		c.findSheetID,
		c.setStyles,
		c.setSizes,
		c.setInfo,
	} {
		err := fn(ctx, excelTable)
		if err != nil {
			return err
		}
	}
	return nil
}
