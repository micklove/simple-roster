package model

import (
	"testing"
)

//public - can be re-used in other tests
var DefaultShifts Shifts
var DefaultFirstShift *Shift

func init() {
	DefaultShifts = make([]Shift, 0)
	DefaultFirstShift = createDefaultShift()
}

func TestShifts_AddShift(t *testing.T) {
	var shifts Shifts = make([]Shift, 0)
	var shiftCount = 100

	for i := 0; i < shiftCount; i++ {
		shift := createDefaultShift()
		shifts, _ = shifts.AddShift(shift)
		validateSliceLength(t, shifts, "Shifts", i+1)
	}
}
