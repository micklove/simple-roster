package app

import (
	"log"
	"os"
)

type Config struct {
	FileDaoStoreName string
	HttpAddress      string
	ErrorLog         *log.Logger
	InfoLog          *log.Logger
}

func (cfg *Config) SetupLogs() {
	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
