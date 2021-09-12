package controller

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ashishkumar68/go-todo/builder"
	"github.com/ashishkumar68/go-todo/command"
	"github.com/ashishkumar68/go-todo/response"
)

func CreateTaskAction(w http.ResponseWriter, req *http.Request) {
	cqrs, err := builder.GetCqrsFacade()
	if err != nil {
		response.SendInternalServerError(w)
		return
	}
	bodyContent, err := ioutil.ReadAll(req.Body)
	if err != nil {
		_ = response.SendBadRequestContentError(w, []byte("Could not read from request"))
		return
	}
	cmd, err := command.CreateTaskCommandByStreamContent(bodyContent)
	if err != nil {
		_ = response.SendBadRequestContentError(w, []byte("Could not process request"))
		return
	}
	cqrs.CommandBus().Send(context.Background(), cmd)

	io.WriteString(w, "POST a new task")
}

func GetTaskAction(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GET an existing task")
}

func GetTasksAction(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "GET current tasks list")
}
