package domain

type Unit struct {
	number int
	unit   string
}

type UnitSet struct {
	base   string
	ratios []Unit
}

func (u UnitSet) Convert(from Unit, qty int) int {
	return qty
}
