package main

import (
	"log"

	"github.com/https-whoyan/swagger_exporter/internal/config"

	convertorS "github.com/https-whoyan/swagger_exporter/internal/servise/convertor"
	excelTableS "github.com/https-whoyan/swagger_exporter/internal/servise/excel"
	exporterS "github.com/https-whoyan/swagger_exporter/internal/servise/exporter"
	excelExpS "github.com/https-whoyan/swagger_exporter/internal/servise/exporter/excel"
	parserS "github.com/https-whoyan/swagger_exporter/internal/servise/parser"

	cmdFlag "github.com/https-whoyan/swagger_exporter/cmd/flag"
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
		exporterClient = excelExpS.NewExcelService(&localCfg)
	}
	err = exporterClient.Export(excelTable)
	if err != nil {
		log.Fatal(err)
	}
}
