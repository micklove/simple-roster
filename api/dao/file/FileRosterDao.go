package dao

import (
	"encoding/json"
	"fmt"
	"github.com/micklove/simple-roster/internal/app"
	"github.com/micklove/simple-roster/internal/app/model"
	"io/ioutil"
)

type FileRosterDao struct {
	*app.Config
}

func (dao FileRosterDao) ByID(ID string) (roster *model.Roster, err error) {
	//roster, _ = model.CreateRoster("MyRoster")
	fileDaoStoreName := dao.Config.FileDaoStoreName
	var rosters []model.Roster = nil
	if rosters, err = readFileToJson(fileDaoStoreName); err != nil {
		return nil, err
	}
	for _, roster := range rosters {
		if roster.ID == ID {
			return &roster, nil
		}
	}
	return roster, nil
}

func (dao FileRosterDao) Save(roster *model.Roster) error {
	return nil
}

func readFileToJson(fileDaoStoreName string) ([]model.Roster, error) {

	var rosters []model.Roster
	var file []byte
	var err error = nil
	file, err = ioutil.ReadFile(fileDaoStoreName)
	if err = json.Unmarshal([]byte(file), &rosters); err != nil {
		return nil, fmt.Errorf("Error parsing [%v] to Roster struct ", fileDaoStoreName)
	}
	return rosters, nil

}
