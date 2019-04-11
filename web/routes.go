package web

import (
	"github.com/micklove/simple-roster/internal/app/config"
	"net/http"
)

//TODO -
// /rosters/{rosterid}
// /rosters/{rosterid}/users/
// /rosters/{rosterid}/shifts
// /rosters/{rosterid}/shifts?start=nnn&end=nnn

func Routes(cfg *app.Config) http.Handler {
	errorHelper := &ErrorHelper{
		*cfg,
	}
	rh := &RosterHandler{
		Config:      cfg,
		ErrorHelper: *errorHelper,
	}
	mux := http.NewServeMux()
	//mux.HandleFunc("/routes", http.HandlerFunc(GetWithID))
	mux.HandleFunc("/rosters", rh.byId())
	mux.HandleFunc("/rosters/", rh.get())
	return mux
}
