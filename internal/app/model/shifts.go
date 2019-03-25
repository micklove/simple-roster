package model

import "errors"

type Shifts []Shift

func (shifts Shifts) AddShift(shift *Shift) (Shifts, error) {
	if shift == nil {
		err := errors.New("cannot add empty shift to Shifts")
		return shifts, err
	}
	return append(shifts, *shift), nil
}
