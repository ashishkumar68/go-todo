package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ashishkumar68/go-todo/controller"
	"github.com/ashishkumar68/go-todo/model"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func runMigration() {
	db, _ := gorm.Open(mysql.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	db.AutoMigrate(&model.Task{})
	fmt.Println("Migrations were executed successully.")
}

func main() {
	go runMigration()
	port := "8080"
	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", controller.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/api/tasks/{taskId}", controller.GetTaskHandler).Methods("GET")
	r.HandleFunc("/api/tasks", controller.GetTasksList).Methods("GET")

	fmt.Println("running server on port", port)
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
