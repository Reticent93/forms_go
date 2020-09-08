package main

import (
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("./templates/form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := ContactDetails{
			Email: r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		_ = details
		tmpl.Execute(w, struct {Success bool}{true})
	})

	log.Fatal(http.ListenAndServe(":4040", nil))
}
