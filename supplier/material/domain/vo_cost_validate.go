package domain

import (
	"errors"
	"time"
)

type CostValidity struct {
	effectiveFrom time.Time
	expiresAt     *time.Time // nil = 永久有效
}

func NewCostValidity(from time.Time, expiresAt *time.Time) (*CostValidity, error) {
	if !from.IsZero() && expiresAt != nil {
		if expiresAt.Before(from) {
			return nil, errors.New("cost expiresAt cannot be before effectiveFrom")
		}
	}

	return &CostValidity{
		effectiveFrom: from,
		expiresAt:     expiresAt,
	}, nil
}

func (v *CostValidity) IsExpired(at time.Time) bool {
	if v == nil {
		return true
	}

	if v.expiresAt == nil {
		return false
	}

	return at.After(*v.expiresAt)
}
