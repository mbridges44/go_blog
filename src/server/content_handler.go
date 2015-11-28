package main
import (
	"net/http"
	"io/ioutil"
)

type content_handler struct {
	Content_type string
}

func (contentHandler *content_handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	css_bytes, _ := ioutil.ReadFile(".." + r.RequestURI)
	w.Header().Set("Content-Type", contentHandler.Content_type)
	w.Write(css_bytes)
}