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

func (c Cells) ReplaceValue(position CeilPosition, ceilValue interface{}) {
	prev, ok := c[position]
	if !ok {
		c[position] = &Ceil{
			Position:  position,
			CeilValue: ceilValue,
		}
		return
	}
	c[position] = &Ceil{
		Position:  position,
		CeilValue: ceilValue,
		CeilStyle: prev.CeilStyle,
	}
}

func (c Cells) Len() int {
	return len(c)
}

func (c Cells) Values() []interface{} {
	values := make([]interface{}, 0, len(c))
	for _, ceil := range c {
		values = append(values, ceil.CeilValue)
	}
	return values
}

func (c Cells) Get(position CeilPosition) *Ceil {
	return c[position]
}

func NewCells() Cells {
	return make(Cells)
}
