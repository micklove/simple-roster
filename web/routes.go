package web

import (
	"github.com/micklove/simple-roster/internal/app/config"
	"github.com/micklove/simple-roster/internal/app/service"
	"net/http"
)

//TODO - semantic urls
// /rosters/:rosterid
// /rosters/:rosterid/users/
// /rosters/:rosterid/shifts
// /rosters/:rosterid/shifts?start=nnn&end=nnn

type Router struct {
	RosterService *service.RosterService
}

func (router *Router) Routes(cfg *app.Config) http.Handler {
	errorHelper := &ErrorHelper{
		*cfg,
	}
	rh := &RosterHandler{
		Config:        cfg,
		ErrorHelper:   *errorHelper,
		RosterService: router.RosterService,
	}
	mux := http.NewServeMux()
	//mux.HandleFunc("/routes", http.HandlerFunc(GetWithID))
	mux.HandleFunc("/rosters", rh.byId())
	mux.HandleFunc("/rosters/", rh.get())
	return mux
}
