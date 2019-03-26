package model

import (
	"fmt"
	"testing"
	"time"
)

const defaultString string = "some note"
const multiLineString = `First line
Second line`

func TestCreateNote(t *testing.T) {
	note := CreateNote(defaultString)
	if decoded, _ := note.DecodeNote(); decoded != defaultString {
		t.Errorf("Note contains incorrect text, expected [%v], was [%v]", defaultString, note.Note)
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
	validateSliceLength(t, notes, "Shifts", 1)

	notes, _ = notes.AddNote(CreateNote(defaultString))
	validateSliceLength(t, notes, "Shifts", 2)
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

func createDefaultNotes() Notes {
	note := CreateNote(defaultString)
	notes := Notes{*note}
	return notes
}
