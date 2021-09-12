package config

import (
	"sync"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ashishkumar68/go-todo/handler"
)

var (
	facadeConfig cqrs.FacadeConfig
	once         sync.Once
)

func GetFacadeConfig() cqrs.FacadeConfig {
	once.Do(func() {
		facadeConfig = cqrs.FacadeConfig{
			CommandHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
				return []cqrs.CommandHandler{
					handler.CreateTaskHandler{EventBus: eb},
				}
			},
		}
	})

	return facadeConfig
}
