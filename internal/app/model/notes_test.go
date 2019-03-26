package model

import (
	"fmt"
	"testing"
	"time"
)

const defaultString string = "some note"
const multiLineString = `First line
Second line`

var DefaultNote *Note

func init() {
	DefaultNote = CreateNote(defaultString)
}

func TestCreateNote(t *testing.T) {
	if decoded, _ := DefaultNote.DecodeNote(); decoded != defaultString {
		t.Errorf("Note contains incorrect text, expected [%v], was [%v]", defaultString, DefaultNote.Note)
	}
}

func TestCreateNoteMultiLine(t *testing.T) {
	note := CreateNote(multiLineString)
	if decoded, _ := note.DecodeNote(); decoded != multiLineString {
		t.Errorf("Note contains incorrect text, expected [%v], was [%v]", multiLineString, note.Note)
	}
}

func Test_NotesAdd(t *testing.T) {
	var notes Notes = make([]Note, 0)
	notes, _ = notes.AddNote(CreateNote(defaultString))
	validateSliceLength(t, notes, "Notes", 1)

	notes, _ = notes.AddNote(CreateNote(defaultString))
	validateSliceLength(t, notes, "Notes", 2)
}

func Test_NotesAddInvalidNote(t *testing.T) {
	var notes Notes = make([]Note, 0)
	if _, err := notes.AddNote(nil); err == nil {
		t.Errorf("expected error when trying to add nil note")
	}
}

func Test_NodesAddInvalidBase64String(t *testing.T) {
	note := CreateNote("Hello World")
	note.Note = "gibberish"
	if _, err := note.DecodeNote(); err == nil {
		t.Errorf("expected Decode to fail on invalid base64")
	}
}

func Test_NoteHasRecentlyCreatedTimeStamp(t *testing.T) {
	note := CreateNote(defaultString)
	now := time.Now()
	fmt.Printf("Note    [%v]\n", note.Date.Unix())
	fmt.Printf("Now     [%v]\n", now.Unix())
	elapsed := now.Sub(note.Date)
	fmt.Printf("Elapsed [%v]\n", elapsed.Seconds())
	expectedMaxTimeDiffInSeconds := 5.0
	if elapsed.Seconds() > expectedMaxTimeDiffInSeconds {
		t.Errorf("Expected Note to contain time stamp created in last [%v] seconds, was [%v] seconds",
			expectedMaxTimeDiffInSeconds, elapsed.Seconds())
	}
}
