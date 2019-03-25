package model

import (
	"fmt"
	"time"
)

//TODO add Role
//TODO add ShiftType
type Shift struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	User  User      `json:"assignee"`
	Notes Notes     `json:"notes"`
}

func CreateShift() {
	fmt.Println("Shift")
}
