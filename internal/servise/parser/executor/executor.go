package executor

import (
	"fmt"
	v3 "github.com/https-whoyan/swagger_exporter/internal/servise/parser/v3"
	"io"
	"os"

	"github.com/https-whoyan/swagger_exporter/internal/config"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	v2 "github.com/https-whoyan/swagger_exporter/internal/servise/parser/v2"
)

func GetJsons(cfg config.Config) ([]*models.JsonInfo, error) {
	osFile, err := os.Open(cfg.Common().JsonFileName)
	if err != nil {
		return nil, err
	}
	defer osFile.Close()
	v, err := detectSwaggerMajorVersion(osFile)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(osFile)
	if err != nil {
		return nil, err
	}
	switch v {
	case 2:
		return v2.NewV2Executor().GetJsons(data)
	case 3:
		return v3.NewV3Executor().GetJsons(data)
	}
	return nil, fmt.Errorf("unsupported swagger major version %d", v)
}
