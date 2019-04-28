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
	Config        *app.Config
}

func (router *Router) Routes() http.Handler {
	errorHelper := &ErrorHelper{
		*router.Config,
	}
	resourceHandler := &ResourceHandler{
		Config:      router.Config,
		ErrorHelper: *errorHelper,
	}
	rh := &RosterHandler{
		Config:        router.Config,
		ErrorHelper:   *errorHelper,
		RosterService: router.RosterService,
	}
	mux := http.NewServeMux()
	//mux.HandleFunc("/routes", http.HandlerFunc(GetWithID))
	mux.HandleFunc("/", resourceHandler.get())
	mux.HandleFunc("/rosters", rh.byId())
	mux.HandleFunc("/rosters/", rh.get())
	return router.logHttpRequest(mux)
}
