package models

type CeilPosition string

func NewCeilPosition(s string) CeilPosition {
	return CeilPosition(s)
}

func (c CeilPosition) String() string {
	return string(c)
}

func (c CeilPosition) Str() string {
	return string(c)
}

type Ceil struct {
	Position  CeilPosition
	CeilValue interface{}
	CeilStyle *CeilStyle
}

type Cells map[CeilPosition]*Ceil

func (c Cells) Add(position CeilPosition, ceilValue interface{}, ceilStyle *CeilStyle) {
	c[position] = &Ceil{
		Position:  position,
		CeilValue: ceilValue,
		CeilStyle: ceilStyle,
	}
}

func (c Cells) Len() int {
	return len(c)
}

func NewCells() Cells {
	return make(Cells)
}
