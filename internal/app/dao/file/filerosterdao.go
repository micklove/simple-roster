package dao

import (
	"encoding/json"
	"fmt"
	"github.com/micklove/simple-roster/internal/app/model"
	"io/ioutil"
	"log"
)

//TEMP DAO - until persistence layer is chosen
type FileRosterDao struct {
	fileDaoStoreName string
}

func NewFileRosterDao(fileDaoStoreName string) *FileRosterDao {
	return &FileRosterDao{
		fileDaoStoreName: fileDaoStoreName,
	}
}

func (frd *FileRosterDao) ByID(ID string) (roster *model.Roster, err error) {
	//roster, _ = model.CreateRoster("MyRoster")
	var rosters []model.Roster = nil
	if rosters, err = readFileToJson(frd.fileDaoStoreName); err != nil {
		err := fmt.Errorf("Error getting Roster with id [%v] Error [%v] ", ID, err)
		log.Println(err)
		return nil, err
	}
	for _, roster := range rosters {
		if roster.ID == ID {
			return &roster, nil
		}
	}
	return roster, nil
}

func (frd *FileRosterDao) All() []string {
	panic("FileRosterDao.All method not implemented")
}

func (frd *FileRosterDao) Save(roster *model.Roster) error {
	panic("FileRosterDao.Save method not implemented")
}

func readFileToJson(fileDaoStoreName string) (rosters []model.Roster, err error) {
	//var rosters []model.Roster
	var file []byte
	//var err error = nil
	if file, err = ioutil.ReadFile(fileDaoStoreName); err != nil {

		err := fmt.Errorf("Error reading file, Error [%v] ", err)
		log.Println(err)
		return nil, err
	}
	if err = json.Unmarshal([]byte(file), &rosters); err != nil {
		return nil, fmt.Errorf("Error parsing [%v] to Roster struct ", fileDaoStoreName)
	}
	return rosters, nil

}
