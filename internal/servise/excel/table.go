package excel

import (
	"github.com/xuri/excelize/v2"
	"math/rand/v2"
)

var (
	tableRowsColors = []string{
		"#ecfdf5",
		"#e8fbf0",
		"#e3f9f6",
		"#ddf7f4",
		"#d9f5f2",
		"#d4f3f0",
		"#cff1ed",
	}
	tableRowsN = len(tableRowsColors)
)

func imitationTableStyle(cfg *config) error {
	for i := 2; i <= cfg.rowsN+1; i++ {
		currFieldFillCfg := excelize.Fill{
			Type:    "pattern",
			Color:   []string{getRandomColor()},
			Pattern: 1,
		}
		currStyleCfg, err := cfg.f.NewStyle(&excelize.Style{
			Fill: currFieldFillCfg,
		})
		if err != nil {
			return err
		}

		for _, column := range columnsArr {
			var (
				ceilStyleInt    int
				parsedCeilStyle *excelize.Style
				cell            = getCell(i, column)
			)
			ceilStyleInt, err = cfg.f.GetCellStyle(cfg.sheetName, cell)
			if err != nil {
				return err
			}
			parsedCeilStyle, err = cfg.f.GetStyle(ceilStyleInt)
			if err != nil {
				return err
			}
			if len(parsedCeilStyle.Fill.Color) != 0 {
				continue
			}

			err = setCeilStyle(cfg, cell, currStyleCfg)
			if err != nil {
				return err
			}
		}
		// end format
	}
	return nil
}

var prevRandomIndex int

func getRandomColor() string {
	var currIndex int
	for {
		randIndex := rand.N(tableRowsN)
		if randIndex != prevRandomIndex {
			currIndex = randIndex
			break
		}
	}
	prevRandomIndex = currIndex
	return tableRowsColors[currIndex]
}
