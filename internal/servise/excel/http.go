package excel

import (
	"errors"
	"github.com/xuri/excelize/v2"
)

const (
	httpMethodGet    = "GET"
	httpMethodPost   = "POST"
	httpMethodPut    = "PUT"
	httpMethodDelete = "DELETE"
	httpMethodPatch  = "PATCH"
)

var (
	colorHttpMethodMapFill = map[string][]string{
		httpMethodGet:    {"#A8E6A2", "#D4FAD4"},
		httpMethodPost:   {"#F4B183", "#F8D6B4"},
		httpMethodPut:    {"#FFD966", "#FFECB3"},
		httpMethodDelete: {"#F28B82", "#FAD2CF"},
		httpMethodPatch:  {"#B39DDB", "#E6D7FF"},
	}
)

func parseHttpMethodAndSetStyle(cfg *config, cell string, httpMethod string) error {
	intFillStyle, err := getHttpStyle(cfg, httpMethod)
	if err != nil {
		return err
	}
	return setCeilStyle(cfg, cell, intFillStyle)
}

func getHttpStyle(cfg *config, httpMethod string) (int, error) {
	colors, ok := colorHttpMethodMapFill[httpMethod]
	if !ok {
		return 0, errors.New("http method not exist")
	}
	fillStyle := &excelize.Style{
		Fill: excelize.Fill{
			Type:    "gradient",
			Color:   colors,
			Shading: 1,
		},
	}
	intFillStyle, err := cfg.f.NewStyle(fillStyle)
	if err != nil {
		return 0, err
	}
	return intFillStyle, nil
}
