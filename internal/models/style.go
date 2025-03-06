package models

import (
	"github.com/https-whoyan/swagger_exporter/internal/utils/slice"
	"github.com/xuri/excelize/v2"
	"google.golang.org/api/sheets/v4"
	"strings"
)

type CeilStyle struct {
	BackgroundFill *Fill      `json:"backgroundFill,omitempty"`
	Borders        *Borders   `json:"borders,omitempty"`
	Alignment      *Alignment `json:"aligment,omitempty"`
	Font           *Font      `json:"font,omitempty"`
}

func FromExelizeCeilStyle(exelizeStyle *excelize.Style) (*CeilStyle, error) {
	var (
		outFill      *Fill
		outBorders   *Borders
		outAlignment *Alignment
		outFont      *Font
	)
	if len(exelizeStyle.Fill.Color) != 0 {
		fillColor, err := ColorFromHex(exelizeStyle.Fill.Color[0])
		if err != nil {
			return nil, err
		}
		outFill = &Fill{
			Color: fillColor,
		}
	}
	if len(exelizeStyle.Border) != 0 {
		notEmptyBorder := slice.FindNotEmptyValue(exelizeStyle.Border)
		borderColor, err := ColorFromHex(notEmptyBorder.Color)
		if err != nil {
			return nil, err
		}
		sheetsColor := borderColor.GoogleColor()
		outBorders = &Borders{
			Border: &Border{
				Color: &sheetsColor,
			},
		}
	}
	if exelizeStyle.Alignment != nil {
		wrapTextStr := "OVERFLOW_CELL"
		if exelizeStyle.Alignment.WrapText {
			wrapTextStr = "WRAP"
		}
		outAlignment = &Alignment{
			Horizontal: exelizeStyle.Alignment.Horizontal,
			Vertical:   exelizeStyle.Alignment.Vertical,
			WrapText:   wrapTextStr,
		}
	}
	if exelizeStyle.Font != nil {
		exelizeFont := exelizeStyle.Font
		color := exelizeFont.Color
		if len(color) == 0 {
			color = "#000000"
		}
		fontColor, err := ColorFromHex(color)
		if err != nil {
			return nil, err
		}
		underBool := false
		if len(exelizeFont.Underline) != 0 && exelizeFont.Underline != "none" {
			underBool = true
		}
		outFont = &Font{
			Bold:   exelizeFont.Bold,
			Italic: exelizeFont.Italic,
			Under:  underBool,
			Family: exelizeFont.Family,
			Size:   exelizeFont.Size,
			Strike: exelizeFont.Strike,
			Color:  fontColor,
		}
	}
	return &CeilStyle{
		BackgroundFill: outFill,
		Borders:        outBorders,
		Alignment:      outAlignment,
		Font:           outFont,
	}, nil
}

func (s *CeilStyle) SheetStyle() *sheets.CellFormat {
	var (
		backgroundColor *sheets.Color
		borders         *sheets.Borders
		textFormat      *sheets.TextFormat
	)
	if s.BackgroundFill != nil {
		sheetsColor := s.BackgroundFill.Color.GoogleColor()
		backgroundColor = &sheetsColor
	}
	if s.Borders != nil {
		border := s.Borders.Border
		border.Style = "SOLID_MEDIUM"
		borders = &sheets.Borders{
			Bottom: border,
			Left:   border,
			Right:  border,
			Top:    border,
		}
	}
	if s.Font != nil {
		textFormat = s.Font.SheetsFont()
	}
	var (
		horizontalAlignment string
		verticalAlignment   string
		wrapStrategy        string
	)
	if s.Alignment != nil {
		horizontalAlignment = s.Alignment.Horizontal
		verticalAlignment = s.Alignment.Vertical
		if verticalAlignment == "center" {
			verticalAlignment = "MIDDLE"
		}
		wrapStrategy = s.Alignment.WrapText
	}
	return &sheets.CellFormat{
		BackgroundColor:     backgroundColor,
		Borders:             borders,
		TextFormat:          textFormat,
		HorizontalAlignment: horizontalAlignment,
		VerticalAlignment:   verticalAlignment,
		WrapStrategy:        wrapStrategy,
	}
}

func (s *CeilStyle) ExcelizeStyle() (*excelize.Style, error) {
	var (
		font     *excelize.Font
		fill     excelize.Fill
		aligment *excelize.Alignment
	)
	out := &excelize.Style{}
	if s.Borders != nil {
		sheetsBorderColor := s.Borders.Border.Color
		color := Color{
			Alpha: sheetsBorderColor.Alpha,
			Blue:  sheetsBorderColor.Blue,
			Green: sheetsBorderColor.Green,
			Red:   sheetsBorderColor.Red,
		}
		colorHex := color.ExelizeColor()
		for _, dir := range []string{
			"left", "right", "top", "bottom",
		} {
			s.addExcelizeBorder(out, dir, colorHex)
		}
	}
	if s.Font != nil {
		font = s.Font.ExelizeFont()
	}
	if s.BackgroundFill != nil {
		fill = excelize.Fill{
			Type:    "pattern",
			Color:   []string{s.BackgroundFill.Color.ExelizeColor()},
			Pattern: 1,
		}
	}
	if s.Alignment != nil {
		aligment = &excelize.Alignment{
			Horizontal: s.Alignment.Horizontal,
			Vertical:   s.Alignment.Vertical,
			WrapText:   s.Alignment.WrapText == "WRAP",
		}
	}
	out.Font = font
	out.Fill = fill
	out.Alignment = aligment
	return out, nil
}

func (s *CeilStyle) addExcelizeBorder(out *excelize.Style, dir string, color string) {
	color = strings.TrimPrefix(color, "#")
	out.Border = append(out.Border, excelize.Border{
		Type:  dir,
		Color: color,
		Style: 1,
	})
}
