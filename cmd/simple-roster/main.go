package main

import (
	"encoding/json"
	"fmt"
	"github.com/micklove/simple-roster/internal/app/model"
)

func main() {
	//note := model.CreateNote("New Note")
	user := model.User{1234, "mick", "love", "micklove", make([]model.Note, 0)}
	user.AddNote(model.CreateNote("Hello World"))
	pretty, _ := json.MarshalIndent(user, "", "\t")
	fmt.Println(string(pretty))
}
