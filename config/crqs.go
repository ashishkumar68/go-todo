package config

import (
	"sync"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ashishkumar68/go-todo/handler"
)

var (
	facadeConfig cqrs.FacadeConfig
	once         sync.Once
)

func GetFacadeConfig() cqrs.FacadeConfig {
	once.Do(func() {
		logger := watermill.NewStdLogger(false, false)
		cqrsMarshaler := cqrs.JSONMarshaler{}
		router, err := message.NewRouter(message.RouterConfig{}, logger)
		if err != nil {
			panic(err)
		}

		facadeConfig = cqrs.FacadeConfig{
			GenerateCommandsTopic: func(commandName string) string {
				return commandName
			},
			CommandHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
				return []cqrs.CommandHandler{
					handler.CreateTaskHandler{EventBus: eb},
				}
			},
			GenerateEventsTopic: func(eventName string) string {
				return "events"
			},
			Router:                router,
			CommandEventMarshaler: cqrsMarshaler,
			Logger:                logger,
		}
	})

	return facadeConfig
}
