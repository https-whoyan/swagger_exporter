package models

import (
	"errors"
	"image/color"
	"strconv"
	"strings"

	utilsColor "github.com/https-whoyan/swagger_exporter/internal/utils/color"

	"google.golang.org/api/sheets/v4"
)

type Color sheets.Color

func (c Color) GoogleColor() sheets.Color {
	return sheets.Color(c)
}

func ColorFromHex(hexColor string) (*Color, error) {
	hexColor = strings.TrimPrefix(hexColor, "#")
	if len(hexColor) != 6 {
		return nil, errors.New("invalid HEX color format")
	}
	r, err := strconv.ParseUint(hexColor[0:2], 16, 8)
	if err != nil {
		return nil, err
	}
	g, err := strconv.ParseUint(hexColor[2:4], 16, 8)
	if err != nil {
		return nil, err
	}
	b, err := strconv.ParseUint(hexColor[4:6], 16, 8)
	if err != nil {
		return nil, err
	}
	return &Color{
		Red:   float64(r) / 255.0,
		Green: float64(g) / 255.0,
		Blue:  float64(b) / 255.0,
		Alpha: 1.0,
	}, nil
}

func (c Color) ExelizeColor() string {
	imgColor := color.RGBA{
		R: uint8(c.Red * 255),
		G: uint8(c.Green * 255),
		B: uint8(c.Blue * 255),
		A: uint8(c.Alpha * 255),
	}
	return utilsColor.ParseHexColor(imgColor)
}
