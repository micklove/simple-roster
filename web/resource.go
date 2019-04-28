package web

import "fmt"

//Use the enum pattern to declare
type Resource string

const (
	ROSTER Resource = "roster"
	SELF   Resource = "self"
)

type Resources []Resource

var ResourceList = Resources{
	ROSTER,
}

func DumpResources() {
	for _, resoure := range ResourceList {
		fmt.Println(resoure)
	}
}
