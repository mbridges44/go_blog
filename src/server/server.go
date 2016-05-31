package main

import (
	"html/template"
	"net/http"
	"regexp"
	"github.com/gorilla/mux"
	"fmt"
	"parser"
	"os"
)

type Server_Config struct {
	Html_templates string
	Static_html string
	Port string
	Javascript_loc string
	Css_loc string
}

var templates *template.Template
var validPath *regexp.Regexp
var config Server_Config


func init() {
	//Parse config file
	parser.ParseJSON("src/server/server_config.json", &config)

	//Set templates
	templates = template.Must(template.ParseGlob(config.Html_templates))
	validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
}

func main() {
	startServerHandlers();
}

func startServerHandlers(){
	cssHandler := &content_handler{Content_type: "text/css"}
	jsHandler := &content_handler{Content_type: "application/javascript"}

	if _, err := os.Stat("src/web/html_static/css"); os.IsNotExist(err) {
		fmt.Printf("DOES NOT EXIST")
	}


	//http.Handle("src/static/", http.FileServer(http.Dir("src/web/js")))
	//http.Handle("web/js/", http.FileServer(http.Dir("web/js/")))
	//http.Handle("/css/", http.FileServer(http.Dir("/src/web/html_static/css/")))
	//cssHandler := http.FileServer(http.Dir("/src/web/html_static/css/"))
	//http.Handle("src/web/html_static/css/", http.StripPrefix("src/web/html_static/css/", cssHandler))
	//http.Handle("/web/html_static/css/", http.StripPrefix("/web/html_static/css/", cssHandler))
	//http.Handle("/html_static/css/", http.StripPrefix("/html_static/css/", http.FileServer(http.Dir("src/web/html_static/css/"))))

	//http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	//fs := http.FileServer(http.Dir("/src/web"))

	r := mux.NewRouter();
	//r.Handle("/", fs)
	r.HandleFunc("/view/{[0-9]+}", makeHandler(viewHandler))
	r.HandleFunc("/edit/{[0-9]+}", makeHandler(editHandler))
	r.HandleFunc("/", redirectHome)
	r.Handle("/html_static/css/{.css}", cssHandler)
	r.Handle("/html_static/js/{.js}", jsHandler)
	r.Handle("/lib/ckeditor/{.js}", jsHandler)

	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":" + config.Port, r)
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request){
		title := validPath.FindStringSubmatch(request.URL.Path)
		if title == nil {
			http.NotFound(writer, request)
			return
		}
		fn(writer, request, title[2])
	}
}

//TODO add more logic to account for other home settings
func redirectHome(writer http.ResponseWriter, req *http.Request){
	//http.Redirect(writer, req, (config.Static_html + "/index.html"), http.StatusOK)

	fmt.Printf("Trying to hit %s\n from %s\n", req.RequestURI, req.Proto)
	http.ServeFile(writer, req, config.Static_html + "/index.html")
}


func notFound(writer http.ResponseWriter, req *http.Request){
	redirectHome(writer, req)
}

func editHandler(writer http.ResponseWriter, r *http.Request, title string){
	p, err := loadEntry(1)

	if err != nil {
		p = &Entry{Title: title}
	} else {
		//DO SOMETHING WITH ERROR, redirect maybe
		fmt.Printf("No entry with id")
	}

	renderTemplate(writer, "edit", p)
}


func renderTemplate(writer http.ResponseWriter, tmpl string, Entry *Entry){
	err := templates.ExecuteTemplate(writer, tmpl + ".html", Entry)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(writer http.ResponseWriter, r *http.Request, title string) {
	p, err := loadEntryString(title)

	if err != nil {
		http.Redirect(writer, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(writer, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
/*	body := r.FormValue("body")
	p := &Entry{Title: title, Body: []byte(body)}
	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)*/
}



