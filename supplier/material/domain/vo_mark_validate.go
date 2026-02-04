package domain

import (
	"errors"
	"time"
)

func NewMarkValidity(start time.Time, expired time.Time) (*Mark, error) {
	if !start.IsZero() && expired.IsZero() {
		if expired.Before(start) {
			return nil, errors.New("mark expiredDate cannot be before startDate")
		}
	}

	return &Mark{
		StartDate:   start,
		ExpiredDate: expired,
	}, nil
}

func (v *Mark) IsExpired(at time.Time) bool {
	if v == nil {
		return true
	}

	if v.ExpiredDate.Before(at) {
		return false
	}

	return at.After(v.ExpiredDate)
}
