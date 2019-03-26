package model

import (
	"testing"
	"time"
)

//public for use in other tests
var DefaultStartTime time.Time
var DefaultEndTime time.Time

const expectedMinutesBetween = 60

func init() {
	DefaultStartTime = time.Now().Add(time.Hour)
	DefaultEndTime = DefaultStartTime.Add(time.Minute * expectedMinutesBetween)
}

func TestCreateShift(t *testing.T) {
	shift := createDefaultShift()

	t.Logf("shift.start, [%v]", shift.Start)
	t.Logf("shift.end ,  [%v]", shift.End)
	validateDefaultUser(t, &shift.User, 0)
	actualMinutesBetweenStartAndEnd := shift.End.Sub(shift.Start)
	if actualMinutesBetweenStartAndEnd.Minutes() != expectedMinutesBetween {
		t.Errorf("expected shift start and end to be %d minutes apart, was [%d]", expectedMinutesBetween, actualMinutesBetweenStartAndEnd)
	}
}

func createDefaultShift() *Shift {
	return CreateShift(DefaultStartTime, DefaultEndTime, *createDefaultUser())
}

//
//func Test_ShiftAddNote(t *testing.T) {
//	t.Fail()
//}
//
