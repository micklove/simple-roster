package model

import (
	"errors"
	"time"
)

//TODO add Role
//TODO add ShiftType
type Shift struct {
	RosterId int64     `json:"rosterId"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	User     User      `json:"assignee"`
	Notes    Notes     `json:"notes"`
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

func CreateShift(start time.Time, end time.Time, user User) *Shift {
	return &Shift{
		Start: start,
		End:   end,
		User:  user,
		Notes: make([]Note, 0),
	}
}
