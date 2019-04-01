package model

import (
	"fmt"
	"testing"
	"time"
)

const DefaultRosterName = "Blah"

var DefaultRoster *Roster
var DefaultTestShift *Shift

func init() {

	DefaultRoster, _ = CreateRoster(DefaultRosterName)
	DefaultTestShift = createDefaultShift(DefaultRoster.ID)
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
	validateSliceLength(t, DefaultRoster.Shifts, "DefaultRoster.Shifts", 1)

	shifts, _ = DefaultRoster.AddShift(DefaultTestShift)
	validateSliceLength(t, shifts, "DefaultRoster.Shifts", 2)
	validateSliceLength(t, DefaultRoster.Shifts, "DefaultRoster.Shifts", 2)
}

func TestCreateShiftsWithRosterId(t *testing.T) {
	roster, _ := CreateRoster("blah")
	shiftCount := 10
	roster.Shifts = CreateShifts(roster.ID, shiftCount)
	ValidateShifts(t, roster.Shifts, roster.ID)
}

// Arrange 	- create multiple shifts, 1 hour apart.
// Act 		- filter by time, next hour
// Assert   - ensure we have 1 matching shift
func TestRoster_Filter_SingleMatchingShift(t *testing.T) {
	roster, _ := CreateRoster("blah")
	duration := time.Hour * 1
	shiftCount := 10
	roster.Shifts = CreateShiftsWithTimeIncrements(shiftCount, roster.ID, time.Now(), duration)
	f := Shift.Between

	for _, shift := range roster.Shifts {
		shift.DumpSummary()
	}
	filtered := roster.FilterByTime(time.Now(), time.Now().Add(duration), f)
	expectedFilteredCount := 1
	if len(filtered) != expectedFilteredCount {
		t.Errorf("expected filter did not match, expected [%v], was [%v]", expectedFilteredCount, len(filtered))
	}
}

// Arrange 	- create multiple shifts, 1 hour apart
// Act 		- filter by time, e.g. next n hours
// Assert   - ensure we have n matching shifts
func TestRoster_Filter_MultipleMatchingShifts(t *testing.T) {
	roster, _ := CreateRoster("blah")
	duration := time.Hour * 1
	shiftCount := 10
	roster.Shifts = CreateShiftsWithTimeIncrements(shiftCount, roster.ID, time.Now(), duration)
	f := Shift.Between

	roster.Dump()

	expectedFilteredCount := 5
	filtered := roster.FilterByTime(time.Now(), time.Now().Add(duration*5), f)
	if len(filtered) != expectedFilteredCount {
		t.Errorf("expected filter did not match, expected [%v], was [%v]", expectedFilteredCount, len(filtered))
	}
}

//create `count` shifts, starting at initialStartTime, ending, afterStart incrementsDuration.
// Each subsequent shift lasts incrementDuration. e.g. 10-11, 11-12, 12-13, etd...
func CreateShiftsWithTimeIncrements(count int, rosterId string, initialStartTime time.Time, incrementsDuration time.Duration) Shifts {
	shifts := CreateShifts(rosterId, count)
	for i, _ := range shifts {
		if i == 0 {
			shifts[i].Start = initialStartTime
		} else {
			shifts[i].Start = shifts[i-1].Start.Add(incrementsDuration)
		}
		//create some notes, (then overwrite the contents) with the time, to make debugging a bit easier
		shifts[i].End = shifts[i].Start.Add(incrementsDuration)
		shifts[i].AddNote(CreateNote("hello"))
		shifts[i].AddNote(CreateNote("world"))
		shifts[i].Notes[0].Note = fmt.Sprintf("S: %v", shifts[i].Start)
		shifts[i].Notes[1].Note = fmt.Sprintf("E: %v", shifts[i].End)
	}
	return shifts
}

//func TestGetAllShifts(t *testing.T) {
//	t.Fail()
//}
//
//func TestGetShiftsByTime(t *testing.T) {
//	t.Fail()
//}
