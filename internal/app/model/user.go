package model

import "errors"

type User struct {
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
