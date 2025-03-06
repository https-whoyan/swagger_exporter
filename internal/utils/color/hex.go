package color

import (
	"fmt"
	"image/color"
)

func ParseHexColor(hexColor color.Color) string {
	r, g, b, _ := hexColor.RGBA()
	return fmt.Sprintf("#%02X%02X%02X", r>>8, g>>8, b>>8)
}
