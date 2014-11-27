package main

import (
	// "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	// "os"
)

type page struct {
	title string
	body  []byte
}

func (p *page) save() error {
	filename := "ss" + ".txt"
	return ioutil.WriteFile(filename, p.body, 0600)
}

func response(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("Hello world."))
}

const lenPath = len("/view/")

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[lenPath:]
	content, _ := ioutil.ReadFile("ss.txt")
	w.Write(content)
	// fmt.Fprintf(w, "<h1>%s</h1><div>view page</div>", title)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	// p, err := loadPage(title)
	p := &page{title: title}
	t, _ := template.ParseFiles("edit.html")
	// t.Execute(p, w)
	t.Execute(w, p)

}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	body := r.FormValue("body")
	p := &page{title: title, body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/", response)

	http.ListenAndServe(":3000", nil)
}
