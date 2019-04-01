package model

import (
	"time"
)

type Schedule struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Shifts `json:"shifts"`
}

//TODO
func (schedule *Schedule) GetShifts(startTime time.Time, endTime time.Time) Shifts {
	return schedule.Shifts
}

//LOOK at fill struct from visual studio go
