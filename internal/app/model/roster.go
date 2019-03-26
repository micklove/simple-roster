package model

import (
	"fmt"
	"time"
)

type Roster struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Shifts Shifts `json:"shifts"`
}

func CreateRoster(name string) *Roster {
	return &Roster{
		Name:   name,
		Shifts: make([]Shift, 0),
	}
}

func (roster *Roster) AddShift(shift *Shift) (Shifts, error) {
	if shifts, err := roster.Shifts.AddShift(shift); err != nil {
		return shifts, err
	} else {
		roster.Shifts = shifts
		return shifts, nil
	}
}

//TODO
func GetAllShifts(startTime time.Time, endTime time.Time) {
	fmt.Println("GetAllShifts")
}

func GetShiftsByTime(startTime time.Time, endTime time.Time) {
	fmt.Println("GetShiftsByTime")
}
