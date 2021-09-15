package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ashishkumar68/go-todo/builder"
	"github.com/ashishkumar68/go-todo/command"
	"github.com/ashishkumar68/go-todo/repository"
	"github.com/ashishkumar68/go-todo/response"
)

func CreateTaskAction(w http.ResponseWriter, req *http.Request) {
	bodyContent, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.SendBadRequestContentError(w, []byte("Could not read from request"))
		return
	}
	cmd, err := command.CreateTaskCommandByStreamContent(bodyContent)
	if err != nil {
		response.SendBadRequestContentError(w, []byte("Could not process request"))
		return
	}
	builder.CBus.Execute(cmd)
	task := repository.GetTaskRepository().FindOneByUuid(cmd.Uuid)
	taskBytes, err := json.Marshal(task)
	if err != nil {
		response.SendBadRequestContentError(w, []byte("Could not process request"))
		return
	}
	response.SendCreated(w, taskBytes)
}

func GetTaskAction(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GET an existing task")
}

func GetTasksAction(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GET current tasks list")
}
