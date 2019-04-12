package model

import (
	"errors"
	"fmt"
	"github.com/segmentio/ksuid"
	"net/url"
	"strings"
	"time"
)

type User struct {
	ID          string  `json:"id"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	DisplayName string  `json:"displayName"`
	Notes       Notes   `json:"notes"`
	AvatarUrl   url.URL `json:"avatarUrl"` //do not display
	Avatar      string  `json:"avatar"`
}

func NewUser(firstName string, lastName string, displayName string, avatar string) (*User, error) {
	var err error
	var avatarUrl *url.URL
	avatarUrl, err = getUrlForAvatar(avatar)
	if err != nil {
		return nil, fmt.Errorf("error parsing url [%v] for avatar, error [%v]", avatar, err)
	}
	var id ksuid.KSUID
	if id, err = ksuid.NewRandomWithTime(time.Now()); err != nil {
		fmt.Printf("error getting new UUID for User ID, err [%v]", err)
		return nil, err
	}
	return &User{
		ID:          id.String(),
		FirstName:   firstName,
		LastName:    lastName,
		DisplayName: displayName,
		Notes:       make([]Note, 0),
		AvatarUrl:   *avatarUrl,
		Avatar:      avatar,
	}, err
}

//TODO - Test
func getUrlForAvatar(urlStr string) (*url.URL, error) {
	if len(urlStr) > 0 {
		var parsedUrl *url.URL
		var err error
		if parsedUrl, err = url.ParseRequestURI(urlStr); err != nil {
			return nil, err
		}
		return parsedUrl, nil

	}
	return nil, nil
}

func (user *User) AddNote(note *Note) error {
	var notes Notes
	notes, err := user.Notes.AddNote(note)
	if err != nil {
		return errors.New("error adding note to User")
	}
	user.Notes = notes
	return nil
}

//TODO - provide shift predicate matchers, for start time / end time

//User Matchers - for FilterByUser methods
func (user User) HasDisplayName(searchNamePrefix string) bool {
	return strings.HasPrefix(user.DisplayName, searchNamePrefix)
}

func (user User) HasFirstName(searchNamePrefix string) bool {
	return strings.HasPrefix(user.FirstName, searchNamePrefix)
}

func (user User) HasLastName(searchNamePrefix string) bool {
	return strings.HasPrefix(user.LastName, searchNamePrefix)
}

func (user User) MatchesAnyName(searchNamePrefix string) bool {
	return strings.HasPrefix(user.FirstName, searchNamePrefix) ||
		strings.HasPrefix(user.LastName, searchNamePrefix) ||
		strings.HasPrefix(user.DisplayName, searchNamePrefix)
}
