package controller

import (
	"io"
	"net/http"
)

func CreateTaskAction(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "POST a new task")
}

func GetTaskAction(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GET an existing task")
}

func GetTasksAction(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GET current tasks list")
}
