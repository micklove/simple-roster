package model

import (
	"reflect"
	"runtime"
	"testing"
)

func init() {
	DefaultShifts = make([]Shift, 0)
	DefaultFirstShift = createDefaultShift(DefaultRosterId)
}

func TestShifts_AddShift(t *testing.T) {
	var shifts Shifts = make([]Shift, 0)
	var shiftCount = 10

	for i := 0; i < shiftCount; i++ {
		shift := createDefaultShift(DefaultRosterId)
		shifts, _ = shifts.AddShift(shift)
		validateSliceLength(t, shifts, "Shifts", i+1)
	}
}

func TestShifts_AddInvalidShift(t *testing.T) {
	var shifts Shifts = make([]Shift, 0)
	if _, err := shifts.AddShift(nil); err == nil {
		t.Errorf("expected error when trying to add nil shift")
	}
}

func TestShifts_FilterByDisplayName(t *testing.T) {
	filterShiftsBySearchName(t, DisplayNamePrefix, User.HasDisplayName)
}

func TestShifts_FilterByFirstName(t *testing.T) {
	filterShiftsBySearchName(t, FirstNamePrefix, User.HasFirstName)
}

func TestShifts_FilterByLastName(t *testing.T) {
	filterShiftsBySearchName(t, LastNamePrefix, User.HasLastName)
}

func filterShiftsBySearchName(t *testing.T, searchNamePrefix string, f func(User, string) bool) {
	shiftCount := 20
	expectedFilterCount := shiftCount / 2
	shifts := createListOfShiftsWithDifferentNames(t, DefaultRosterId, shiftCount)
	filteredShifts := shifts.FilterByUser(searchNamePrefix, f)
	actualFilteredShiftsLength := len(filteredShifts)

	if actualFilteredShiftsLength != expectedFilterCount {
		actualFuncName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		t.Errorf("error filter with predicate [%v], expected filtered list of size [%v], got [%v]",
			actualFuncName, expectedFilterCount, actualFilteredShiftsLength)
	}
}
