package main

import (
	"log"

	"github.com/https-whoyan/swagger_exporter/internal/config"

	cmdFlag "github.com/https-whoyan/swagger_exporter/cmd/flag"
	convertorS "github.com/https-whoyan/swagger_exporter/internal/servise/convertor"
	excelTableS "github.com/https-whoyan/swagger_exporter/internal/servise/excel"
	exporterS "github.com/https-whoyan/swagger_exporter/internal/servise/exporter"
	excelExpC "github.com/https-whoyan/swagger_exporter/internal/servise/exporter/excel"
	sheetsExpC "github.com/https-whoyan/swagger_exporter/internal/servise/exporter/sheets"
	parserS "github.com/https-whoyan/swagger_exporter/internal/servise/parser/executor"
)

func main() {
	cfg, err := cmdFlag.ParseFlag()
	if err != nil {
		log.Fatal(err)
	}

	jsons, err := parserS.GetJsons(cfg)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := convertorS.Convert(cfg, jsons)
	if err != nil {
		log.Fatal(err)
	}

	excelTable, err := excelTableS.GetExcelTable(rows)
	if err != nil {
		log.Fatal(err)
	}

	var exporterClient exporterS.Service
	switch cfg.RunMode() {
	case config.RunModeLocal:
		localCfg, _ := cfg.Config().(config.LocalConfig)
		exporterClient = excelExpC.NewExcelService(&localCfg)
	case config.RunModeGoogleSheets:
		googleSheetsCfg, _ := cfg.Config().(config.GoogleSheetsConfig)
		exporterClient, err = sheetsExpC.NewSheetsClient(&googleSheetsCfg)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown run type: %s", cfg.RunMode())
	}
	err = exporterClient.Export(excelTable)
	if err != nil {
		log.Fatal(err)
	}
}
