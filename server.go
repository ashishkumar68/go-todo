package main

import (
	"fmt"
	"net/http"

	"github.com/ashishkumar68/go-todo/controller"
	"github.com/ashishkumar68/go-todo/database"
	"github.com/ashishkumar68/go-todo/model"
	"github.com/gorilla/mux"
)

func init() {
	go runMigration()
}

func runMigration() {
	db, _ := database.GetManager()
	db.Migrator().DropTable(&model.Task{})
	db.Migrator().CreateTable(&model.Task{})

	fmt.Println("Migrations were executed successully.")
}

func main() {
	port := "8080"
	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", controller.CreateTaskAction).Methods("POST")
	r.HandleFunc("/api/tasks/{taskId}", controller.GetTaskAction).Methods("GET")
	r.HandleFunc("/api/tasks", controller.GetTasksAction).Methods("GET")

	fmt.Println("running server on port", port)
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
