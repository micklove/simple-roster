package model

import "time"

type Note struct {
	Date time.Time `json:"date"`
	Note string    `json:"note"`
}

func CreateNote(note string) *Note {
	return &Note{
		Note: note,
		Date: time.Now(),
	}
}
