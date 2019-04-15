package model

import (
	"fmt"
	"github.com/micklove/simple-roster/internal/pkg/UUID"
	"reflect"
	"testing"
	"time"
)

var DefaultRosterId = "some-roster-id"

//public for use in other tests
var DefaultStartTime time.Time
var DefaultEndTime time.Time
var DefaultShift *Shift

const DefaultFirstName = "ronnie"
const DefaultLastName = "osullivan"
const DefaultDisplayName = "Rocket Ronnie"
const DefaultUserRole = "snooker player"

const DefaultNoteString = "Hello World"
const DefaultAvatarUrl = "http://localhost:8080/ronnie.jpg"

const DisplayNamePrefix = "maestro"
const FirstNamePrefix = "stephen"
const LastNamePrefix = "hendry"

var DefaultShifts Shifts
var DefaultFirstShift *Shift

//Attempt at a reasonably generic method :)
//Take a composite struct, e.g. Notes (== []Note) and test the length
func validateSliceLength(t *testing.T, obj interface{}, typeName string, expectedNoteLength int) {

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		genericSlice := reflect.ValueOf(obj)

		//t.Logf("struct length [%v]", genericSlice.Len())

		if genericSlice.Len() != expectedNoteLength {
			t.Errorf("Expected %v of count [%v], was [%v]", typeName, expectedNoteLength, genericSlice.Len())
		}
	} else {
		t.Errorf("Given obj %v is not an Array", typeName)
		t.Fail()
	}
}

func ValidateShifts(t *testing.T, shifts Shifts, expectedRosterId string) {
	for _, shift := range shifts {
		if shift.RosterId != expectedRosterId {
			t.Errorf("expected RosterId of [%v], was [%v]", expectedRosterId, shift.RosterId)
			t.Fail()
		}
	}
}

//Create a list of shifts, where half of the list have a different first, last name
func createListOfShiftsWithDifferentNames(t *testing.T, rosterId string, count int) Shifts {
	shifts := CreateShifts(rosterId, count)

	//nb: can't use the range variable to update each shift, as it's a local copy -
	//    see https://yourbasic.org/golang/gotcha-change-value-range/
	for i, _ := range shifts {
		if i%2 == 0 {
			shifts[i].User.DisplayName = fmt.Sprintf("%v-%v", DisplayNamePrefix, i)
			shifts[i].User.FirstName = fmt.Sprintf("%v-%v", FirstNamePrefix, i)
			shifts[i].User.LastName = fmt.Sprintf("%v-%v", LastNamePrefix, i)
		}
	}
	return shifts
}

//Create a list, where half of the list have a different first, last name
func CreateShifts(rosterId string, shiftCount int) Shifts {
	var shifts Shifts = make([]Shift, 0, shiftCount)
	for i := 0; i < shiftCount; i++ {
		shift := CreateShift(rosterId, time.Now(), time.Now().Add(time.Hour*24), *CreateDefaultUser(), DefaultUserRole)
		shifts, _ = shifts.AddShift(shift)
	}
	return shifts
}

func CreateDefaultUser() *User {
	var uuidGenerator = &UUID.KSUUIDGenerator{}
	user, _ := NewUser(DefaultFirstName, DefaultLastName, DefaultDisplayName, DefaultAvatarUrl, uuidGenerator)
	return user
}

func createDefaultShift(rosterId string) *Shift {
	return CreateShift(rosterId, DefaultStartTime, DefaultEndTime, *CreateDefaultUser(), DefaultUserRole)
}
