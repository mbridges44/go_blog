package main
import (
	"net/http"
	"io/ioutil"
)

type content_handler struct {
	Content_type string
}

func (contentHandler *content_handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	//REMOVE HARDCODING
	content_bytes, err := ioutil.ReadFile("src/web" + r.RequestURI)
	if(err != nil){
		panic(err.Error());
	}

	w.Header().Set("Content-Type", contentHandler.Content_type)
	w.Write(content_bytes)
}