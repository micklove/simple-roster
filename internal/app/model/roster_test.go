package model

import (
	"testing"
)

const DefaultRosterName = "Blah"

var DefaultRoster *Roster
var DefaultTestShift *Shift

func init() {
	DefaultRoster, _ = CreateRoster(DefaultRosterName)
	DefaultTestShift = createDefaultShift()
}

func TestRoster_CreateRoster(t *testing.T) {
	if DefaultRoster.Shifts == nil {
		t.Error("Expected roster to have empty shifts")
	}
	if DefaultRoster.Name != DefaultRosterName {
		t.Errorf("Expected roster to have name [%v], was [%v]", DefaultRosterName, DefaultRoster.Name)
	}
}

func TestRoster_AddShift(t *testing.T) {
	var shifts, _ = DefaultRoster.AddShift(DefaultTestShift)
	validateSliceLength(t, shifts, "DefaultRoster.Shifts", 1)

	shifts, _ = DefaultRoster.AddShift(DefaultTestShift)
	validateSliceLength(t, shifts, "DefaultRoster.Shifts", 2)
	validateSliceLength(t, DefaultRoster.Shifts, "DefaultRoster.Shifts", 2)
}

//func TestGetAllShifts(t *testing.T) {
//	t.Fail()
//}
//
//func TestGetShiftsByTime(t *testing.T) {
//	t.Fail()
//}
