package main

import (
	"flag"
	"github.com/micklove/simple-roster/internal/app/config"
	"github.com/micklove/simple-roster/internal/app/dao/file"
	"github.com/micklove/simple-roster/internal/app/service"
	"github.com/micklove/simple-roster/internal/pkg/UUID"
	"github.com/micklove/simple-roster/web"
	"log"
	"net/http"
)

//https://stackoverflow.com/questions/28256923/import-cycle-not-allowed

func main() {

	cfg := configure()
	handleArgs(cfg)
	router := &web.Router{
		RosterService: &service.RosterService{
			RosterDao: dao.NewFileRosterDao(cfg),
		},
	}
	mux := router.Routes(cfg)

	// set the ErrorLog field so that the server now uses the custom errorLog logger
	srv := &http.Server{
		Addr:     cfg.HttpAddress,
		ErrorLog: cfg.ErrorLog,
		Handler:  mux,
	}
	log.Fatal(srv.ListenAndServe())
}

//TODO - use functional options here
func configure() *app.Config {
	cfg := &app.Config{
		Generator:        UUID.KSUUIDGenerator{},
		FileDaoStoreName: "/Users/lovemi/dev/_projects/go/simple-roster/internal/app/dao/file/rosters-test.json",
	}
	//Choose the implementation(s)
	//cfg.Generator = UUID.KSUUIDGenerator{}
	//cfg.FileDaoStoreName = "/Users/lovemi/dev/_projects/go/simple-roster/internal/app/dao/file/rosters-test.json"
	cfg.SetupLogs()
	return cfg
}

func handleArgs(cfg *app.Config) {
	addr := flag.String("addr", "0.0.0.0:8080", "HTTP network address")
	flag.Parse()
	log.Printf("param addr = [%v]", *addr)
	cfg.HttpAddress = *addr
}
