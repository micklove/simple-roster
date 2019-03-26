package model

import (
	"testing"
)

//public - used in other tests
const DefaultFirstName = "ronnie"
const DefaultLastName = "osullivan"
const DefaultDisplayName = "Rocket Ronnie"
const DefaultNoteString = "Hello World"

var DefaultUser *User

func init() {
	DefaultUser = createDefaultUser()
}

func Test_UserCreate(t *testing.T) {
	validateDefaultUser(t, DefaultUser, 0)
}

func Test_UserAddNote(t *testing.T) {
	user := createDefaultUser()
	note := CreateNote(DefaultNoteString)
	err := user.AddNote(note)

	if err != nil {
		t.Errorf("Error adding note to user")
	}
	validateDefaultUser(t, user, 1)
}

func TestUser_HasDisplayName(t *testing.T) {
	matchNameField(t, DefaultUser.HasDisplayName, DefaultDisplayName, "DisplayName")
}

func TestUser_HasFirstName(t *testing.T) {
	matchNameField(t, DefaultUser.HasFirstName, DefaultFirstName, "FirstName")
}

func TestUser_HasLastName(t *testing.T) {
	matchNameField(t, DefaultUser.HasLastName, DefaultLastName, "LastName")
}

func TestUser_MatchesAnyName(t *testing.T) {
	matchNameField(t, DefaultUser.MatchesAnyName, DefaultFirstName, "FirstName")
	matchNameField(t, DefaultUser.MatchesAnyName, DefaultLastName, "LastName")
	matchNameField(t, DefaultUser.MatchesAnyName, DefaultDisplayName, "DisplayName")
}

//f func(shift Shift) bool
func matchNameField(t *testing.T, f func(s string) bool, searchString string, searchField string) {
	if !f(searchString) {
		t.Errorf("Expected search to match Name Field %v prefix of [%v]", searchField, searchString)
	}
}

func createDefaultUser() *User {
	return CreateUser(DefaultFirstName, DefaultLastName, DefaultDisplayName)
}

func validateDefaultUser(t *testing.T, user *User, expectedNotesLength int) {
	validateUser(t, user, DefaultFirstName, DefaultLastName, DefaultDisplayName, expectedNotesLength)
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
