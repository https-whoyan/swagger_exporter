package models

import (
	"github.com/xuri/excelize/v2"
	"google.golang.org/api/sheets/v4"
)

type Font struct {
	Bold   bool
	Italic bool
	Under  bool
	Family string
	Size   float64
	Strike bool
	Color  *Color
}

func (f Font) SheetsFont() *sheets.TextFormat {
	return &sheets.TextFormat{
		Bold:          f.Bold,
		Italic:        f.Italic,
		Underline:     f.Under,
		FontFamily:    f.Family,
		FontSize:      int64(f.Size),
		Strikethrough: f.Strike,
	}
}

func (f Font) ExelizeFont() *excelize.Font {
	underlineStr := "none"
	if f.Under {
		underlineStr = "single"
	}
	return &excelize.Font{
		Bold:      f.Bold,
		Italic:    f.Italic,
		Underline: underlineStr,
		Family:    f.Family,
		Size:      f.Size,
		Strike:    f.Strike,
	}
}
