package web

import (
	"encoding/json"
	"fmt"
	"github.com/micklove/simple-roster/internal/app/config"
	"github.com/micklove/simple-roster/internal/app/model"
	"io"
	"net/http"
)

type RosterHandler struct {
	Config      *app.Config
	ErrorHelper ErrorHelper
}

func (rh RosterHandler) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO handle request
		w.Header().Set("Content-Type", "application/json")
		rh.Config.InfoLog.Printf("get()")
		switch r.Method {
		case "GET":
			rh.handleGet(w, r)
		case "POST":
			rh.handlePost(w, r)
		case "PUT":
			rh.handlePut(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (rh RosterHandler) byId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		rh.Config.InfoLog.Printf("byId()")
		switch r.Method {
		case "GET":
			rh.handleGetById(w, r)
		case "POST":
			rh.handlePost(w, r)
		case "PUT":
			rh.handlePut(w, r)
		default:
			rh.ErrorHelper.clientError(w, http.StatusMethodNotAllowed)
		}
	}
}

func (rh *RosterHandler) handleGetById(w http.ResponseWriter, r *http.Request) {
	var err error
	var rosterId string
	rosterId = r.URL.Query().Get("id")
	key := getKey(r)
	rh.Config.InfoLog.Printf("key: [%v]", key)

	if len(rosterId) < 1 {
		//rh.ErrorHelper.notFound(w) // Use the notFound() helper.
		rh.ErrorHelper.clientError(w, http.StatusBadRequest)
		return
	}
	//rosterId := "1JLL9NxeyvqhLnDGcMD8MM20H9p"
	var roster *model.Roster
	if roster, err = rh.Config.RosterService.ByID(rosterId); err != nil {
		rh.ErrorHelper.serverError(w, err)
		fmt.Errorf("error retrieving roster id [%v]", rosterId)
	}
	rh.writeResponse(w, roster)
}

func (rh *RosterHandler) writeResponse(w http.ResponseWriter, roster *model.Roster) {
	if roster == nil {
		w.WriteHeader(http.StatusNotFound)
	}
	rosterWrapper := NewRosterWrapper(roster)
	response := NewResponse(rosterWrapper)
	//jr, _ := json.MarshalIndent(response, "", "\t")

	pretty, _ := json.MarshalIndent(response, "", "\t")
	respStr := string(pretty)
	//fmt.Println(respStr)
	rh.Config.InfoLog.Printf("Response:\n%v\n", respStr)

	io.WriteString(w, respStr)
}

func (rh *RosterHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	var err error
	var rosterId string
	rosterId = r.URL.Query().Get("id")
	key := getKey(r)
	rh.Config.InfoLog.Printf("key: [%v]", key)

	if len(rosterId) < 1 {
		rh.ErrorHelper.clientError(w, http.StatusBadRequest)
		return
	}
	var roster *model.Roster
	if roster, err = rh.Config.RosterService.ByID(rosterId); err != nil {
		rh.ErrorHelper.serverError(w, err)
		fmt.Errorf("error retrieving roster id [%v]", rosterId)
	}
	rh.writeResponse(w, roster)
}

//Create
func (rh *RosterHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("handlePost - Error %v NOT implemented", r.Method)
	rh.ErrorHelper.serverError(w, err)
}

//Update
func (rh *RosterHandler) handlePut(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("handlePut - Error %v NOT implemented", r.Method)
	rh.ErrorHelper.serverError(w, err)
}

//See https://gist.github.com/tomnomnom/52dfa67c7a8c9643d7ce

func getKey(r *http.Request) string {
	return r.URL.Path[len("/rosters"):]
}
