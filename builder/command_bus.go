package builder

import (
	"sync"

	"github.com/ashishkumar68/go-todo/command"
	"github.com/ashishkumar68/go-todo/handler"
	"github.com/gogolfing/cbus"
)

var (
	CBus *cbus.Bus
	once sync.Once
)

func GetCommandBus() *cbus.Bus {
	once.Do(func() {
		CBus = &cbus.Bus{}
	})

	return CBus
}

func init() {
	GetCommandBus()

	CBus.Handle(&command.CreateTaskCommand{}, handler.GetTaskHandler())
}
