package excel

import "github.com/xuri/excelize/v2"

const fontSize = 12
const fontFamily = "Segoe UI"

var (
	defaultFont = &excelize.Font{
		Size:   fontSize,
		Family: fontFamily,
	}
	headerStyle = &excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#499175"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: true,
			Size: fontSize,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}
	borderStyle = &excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Font: defaultFont,
	}
	jsonStyle = &excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical: "top",
		},
		Font: defaultFont,
	}
	centerStyle = &excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: defaultFont,
	}
	commonStyle = &excelize.Style{
		Font: &excelize.Font{
			Family: "Segoe UI",
			Size:   fontSize,
		},
	}
	boldStyle = &excelize.Style{
		Font: &excelize.Font{
			Family: "Segoe UI",
			Size:   fontSize,
			Bold:   true,
		},
	}

	tableStyle = &excelize.Table{
		StyleName:       "TableStyleLight9", // random idk)))
		ShowHeaderRow:   &trueBool,
		ShowFirstColumn: true,
	}
)

var (
	headerStyleInt, borderStyleInt, jsonStyleInt, centerStyleInt, commonStyleInt, boldTextStyle int

	jsonColumns = []string{
		"F", "G", "H",
	}
	centerStylesColumns = []string{
		"A", "B",
	}
	boldStylesColumns = []string{
		"A", "B",
	}
)

func initStyles(cfg *config) (err error) {
	headerStyleInt, err = cfg.f.NewStyle(headerStyle)
	if err != nil {
		return
	}
	borderStyleInt, err = cfg.f.NewStyle(borderStyle)
	if err != nil {
		return
	}
	jsonStyleInt, err = cfg.f.NewStyle(jsonStyle)
	if err != nil {
		return
	}
	centerStyleInt, err = cfg.f.NewStyle(centerStyle)
	if err != nil {
		return
	}
	commonStyleInt, err = cfg.f.NewStyle(commonStyle)
	if err != nil {
		return
	}
	boldTextStyle, err = cfg.f.NewStyle(boldStyle)
	if err != nil {
		return
	}
	return
}

func setAllRowsStyles(cfg *config) error {
	for i := 1; i <= cfg.rowsN+1; i++ {
		for _, column := range columnsArr {
			cell := getCell(i, column)
			err := setCeilStyle(
				cfg,
				cell,
				commonStyleInt,
			)
			if err != nil {
				return err
			}
			// border
			err = setCeilStyle(cfg, cell, borderStyleInt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
