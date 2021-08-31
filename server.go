package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func createTaskHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "POST a new task")
}

func getTaskHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GET an existing task")
}

func getTasksList(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GET current tasks list")
}

func main() {
	port := "8080"
	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", createTaskHandler).Methods("POST")
	r.HandleFunc("/api/tasks/{taskId}", getTaskHandler).Methods("GET")
	r.HandleFunc("/api/tasks", getTasksList).Methods("GET")

	fmt.Println("running server on port", port)
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
