package main

import (
	"encoding/json"
	"fmt"
	"github.com/micklove/simple-roster/api/dao/file"
	"github.com/micklove/simple-roster/api/service"
	"github.com/micklove/simple-roster/internal/app"
	"github.com/micklove/simple-roster/internal/app/model"
	"github.com/micklove/simple-roster/internal/pkg/UUID"
)

func main() {

	config := configure()
	var roster *model.Roster
	var err error
	roster, err = config.RosterService.ByID("1JLL9NxeyvqhLnDGcMD8MM20H9p")

	if err != nil {
		panic("Error getting roster response")
	}
	pretty, _ := json.MarshalIndent(roster, "", "\t")
	fmt.Println(string(pretty))
}

//TODO - use functional options here
func configure() *app.Config {
	config := &app.Config{
		RosterService: &service.RosterService{},
	}
	//Choose the implementation(s)
	config.Generator = UUID.KSUUIDGenerator{}
	config.RosterService.RosterDao = dao.FileRosterDao{config}
	config.FileDaoStoreName = "/Users/lovemi/dev/_projects/go/simple-roster/api/dao/file/rosters-test.json"
	return config
}
