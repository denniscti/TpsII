package domain

import "errors"

type SupplierProductNo struct {
	value string
}

func NewSupplierProductNo(v string) (SupplierProductNo, error) {
	if v == "" {
		return SupplierProductNo{}, errors.New("supplierProductNo cannot be empty")
	}
	return SupplierProductNo{value: v}, nil
}

func (s SupplierProductNo) Value() string {
	return s.value
}
