package dao

import (
	"github.com/micklove/simple-roster/internal/app/model"
)

type MockRosterDao struct {
	*model.Roster
}

func (rd *MockRosterDao) ByID(ID string) (roster *model.Roster, err error) {

	//Allow users to set "not found" expectation for Response
	if rd.Roster.ID == "NOT FOUND" {
		return nil, err
	}
	if rd.Roster == nil {
		panic("Initialise the mock by calling mock.SetMockResponse")
	}
	if rd.Roster.ID == ID {
		return rd.Roster, nil
	}
	return rd.Roster, nil
}

func (frd *MockRosterDao) All() []string {
	panic("FileRosterDao.All method not implemented")
}

func (rd *MockRosterDao) Save(roster *model.Roster) error {
	return nil
}

func (rd *MockRosterDao) SetMockResponse(response *model.Roster) *model.Roster {
	rd.Roster = response
	return rd.Roster
}
