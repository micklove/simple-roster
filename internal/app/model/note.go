package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

type Note struct {
	Date time.Time `json:"date"`
	Note string    `json:"note"`
}

func NewNote(note string) *Note {
	encoded := base64.StdEncoding.EncodeToString([]byte(note))
	return &Note{
		Note: encoded,
		Date: time.Now(),
	}
}

//So that we can accommodate multi-line notes, the string is base64 encoded on creation, decode
func (note *Note) DecodeNote() (decoded string, err error) {
	return decode(note.Note)
}

func decode(noteStr string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(noteStr)
	if err != nil {
		return "", fmt.Errorf("decode error [%v]: cannot decode [%v]\n", err, noteStr)
	}
	return string(decoded), err
}

//Ensure the decoded json, e.g. from the File dao, is valid base64
// e.g. To get a base64 string In osx, base64 -i- <<< "hello world"
func (note *Note) UnmarshalJSON(data []byte) error {

	//fmt.Printf("UnmarshalJSON String = [%v]\n", string(data))
	//Unmarshal the type Note to NewNote struct, same fields, no methods (prevents infinite loop on the Unmarshal)
	aux := &struct {
		Date time.Time `json:"date"`
		Note string    `json:"note"`
	}{}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	decoded, err := decode(aux.Note)
	note.Note = decoded
	note.Date = aux.Date
	return err
}

//
//func (note *Note) MarshalJSON() ([]byte, error) {
//	m := map[string]string{
//		"id":   fmt.Sprintf("0x%08x", a.Id),
//		"name": a.Name,
//	}
//	return json.Marshal(m)
//}
