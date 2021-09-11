package database

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestItUsesOneManagerForAllGoRoutines(t *testing.T) {
	var wg sync.WaitGroup
	var firstInstance *gorm.DB
	mutex := &sync.Mutex{}
	noRoutines := 100
	wg.Add(noRoutines)

	for i := 0; i < noRoutines; i++ {
		go func(currentIndex int) {
			defer wg.Done()
			var manager *gorm.DB
			if firstInstance == nil {
				mutex.Lock()
				defer mutex.Unlock()
				db, err := GetManager()
				if err != nil {
					t.Error("Failed test while trying to fetch first DB instance")
				}
				firstInstance = db
			}
			manager, err := GetManager()
			if err != nil {
				t.Error("Failed test while trying to fetch first DB instance")
			}
			assert.Same(t, firstInstance, manager)
		}(i)
	}

	wg.Wait()
}
