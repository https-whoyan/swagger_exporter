package parser

import "github.com/https-whoyan/swagger_exporter/internal/models"

type Parser interface {
	GetJsons(data []byte) ([]*models.JsonInfo, error)
}
