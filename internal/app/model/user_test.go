package model

import (
	"testing"
)

const defaultFirstName = "ronnie"
const defaultLastName = "osullivan"
const defaultDisplayName = "Rocket Ronnie"
const defaultNote = "Hello World"

func Test_UserCreate(t *testing.T) {
	user := createDefaultUser()
	validateDefaultUser(t, user, 0)
}

func Test_UserAddNote(t *testing.T) {
	user := createDefaultUser()
	note := CreateNote(defaultNote)
	err := user.AddNote(note)

	if err != nil {
		t.Errorf("Error adding note to user")
	}
	validateDefaultUser(t, user, 1)
}

func createDefaultUser() *User {
	return CreateUser(defaultFirstName, defaultLastName, defaultDisplayName)
}

func validateDefaultUser(t *testing.T, user *User, expectedNotesLength int) {
	validateUser(t, user, defaultFirstName, defaultLastName, defaultDisplayName, expectedNotesLength)
}

func validateUser(t *testing.T, user *User, expectedFirstName string, expectedLastName string, expectedDisplayName string, expectedNotesLength int) {

	if user.FirstName != expectedFirstName {
		t.Errorf("Expected first name to be [%v], was [%v]", expectedFirstName, user.FirstName)
	}

	if user.LastName != expectedLastName {
		t.Errorf("Expected last name to be [%v], was [%v]", expectedLastName, user.LastName)
	}

	if user.DisplayName != expectedDisplayName {
		t.Errorf("Expected last name to be [%v], was [%v]", expectedDisplayName, user.DisplayName)
	}

	if len(user.Notes) != expectedNotesLength {
		t.Errorf("Expected user Notes to have length [%v], was [%v]", expectedNotesLength, len(user.Notes))
	}
}
