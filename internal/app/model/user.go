package model

import (
	"errors"
	"strings"
)

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DisplayName string `json:"displayName"`
	Notes       Notes  `json:"notes"`
}

func CreateUser(firstName string, lastName string, displayName string) *User {
	return &User{
		FirstName:   firstName,
		LastName:    lastName,
		DisplayName: displayName,
		Notes:       make([]Note, 0),
	}
}

func (user *User) AddNote(note *Note) error {
	var notes Notes
	notes, err := user.Notes.AddNote(note)
	if err != nil {
		return errors.New("error adding note to User")
	}
	user.Notes = notes
	return nil
}

//User Matchers - for Filter methods
func (user *User) HasDisplayName(searchNamePrefix string) bool {
	return strings.HasPrefix(user.DisplayName, searchNamePrefix)
}

func (user *User) HasFirstName(searchNamePrefix string) bool {
	return strings.HasPrefix(user.FirstName, searchNamePrefix)
}

func (user *User) HasLastName(searchNamePrefix string) bool {
	return strings.HasPrefix(user.LastName, searchNamePrefix)
}

func (user *User) MatchesAnyName(searchNamePrefix string) bool {
	return strings.HasPrefix(user.FirstName, searchNamePrefix) ||
		strings.HasPrefix(user.LastName, searchNamePrefix) ||
		strings.HasPrefix(user.DisplayName, searchNamePrefix)
}
