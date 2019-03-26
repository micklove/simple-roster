package model

import (
	"fmt"
	"testing"
	"time"
)

//public - can be re-used in other tests
var DefaultShifts Shifts
var DefaultFirstShift *Shift

const displayNamePrefix = "rocket"
const firstNamePrefix = "ronnie"
const lastNamePrefix = "osullivan"

func init() {
	DefaultShifts = make([]Shift, 0)
	DefaultFirstShift = createDefaultShift()
}

func TestShifts_AddShift(t *testing.T) {
	var shifts Shifts = make([]Shift, 0)
	var shiftCount = 10

	for i := 0; i < shiftCount; i++ {
		shift := createDefaultShift()
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
	filterShiftsBySearchName(t, displayNamePrefix, User.HasDisplayName)
}

func TestShifts_FilterByFirstName(t *testing.T) {
	filterShiftsBySearchName(t, "hello", User.HasFirstName)
}

func TestShifts_FilterByLastName(t *testing.T) {
	filterShiftsBySearchName(t, DefaultLastName, User.HasFirstName)
}

func filterShiftsBySearchName(t *testing.T, searchNamePrefix string, f func(User, string) bool) {
	shiftCount := 20
	expectedFilterCount := shiftCount / 2
	shifts := createListOfShifts(t, shiftCount)
	filteredShifts := shifts.Filter(searchNamePrefix, f)
	actualFilteredShiftsLength := len(filteredShifts)

	if actualFilteredShiftsLength != expectedFilterCount {
		t.Errorf("error expected filtered list of size [%v], got [%v]",
			expectedFilterCount, actualFilteredShiftsLength)
	}
}

//Create a list, where half of the list have a different first, last name
func createListOfShifts(t *testing.T, shiftCount int) Shifts {
	var shifts Shifts = make([]Shift, 0)
	for i := 0; i < shiftCount; i++ {
		shift := CreateShift(time.Now(), time.Now().Add(time.Minute*2), *createDefaultUser())
		if i%2 == 0 {
			shift.User.DisplayName = fmt.Sprintf("%v-%v", displayNamePrefix, i)
			shift.User.FirstName = fmt.Sprintf("%v-%v", firstNamePrefix, i)
			shift.User.LastName = fmt.Sprintf("%v-%v", lastNamePrefix, i)
		}
		shifts, _ = shifts.AddShift(shift)
		validateSliceLength(t, shifts, "Shifts", i+1)
	}
	return shifts

}
