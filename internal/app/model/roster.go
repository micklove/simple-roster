package model

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"time"
)

type Roster struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Shifts Shifts `json:"shifts"`
}

func CreateRoster(name string) (*Roster, error) {

	if id, err := ksuid.NewRandomWithTime(time.Now()); err != nil {
		fmt.Printf("error getting new UUID for Roster ID, err [%v]", err)
		return nil, err
	} else {
		return &Roster{
			ID:     id.String(),
			Name:   name,
			Shifts: make([]Shift, 0),
		}, nil
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
//func GetAllShifts(startTime time.Time, endTime time.Time) {
//	fmt.Println("GetAllShifts")
//}
//
//func GetShiftsByTime(startTime time.Time, endTime time.Time) {
//	fmt.Println("GetShiftsByTime")
//}
