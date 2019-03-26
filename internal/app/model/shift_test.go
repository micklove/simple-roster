package model

import (
	"testing"
	"time"
)

//public for use in other tests
var DefaultStartTime time.Time
var DefaultEndTime time.Time
var DefaultShift *Shift

const expectedMinutesBetween = 60

func init() {
	DefaultStartTime = time.Now().Add(time.Hour)
	DefaultEndTime = DefaultStartTime.Add(time.Minute * expectedMinutesBetween)
	DefaultShift = createDefaultShift()
}

func TestCreateShift(t *testing.T) {
	t.Logf("shift.start, [%v]", DefaultShift.Start)
	t.Logf("DefaultShift.end ,  [%v]", DefaultShift.End)
	validateDefaultUser(t, &DefaultShift.User, 0)

	if duration, err := DefaultShift.GetShiftDuration(); err != nil {
		t.Errorf("Unexpected error getting duration for shift, Start [%v], End [%v]", DefaultShift.Start, DefaultShift.End)
	} else {
		if duration.Minutes() != expectedMinutesBetween {
			t.Errorf("expected DefaultShift start and end to be %d minutes apart, was [%v]",
				expectedMinutesBetween, duration.Minutes())
		}
	}
}

func TestShift_Duration(t *testing.T) {
	if duration, err := DefaultShift.GetShiftDuration(); err != nil {
		t.Errorf("Unexpected error getting duration for shift, Start [%v], End [%v]", DefaultShift.Start, DefaultShift.End)
	} else {
		if duration.Minutes() != expectedMinutesBetween {
			t.Errorf("expected DefaultShift start and end to be %d minutes apart, was [%v]",
				expectedMinutesBetween, duration.Minutes())
		}
	}
}

func TestShift_InvalidStartOrEndTime(t *testing.T) {
	invalidStart := time.Now().Add(time.Minute * 10)
	invalidEnd := time.Now()
	shift := CreateShift(invalidStart, invalidEnd, *createDefaultUser())
	if _, err := shift.GetShiftDuration(); err == nil {
		t.Errorf("expected error when endTime [%v] is before startTime [%v]", invalidEnd, invalidStart)
	}
}

func TestShift_AddNote(t *testing.T) {
	for i := 0; i < 10; i++ {
		if err := DefaultShift.AddNote(DefaultNote); err != nil {
			t.Errorf("error adding note [%v] to shift", DefaultNote.Note)
		}
		validateSliceLength(t, DefaultShift.Notes, "Shift.Notes", i+1)
	}
}

func TestShift_AddInvalidNote(t *testing.T) {
	if err := DefaultShift.AddNote(nil); err == nil {
		t.Errorf("error adding note [%v] to shift", nil)
	}
}

func createDefaultShift() *Shift {
	return CreateShift(DefaultStartTime, DefaultEndTime, *createDefaultUser())
}
