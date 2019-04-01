package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

//TODO add Role
//TODO add ShiftType
//TODO add ID
//TODO add durationInMinutes
type Shift struct {
	RosterId string    `json:"rosterId"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	User     User      `json:"assignee"`
	Notes    Notes     `json:"notes"`
}

type ShiftSummary struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	STime string    `json:"s_time"`
	ETime string    `json:"e_time"`
}

func CreateShift(rosterId string, start time.Time, end time.Time, user User) *Shift {
	return &Shift{
		RosterId: rosterId,
		Start:    start,
		End:      end,
		User:     user,
		Notes:    make([]Note, 0),
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

//Shift Matchers - for FilterByTime methods
func (shift Shift) Between(start time.Time, end time.Time) bool {
	fmt.Printf("S: [%v], E: [%v], Shift [%v]\n", start, end, shift)
	afterStart := shift.afterStart(start)
	beforeEnd := shift.beforeEnd(end)
	return afterStart && beforeEnd
}

// nb: variance of 1 second allowed
func (shift Shift) afterStart(time time.Time) bool {
	fmt.Printf("S: [%v], Shift [%v]\n", time, shift)
	diff := time.Sub(shift.Start).Seconds()
	return diff < 1
}

// nb: variance of 1 second allowed
func (shift Shift) beforeEnd(time time.Time) bool {
	fmt.Printf("E: [%v], Shift [%v]\n", time, shift)
	diff := time.Sub(shift.End).Seconds()
	return diff > -1
}

func (shift Shift) Dump() {
	dumpToJSON(shift)
}

func (shift Shift) DumpSummary() {
	summary := &ShiftSummary{
		shift.Start,
		shift.End,
		fmt.Sprintf("%02d:%02d:%02d", shift.Start.Hour(), shift.Start.Minute(), shift.Start.Second()),
		fmt.Sprintf("%02d:%02d:%02d", shift.End.Hour(), shift.End.Minute(), shift.End.Second()),
	}
	dumpToJSON(summary)
}

func dumpToJSON(obj interface{}) {
	b, _ := json.MarshalIndent(obj, "", "\t")
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n"))
}
