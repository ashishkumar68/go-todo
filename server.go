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
	err := db.Debug().Migrator().DropTable(&model.User{}, &model.Task{})
	if err != nil {
		fmt.Println("Failed to drop table.")
		return
	}
	err = db.Debug().Migrator().CreateTable(&model.User{}, &model.Task{})
	if err != nil {
		fmt.Println("Failed to create table.")
		return
	}
	fmt.Println("Migrations were executed successfully.")
}

func main() {
	port := "8080"
	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", controller.CreateTaskAction).Methods("POST")
	r.HandleFunc("/api/tasks/{taskId}", controller.GetTaskAction).Methods("GET")
	r.HandleFunc("/api/tasks", controller.GetTasksAction).Methods("GET")

	fmt.Println("running server on port", port)
	http.Handle("/", r)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		fmt.Println("Could not start the HTTP server.")
		fmt.Println(err)
	}
}
