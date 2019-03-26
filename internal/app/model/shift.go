package model

import (
	"errors"
	"time"
)

//TODO add Role
//TODO add ShiftType
//TODO add ID
//TODO add durationInMinutes
type Shift struct {
	RosterId int64     `json:"rosterId"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	User     User      `json:"assignee"`
	Notes    Notes     `json:"notes"`
}

func CreateShift(start time.Time, end time.Time, user User) *Shift {
	return &Shift{
		Start: start,
		End:   end,
		User:  user,
		Notes: make([]Note, 0),
	}
}

func (shift *Shift) AddNote(note *Note) error {
	var notes Notes
	notes, err := shift.Notes.AddNote(note)
	if err != nil {
		return errors.New("error adding note to shift")
	}
	shift.Notes = notes
	return nil
}

func (shift *Shift) GetShiftDuration() (time.Duration, error) {
	duration := shift.End.Sub(shift.Start)
	if duration.Seconds() < 0 {
		err := errors.New("cannot add nil link to error, ignoring")
		return duration, err
	}
	return duration, nil
}
