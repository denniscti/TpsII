package domain

import "errors"

type MaterialNo struct {
	value string
}

func NewMaterialNo(v string) (MaterialNo, error) {
	if v == "" {
		return MaterialNo{}, errors.New("materialNo cannot be empty")
	}
	return MaterialNo{value: v}, nil
}

func (m MaterialNo) Value() string {
	return m.value
}
