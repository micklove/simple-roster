package model

import "errors"

type Notes []Note

func (notes Notes) AddNote(note *Note) (Notes, error) {
	if note == nil {
		err := errors.New("cannot add empty note to Notes")
		return notes, err
	}
	return append(notes, *note), nil
}
