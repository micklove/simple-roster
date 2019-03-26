package model

import (
	"fmt"
	"time"
)

type Schedule struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Shifts Shifts `json:"shifts"`
}

//TODO
func GetShifts(startTime time.Time, endTime time.Time) {
	fmt.Println("Roster")
}
