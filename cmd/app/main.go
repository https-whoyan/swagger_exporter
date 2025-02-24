package main

import (
	"log"
	"os"

	"github.com/https-whoyan/swagger_exporter/cmd/flag"
	"github.com/https-whoyan/swagger_exporter/internal/servise/convertor"
	"github.com/https-whoyan/swagger_exporter/internal/servise/excel"
	"github.com/https-whoyan/swagger_exporter/internal/servise/parser"
)

func main() {
	cfg, err := flag.ParseFlag()
	if err != nil {
		log.Fatal(err)
	}

	jsons, err := parser.GetJsons(cfg)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := convertor.Convert(cfg, jsons)
	if err != nil {
		log.Fatal(err)
	}

	buff, err := excel.GetBuffer(rows)
	if err != nil {
		log.Fatal(err)
	}

	if err = os.WriteFile(cfg.OutputFileName, buff.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}
