package flag

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/https-whoyan/swagger_exporter/internal/config"
)

const defaultOutXlsxFile = "output.xlsx"

func ParseFlag() (config.Config, error) {
	// Mode
	runModePtr := flag.String("mode", config.RunModeLocal.String(), "run mode")
	var runMode = config.RunModeLocal
	// Common
	jsonPath := flag.String("json", "", "Путь к JSON файлу Swagger")
	microserviceNamePtr := flag.String("microservice", "", "Название микросервиса, где располагается сваггер.")

	// Local
	outputPath := flag.String("output", "swagger.xlsx", "Путь к выходному Excel файлу")
	// Export
	googleSheetsCredsFile := flag.String("creds", "", "Путь к cread'ам google oauth service client'а")
	flag.Parse()

	// Common
	if ptrStrToStr(runModePtr) != "" {
		runModeStrChecked := ptrStrToStr(runModePtr)
		switch runModeStrChecked {
		case config.RunModeLocal.String():
			runMode = config.RunModeLocal
		case config.RunModeGoogleSheets.String():
			runMode = config.RunModeGoogleSheets
		default:
			return nil, errors.New("неправильный run mode")
		}
	}
	if ptrStrToStr(jsonPath) == "" {
		return nil, errors.New("ошибка: Укажите -json")
	}
	err := checkFile(*jsonPath)
	if err != nil {
		return nil, err
	}
	var microservice = ""
	if microserviceNamePtr != nil && *microserviceNamePtr != "" {
		microservice = *microserviceNamePtr
	}
	// Local
	var outputDst = defaultOutXlsxFile
	if outputPath != nil && *outputPath != "" {
		outputDst = *outputPath
	}
	// Sheets
	var sheetCreedsCfg = ptrStrToStr(googleSheetsCredsFile)
	switch runMode {
	case config.RunModeLocal:
		return &config.LocalConfig{
			CommonConfig: config.CommonConfig{
				JsonFileName: *jsonPath,
				Microservice: microservice,
			},
			OutputFileName: outputDst,
		}, nil
	default:
		return &config.GoogleSheetsConfig{
			CommonConfig: config.CommonConfig{
				JsonFileName: *jsonPath,
				Microservice: microservice,
			},
			GoogleSheetsCredsFile: sheetCreedsCfg,
		}, nil
	}
}

func ptrStrToStr(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func checkFile(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file not exists: %s", filePath)
	}
	return nil
}
