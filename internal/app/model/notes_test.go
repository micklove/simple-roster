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
	DefaultNote = NewNote(defaultString)
}

func TestCreateNote(t *testing.T) {
	if decoded, _ := DefaultNote.DecodeNote(); decoded != defaultString {
		t.Errorf("Note contains incorrect text, expected [%v], was [%v]", defaultString, DefaultNote.Note)
	}
}

func TestCreateNoteMultiLine(t *testing.T) {
	note := NewNote(multiLineString)
	if decoded, _ := note.DecodeNote(); decoded != multiLineString {
		t.Errorf("Note contains incorrect text, expected [%v], was [%v]", multiLineString, note.Note)
	}
}

func Test_NotesAdd(t *testing.T) {
	var notes Notes = make([]Note, 0)
	notes, _ = notes.AddNote(NewNote(defaultString))
	validateSliceLength(t, notes, "Notes", 1)

	notes, _ = notes.AddNote(NewNote(defaultString))
	validateSliceLength(t, notes, "Notes", 2)
}

func Test_NotesAddInvalidNote(t *testing.T) {
	var notes Notes = make([]Note, 0)
	if _, err := notes.AddNote(nil); err == nil {
		t.Errorf("expected error when trying to add nil note")
	}
}

func Test_NodesAddInvalidBase64String(t *testing.T) {
	note := NewNote("Hello World")
	note.Note = "gibberish"
	if _, err := note.DecodeNote(); err != nil {
		return //Expected Error, ignore
	}
	t.Errorf("expected Decode to fail on invalid base64")
}

func Test_NoteHasRecentlyCreatedTimeStamp(t *testing.T) {
	note := NewNote(defaultString)
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

//TODO Fix - json object to parse, unmarshalJSONText,  is invalid
//
//func TestNote_UnmarshalJSON(t *testing.T) {
//	//expectedDateStr := "2019-04-03T16:27:58.111197+11:00"
//	//unmarshalJSONText := `[{
//	//	\\"date\\": \\"2019-04-03T16:27:58.111197+11:00\\",
//	//	\\"note\\": \\"aGVsbG9cbldvcmxkISEK\\"
//	//}]
//	//`
//	expectedDateStr := time.Now().String()
//	unmarshalJSONText := fmt.Sprintf(string(`{"date": "%v","note": "aGVsbG9cbldvcmxkISEK"}`), expectedDateStr)
//	note := &Note{}
//
//	//See https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/2678
//	if err := note.UnmarshalJSON([]byte(unmarshalJSONText)); err != nil {
//		//msg := fmt.Sprintf("Unexpected UnMarshal error in parsing Note JSON [%v]", unmarshalJSONText)
//		msg := "unexpected marshall error"
//		t.Error(msg)
//		t.Fail() ///explicit fail here , as test was not failing on output with ---FAIL, not sure why
//		panic("Exit here, test failed")
//	}
//	actualDate := note.Date.String()
//	if actualDate != expectedDateStr {
//		t.Errorf("Expected Note to have expected date [%v] , after UnMarshal", expectedDateStr)
//	}
//
//}
