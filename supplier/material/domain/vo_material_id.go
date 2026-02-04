package domain

import "errors"

type MaterialID struct {
	value string
}

func NewMaterialID(v string) (MaterialID, error) {
	if v == "" {
		return MaterialID{}, errors.New("materialID cannot be empty")
	}
	return MaterialID{value: v}, nil
}

func (id MaterialID) Value() string {
	return id.value
}
