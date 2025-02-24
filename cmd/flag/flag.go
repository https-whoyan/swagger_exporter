package flag

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/https-whoyan/swagger_exporter/internal/config"
)

const defaultOutXlsxFile = "output.xlsx"

func ParseFlag() (*config.Config, error) {
	jsonPath := flag.String("json", "", "Путь к JSON файлу Swagger")
	outputPath := flag.String("output", "swagger.xlsx", "Путь к выходному Excel файлу")
	microserviceName := flag.String("microservice", "", "Название микросервиса, где располагается сваггер.")
	flag.Parse()
	if jsonPath == nil || *jsonPath == "" {
		return nil, errors.New("ошибка: Укажите -json")
	}
	if _, err := os.Stat(*jsonPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("файл %s не найден", *jsonPath)
	}
	var outputDst = defaultOutXlsxFile
	if outputPath != nil && *outputPath != "" {
		outputDst = *outputPath
	}
	var microservice = ""
	if microserviceName != nil && *microserviceName != "" {
		microservice = *microserviceName
	}
	return &config.Config{
		JsonFileName:   *jsonPath,
		OutputFileName: outputDst,
		Microservice:   microservice,
	}, nil
}
