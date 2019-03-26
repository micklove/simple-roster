package model

import (
	"errors"
)

type Shifts []Shift

func (shifts Shifts) AddShift(shift *Shift) (Shifts, error) {
	if shift == nil {
		err := errors.New("cannot add empty shift to Shifts")
		return shifts, err
	}
	return append(shifts, *shift), nil
}

//Filter returns a new slice containing all shifts in the slice that satisfy the User predicate f.
// e.g. first name match, display name match, etc...
// https://gobyexample.com/collection-functions
func Filter(shifts Shifts, f func(user User) bool) Shifts {
	vsf := make([]Shift, 0)
	for _, v := range shifts {
		if f(v.User) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
