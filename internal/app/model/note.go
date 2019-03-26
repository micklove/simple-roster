package model

import (
	"encoding/base64"
	"fmt"
	"time"
)

//TODO - base64
type Note struct {
	Date time.Time `json:"date"`
	Note string    `json:"note"`
}

//TODO - Encode
func CreateNote(note string) *Note {
	encoded := base64.StdEncoding.EncodeToString([]byte(note))

	return &Note{
		Note: encoded,
		Date: time.Now(),
	}
}

//So that we can accommodate multi-line notes, the string is base64 encoded on creation, decode
func (note *Note) DecodeNote() (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(note.Note)
	if err != nil {
		fmt.Printf("decode error [%v]: cannot decode [%v]", err, note.Note)
		return note.Note, err
	}
	return string(decoded), nil
}
