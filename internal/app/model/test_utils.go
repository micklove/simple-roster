package model

import (
	"reflect"
	"testing"
)

//Attempt at a reasonably generic method :)
//Take a composite struct, e.g. Notes (== []Note) and test the length
func validateSliceLength(t *testing.T, obj interface{}, typeName string, expectedNoteLength int) {

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		genericSlice := reflect.ValueOf(obj)

		t.Logf("struct length [%v]", genericSlice.Len())

		if genericSlice.Len() != expectedNoteLength {
			t.Errorf("Expected %v of count [%v], was [%v]", typeName, expectedNoteLength, genericSlice.Len())
		}
	} else {
		t.Errorf("Given obj %v is not an Array", typeName)
		t.Fail()
	}
}
