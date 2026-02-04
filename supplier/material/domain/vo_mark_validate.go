package domain

import (
	"errors"
	"time"
)

type MarkValidity struct {
	startDate   time.Time
	expiredDate *time.Time // nil = 永久有效
}

func NewMarkValidity(start time.Time, expired *time.Time) (*MarkValidity, error) {
	if !start.IsZero() && expired != nil {
		if expired.Before(start) {
			return nil, errors.New("mark expiredDate cannot be before startDate")
		}
	}

	return &MarkValidity{
		startDate:   start,
		expiredDate: expired,
	}, nil
}

func (v *MarkValidity) IsExpired(at time.Time) bool {
	if v == nil {
		return true
	}

	if v.expiredDate == nil {
		return false
	}

	return at.After(*v.expiredDate)
}
