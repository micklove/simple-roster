package model

import (
	"encoding/json"
	"fmt"
	"github.com/micklove/simple-roster/internal/pkg/UUID"
	"os"
)

type Roster struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Shifts `json:"shifts"`
}

func CreateRoster(name string, uuidGenerator UUID.Generator) (roster *Roster, err error) {
	var generatedId string
	if generatedId, err = uuidGenerator.Create(); err != nil {
		return nil, fmt.Errorf("error getting new UUID to use Roster ID, err [%v]", err)
	}
	return &Roster{
		ID:     generatedId,
		Name:   name,
		Shifts: make([]Shift, 0),
	}, nil
}

//Add a shift to the Shifts collection on the Roster
// nb: Would have preferred to use an 'Embedded field' here, (i.e. without the reference to
//     roster.shifts.AddShift, with just roster.AddShift, and save recreating the method
//     However, the underlying 'append' method on Shifts returns a copy of the []Shift array
//      - See See https://golang.org/ref/spec#Struct_types
func (roster *Roster) AddShift(shift *Shift) (Shifts, error) {
	var shifts Shifts
	var err error
	shift.RosterId = roster.ID
	if shifts, err = roster.Shifts.AddShift(shift); err != nil {
		return shifts, err
	}
	roster.Shifts = shifts
	return shifts, nil
}

//TODO
//func GetAllShifts(startTime time.Time, endTime time.Time) {
//	fmt.Println("GetAllShifts")
//}
//
//func GetShiftsByTime(startTime time.Time, endTime time.Time) {
//	fmt.Println("GetShiftsByTime")
//}

func (roster *Roster) Dump() {
	dumpRoster(roster)
}

//TODO - move to it's own pkg
func dumpRoster(obj interface{}) {
	b, _ := json.MarshalIndent(obj, "", "\t")
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n"))
}
