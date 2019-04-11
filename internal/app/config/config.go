package app

import (
	"github.com/micklove/simple-roster/internal/app/service"
	"github.com/micklove/simple-roster/internal/pkg/UUID"
	"log"
	"os"
)

type Config struct {
	UUID.Generator
	RosterService    *service.RosterService
	FileDaoStoreName string
	HttpAddress      string
	ErrorLog         *log.Logger
	InfoLog          *log.Logger
}

func (cfg *Config) SetupLogs() {
	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
