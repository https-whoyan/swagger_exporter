package excel

import (
	"github.com/https-whoyan/swagger_exporter/internal/models"
	"github.com/xuri/excelize/v2"
)

type config struct {
	f         *excelize.File
	sheetName string
	rows      []*models.Row
	rowsN     int
}
