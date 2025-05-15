// currently using go's own webserver tutorial as a base 
// available at https://go.dev/doc/articles/wiki/

package main

import (
    // "fmt"
    "os"
    "log"
    "net/http"
    "html/template"
)

type Page struct {
    Title string
    Body []byte // byte slice instead of string because of unix's c-isms
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}


func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)

    if err != nil {
        return nil, err
    }

    return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles("templates/" + tmpl + ".html")
    t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]

    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }

    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]

    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }

    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page { Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
    addr := os.Getenv("ADDR")
    if len(addr) == 0 {
        addr = "localhost"
    }

    port := ":8000"

    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)

    log.Printf("Server is listening at %s%s...", addr, port)
    log.Fatal(http.ListenAndServe(port, nil))
}

