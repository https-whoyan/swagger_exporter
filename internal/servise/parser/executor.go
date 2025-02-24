package parser

import (
	"github.com/https-whoyan/swagger_exporter/internal/config"
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"os"
)

func GetJsons(cfg *config.Config) ([]*models.JsonInfo, error) {
	osFile, err := os.Open(cfg.JsonFileName)
	if err != nil {
		return nil, err
	}
	defer osFile.Close()
	return parseSwagger(osFile)
}
