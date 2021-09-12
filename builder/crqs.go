package builder

import (
	"sync"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ashishkumar68/go-todo/config"
)

var (
	cqrsFacade *cqrs.Facade
	once       sync.Once
)

func GetCqrsFacade() (*cqrs.Facade, error) {
	once.Do(func() {
		cf, err := cqrs.NewFacade(config.GetFacadeConfig())
		if err != nil {
			panic(err)
		}
		cqrsFacade = cf
	})

	return cqrsFacade, nil
}
