package web

import "github.com/micklove/simple-roster/internal/app/model"

// simple wrapper to ensure we get the following json response struct
//{
//	"data": {
//		"roster": {
//			"rosterId": "1JLL9NxeyvqhLnDGcMD8MM20H9p",
//			"start": "2019-04-03T17:27:58.111176+11:00",
//			etc...
//		}
//	}
//}

type RosterWrapper struct {
	Roster *model.Roster `json:"roster"`
}

func NewRosterWrapper(roster *model.Roster) *RosterWrapper {
	return &RosterWrapper{
		Roster: roster,
	}
}
