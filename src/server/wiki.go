package main
import "strconv"
/*import (
	"io/ioutil"

)*/

type Entry struct {
	entryID int
	Title string
	Body string
	Author string

}

func NewEntry(Title , Body string, Author string) *Entry{
	return &Entry{Title: Title, Body: Body, Author: Author}
}


func (p *Entry) save() {
	getEntry(1)
}


func loadEntry(id int) (*Entry, error) {

	return getEntry(id), nil
}

func loadEntryString(id string) (*Entry, error) {
	i, err := strconv.Atoi(id)
	if(err != nil){
		panic(err.Error());
	}

	return getEntry(i), nil
}

