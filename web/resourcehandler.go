package web

import (
	"encoding/json"
	"fmt"
	"github.com/micklove/simple-roster/internal/app/config"
	"io"
	"net/http"
)

type ResourceHandler struct {
	Config      *app.Config
	ErrorHelper ErrorHelper
}

func (rh ResourceHandler) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

func (rh *ResourceHandler) writeResponse(w http.ResponseWriter) {

	rel := "roster"
	href := "/" + rel

	x := struct {
		HREF string `json:"href"`
		REL1 string `json:"rel"`
	}{
		href,
		rel,
	}

	response := NewResponse(x)
	//jr, _ := json.MarshalIndent(response, "", "\t")

	pretty, _ := json.MarshalIndent(response, "", "  ")
	respStr := string(pretty)
	//fmt.Println(respStr)
	rh.Config.InfoLog.Printf("Response:\n%v\n", respStr)

	io.WriteString(w, respStr)
}

func (rh *ResourceHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	rh.writeResponse(w)
}

//Create
func (rh *ResourceHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("handlePost - Error %v NOT implemented", r.Method)
	rh.ErrorHelper.serverError(w, err)
}

//Update
func (rh *ResourceHandler) handlePut(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("handlePut - Error %v NOT implemented", r.Method)
	rh.ErrorHelper.serverError(w, err)
}
