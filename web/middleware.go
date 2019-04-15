package web

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func (router *Router) logHttpRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router.Config.InfoLog.Printf("MICK %s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		// dump a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("\n# REQ -----------\n " + string(requestDump) + "# REQ END -----------\n\n")
		next.ServeHTTP(w, r)
	})
}
