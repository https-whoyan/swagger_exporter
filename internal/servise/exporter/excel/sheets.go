package excel

import (
	"bytes"
	configPkg "github.com/https-whoyan/swagger_exporter/internal/config"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"github.com/https-whoyan/swagger_exporter/internal/servise/exporter"
	"github.com/xuri/excelize/v2"
	"os"
)

type excelService struct {
	cfg       *configPkg.LocalConfig
	sheetName string
}

func NewExcelService(cfg *configPkg.LocalConfig) exporter.Service {
	sheetName := models.DefaultSheetName
	if len(cfg.Microservice) != 0 {
		sheetName = cfg.Microservice
	}
	return &excelService{
		cfg:       cfg,
		sheetName: sheetName,
	}
}

func (s *excelService) Export(excelTable *models.ExcelTable) error {
	f := excelize.NewFile()
	defer f.Close()
	if excelTable.Cells.Len() == 0 {
		return s.safeBuffer(&bytes.Buffer{})
	}
	sheetName := models.DefaultSheetName
	if len(s.cfg.Microservice) != 0 {
		sheetName = s.cfg.Microservice
	}
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}
	f.SetActiveSheet(index)
	for _, fn := range []func(f *excelize.File, t *models.ExcelTable) error{
		s.setColumnsWidth,
		s.setRowsHeight,
		s.setStyles,
		s.setCellsInfo,
	} {
		err = fn(f, excelTable)
		if err != nil {
			return err
		}
	}
	buff := bytes.Buffer{}
	_, err = f.WriteTo(&buff)
	if err != nil {
		return err
	}
	return s.safeBuffer(&buff)
}

func (s *excelService) safeBuffer(buff *bytes.Buffer) error {
	err := os.WriteFile(s.cfg.OutputFileName, buff.Bytes(), 0644)
	return err
}
