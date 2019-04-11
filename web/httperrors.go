package web

import (
	"fmt"
	"github.com/micklove/simple-roster/internal/app/config"
	"net/http"
	"runtime/debug"
)

type ErrorHelper struct {
	app.Config
}

func (httpErr ErrorHelper) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	httpErr.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func (httpErr ErrorHelper) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (httpErr ErrorHelper) notFound(w http.ResponseWriter) {
	httpErr.clientError(w, http.StatusNotFound)
}
