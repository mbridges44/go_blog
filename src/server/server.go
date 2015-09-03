package main

import (
	"html/template"
	"net/http"
	"regexp"
)

type Server_Config struct {
	Html_templates string
	Port string
}

 var templates *template.Template
 var validPath *regexp.Regexp
 var config  = Server_Config{}

func init() {
	//Parse config file
	ParseJSON("server_config.json", config)
	println(config.Html_templates)

	//Set templates
	templates = template.Must(template.ParseFiles("../web/html_templates/edit.html", "../web/html_templates/view.html"))
	validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
}

func main() {

}

func startServerHandlers(){
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/", redirectHome)
	http.ListenAndServe(":8080", nil)
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

func redirectHome(http.ResponseWriter, *http.Request){
//TODO figure out redirect home stuff
	
}

func editHandler(writer http.ResponseWriter, r *http.Request, title string){
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(writer, "edit", p)
}


func renderTemplate(writer http.ResponseWriter, tmpl string, page *Page){
	err := templates.ExecuteTemplate(writer, tmpl + ".html", page)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(writer http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(writer, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(writer, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}



