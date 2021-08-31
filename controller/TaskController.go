package controller

import (
	"io"
	"net/http"
)

func CreateTaskHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "POST a new task")
}

func GetTaskHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GET an existing task")
}

func GetTasksList(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GET current tasks list")
}
