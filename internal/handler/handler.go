package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
	if err = tmpl.Execute(w, ""); err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/tasks.html")
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
	if err = tmpl.Execute(w, ""); err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/create-task.html")
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
	if err = tmpl.Execute(w, ""); err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}

}

func AskTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ask-task"))

}
