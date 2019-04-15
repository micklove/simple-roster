package web

import (
	"encoding/json"
	"github.com/micklove/simple-roster/internal/app/config"
	"github.com/micklove/simple-roster/internal/app/dao"
	"github.com/micklove/simple-roster/internal/app/model"
	"github.com/micklove/simple-roster/internal/app/service"
	"github.com/micklove/simple-roster/internal/pkg/UUID"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//See https://blog.questionable.services/article/testing-http-handlers-go/
var DefaultEndpoint = "http://localhost:8080/rosters"

//func TestRosterById(t *testing.T) {
//	//p := Routes(nil)
//	//
//	//var ok bool
//
//	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
//	// pass 'nil' as the third parameter.
//	req, err := http.NewRequest("GET", DefaultEndpoint, nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
//	rr := httptest.NewRecorder()
//
//	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
//	// directly and pass in our Request and ResponseRecorder.
//	Routes(nil).ServeHTTP(rr, req)
//
//	// Check the status code is what we expect.
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//	//}
//	//p.Get("/foo/:name", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	//	ok = true
//	//	t.Logf("%#v", r.URL.Query())
//	//	if got, want := r.URL.Query().Get(":name"), "keith"; got != want {
//	//		t.Errorf("got %q, want %q", got, want)
//	//	}
//	//}))
//	//
//	//p.ServeHTTP(nil, newRequest("GET", "/foo/keith?a=b", nil))
//	//if !ok {
//	//	t.Error("handler not called")
//	//}
//}

//
//func TestRosters(t *testing.T) {
//	rosterName := "test-roster"
//	mockRoster, _ := model.CreateRoster(rosterName)
//
//	rr, res, _ := getDefaultRecorderAndResponse(t, mockRoster, DefaultEndpoint)
//	expectedHttpStatus := http.StatusOK
//	validateRosterResponse(t, rr, res, mockRoster.ID, rosterName, mockRoster.Shifts, expectedHttpStatus)
//
//}
var uuidGenerator = &UUID.KSUUIDGenerator{}

func TestRosterById(t *testing.T) {
	rosterName := "test-roster"
	mockRoster, _ := model.CreateRoster(rosterName, uuidGenerator)
	expectedID := mockRoster.ID

	endpointWithId := DefaultEndpoint + "?id=" + expectedID
	rr, res, _ := getDefaultRecorderAndResponse(t, mockRoster, endpointWithId)
	expectedHttpStatus := http.StatusOK
	validateRosterResponse(t, rr, res, expectedID, rosterName, mockRoster.Shifts, expectedHttpStatus)
}

func TestRosterNotFoundResponse(t *testing.T) {
	mockRoster, _ := model.CreateRoster("non existent roster", uuidGenerator)
	//mockRoster, _ := model.CreateRoster("non existent roster")
	mockRoster.ID = "NOT FOUND"
	expectedHttpStatus := http.StatusNotFound
	rr, res, _ := getDefaultRecorderAndResponse(t, mockRoster, DefaultEndpoint+"?id=JUNK")
	validateRosterResponse(t, rr, res, "", "", mockRoster.Shifts, expectedHttpStatus)
}

func validateRosterResponse(t *testing.T,
	rr *httptest.ResponseRecorder,
	response *Response,
	expectedID string,
	expectedName string,
	expectedShifts model.Shifts,
	expectedHttpStatus int) {

	//is the response.Data interface the correct type? (RosterWrapper)
	if rosterWrapper, ok := response.Data.(RosterWrapper); ok {
		roster := rosterWrapper.Roster
		validateRosterShifts(t, roster, expectedShifts)
		if expectedID != "" {
			validateRosterId(t, roster, expectedID)
		}
		if expectedName != "" {
			validateRosterName(t, roster, expectedName)
		}
	}
	validateHttpStatus(t, rr, expectedHttpStatus)
}

func validateRosterShifts(t *testing.T, roster *model.Roster, expectedShifts model.Shifts) {
	want := len(expectedShifts)
	got := len(roster.Shifts)
	if want != got {
		t.Errorf("Expected [%v] shifts, got [%v] shifts", want, got)
	}
}

func validateRosterId(t *testing.T, roster *model.Roster, expectedId string) {
	want := expectedId
	got := roster.ID
	if want != got {
		t.Errorf("Expected roster ID [%v], got [%v]", want, got)
	}
}

func validateRosterName(t *testing.T, roster *model.Roster, expectedName string) {
	want := expectedName
	got := roster.Name
	if want != got {
		t.Errorf("Expected roster Name [%v], Got [%v]", want, got)
	}
}

func validateHttpStatus(t *testing.T, rr *httptest.ResponseRecorder, expectedHttpStatus int) {
	want := expectedHttpStatus
	got := rr.Code

	if want != got {
		t.Errorf("expected http status of [%v], Got [%v]", want, got)
	}
}

func getDefaultRecorderAndResponse(t *testing.T, roster *model.Roster, url string) (rr *httptest.ResponseRecorder, res *Response, body []byte) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	//Use mock dao, in RosterService
	rr = httptest.NewRecorder()
	mockRosterDao := &dao.MockRosterDao{}
	mockRosterDao.SetMockResponse(roster)

	cfg := &app.Config{
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	router := &Router{
		Config: cfg,
		RosterService: &service.RosterService{
			RosterDao: mockRosterDao,
		},
	}
	handler := router.Routes()
	handler.ServeHTTP(rr, req)

	body, err = ioutil.ReadAll(rr.Body)
	if err != nil {
		log.Fatal(err)
	}

	//var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Errorf("could not unmarshall response to Response type for: %s \n", string(body))
	}

	return rr, res, body
}
