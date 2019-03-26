package main

import (
	"encoding/json"
	"fmt"
	"github.com/micklove/simple-roster/internal/app/model"
)

func main() {
	//note := model.CreateNote("New Note")
	avatarUrl := "http://localhost:8080/ronnie.jpg"
	user, _ := model.CreateUser("michael", "love", "mick", avatarUrl)
	//user := model.User{1234, "mick", "love", "micklove", make([]model.Note, 0), nil}
	if err := user.AddNote(model.CreateNote("Hello World")); err != nil {
		fmt.Print("error creating sample user")
	}

	var pretty, _ = json.MarshalIndent(user, "", "\t")
	fmt.Println(string(pretty))

	roster, _ := model.CreateRoster("MyRoster")
	pretty, _ = json.MarshalIndent(roster, "", "\t")
	fmt.Println(string(pretty))
}
