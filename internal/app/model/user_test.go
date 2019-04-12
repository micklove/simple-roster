package model

import (
	"net/url"
	"testing"
)

var DefaultUser *User

func init() {
	DefaultUser = CreateDefaultUser()
}

func Test_UserCreate(t *testing.T) {
	validateDefaultUser(t, DefaultUser, 0)
}

func Test_UserAddNote(t *testing.T) {
	user := CreateDefaultUser()
	note := NewNote(DefaultNoteString)
	err := user.AddNote(note)

	if err != nil {
		t.Errorf("Error adding note to user")
	}
	validateDefaultUser(t, user, 1)
}

func Test_UserAddInvalidNote(t *testing.T) {
	user := CreateDefaultUser()
	if err := user.AddNote(nil); err == nil {
		t.Errorf("error expected to fail when adding nil note to User")
	}
}

func TestUser_HasDisplayName(t *testing.T) {
	matchNameField(t, *DefaultUser, User.HasDisplayName, DefaultDisplayName, "DisplayName")
}

func TestUser_HasFirstName(t *testing.T) {
	matchNameField(t, *DefaultUser, User.HasFirstName, DefaultFirstName, "FirstName")
}

func TestUser_HasLastName(t *testing.T) {
	matchNameField(t, *DefaultUser, User.HasLastName, DefaultLastName, "LastName")
}

func TestUser_MatchesAnyName(t *testing.T) {
	matchNameField(t, *DefaultUser, User.MatchesAnyName, DefaultFirstName, "FirstName")
	matchNameField(t, *DefaultUser, User.MatchesAnyName, DefaultLastName, "LastName")
	matchNameField(t, *DefaultUser, User.MatchesAnyName, DefaultDisplayName, "DisplayName")
}

//f func(shift Shift) bool
func matchNameField(t *testing.T, user User, f func(User, string) bool, searchString string, searchField string) {
	if !f(user, searchString) {
		t.Errorf("Expected search to match Name Field %v prefix of [%v]", searchField, searchString)
	}
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

	var parsedUrl *url.URL
	var err error
	if parsedUrl, err = url.ParseRequestURI(user.Avatar); err != nil {
		t.Errorf("Error parsing url [%v]", user.Avatar)
	}

	if user.AvatarUrl != *parsedUrl {
		t.Errorf("Avatar urls do not match, want [%v], got [%v]", user.AvatarUrl, parsedUrl)
	}

}
