package dao

import (
	"github.com/micklove/simple-roster/internal/app/model"
)

type RosterDao interface {
	ByID(ID string) (*model.Roster, error)
	Save(roster *model.Roster) error
	All() []string
}
