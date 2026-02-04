package domain

import "errors"

type Category struct {
	values []string
}

func NewCategory(v []string) (Category, error) {
	if len(v) == 0 {
		return Category{}, errors.New("category cannot be empty")
	}
	return Category{values: v}, nil
}

func (c Category) Values() []string {
	return c.values
}
