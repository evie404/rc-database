package memmapmutex

import (
	"sync"
)

type MemoryMapMutexDatabase struct {
	data  map[string][]byte
	mutex sync.Mutex
}

func NewMemoryMapMutexDatabase() *MemoryMapMutexDatabase {
	return &MemoryMapMutexDatabase{
		data: map[string][]byte{},
	}
}

func (db *MemoryMapMutexDatabase) Get(key string) ([]byte, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if value, found := db.data[key]; found {
		return value, nil
	}

	return nil, nil
}

func (db *MemoryMapMutexDatabase) Set(key string, value []byte) error {
	db.mutex.Lock()

	db.data[key] = value

	db.mutex.Unlock()

	return nil
}
