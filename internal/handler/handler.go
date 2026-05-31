package handler

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tasks"))

}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create-task"))

}

func AskTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ask-task"))

}
