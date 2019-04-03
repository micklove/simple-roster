package app

import (
	"github.com/micklove/simple-roster/api/service"
	"github.com/micklove/simple-roster/internal/pkg/UUID"
)

type Config struct {
	UUID.Generator
	RosterService    *service.RosterService
	FileDaoStoreName string
}
