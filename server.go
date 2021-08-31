package main

import (
	"fmt"
	"net/http"

	"github.com/ashishkumar68/go-todo/controller"
	"github.com/gorilla/mux"
)

func main() {
	port := "8080"
	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", controller.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/api/tasks/{taskId}", controller.GetTaskHandler).Methods("GET")
	r.HandleFunc("/api/tasks", controller.GetTasksList).Methods("GET")

	fmt.Println("running server on port", port)
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
