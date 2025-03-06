package exporter

import "github.com/https-whoyan/swagger_exporter/internal/models"

type Service interface {
	Export(excelTable *models.ExcelTable) error
}
