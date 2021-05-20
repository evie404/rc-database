package memmaprwmutex

import (
	"sync"
)

type MemoryMapRWMutexDatabase struct {
	data  map[string][]byte
	mutex sync.RWMutex
}

func NewMemoryMapRWMutexDatabase() *MemoryMapRWMutexDatabase {
	return &MemoryMapRWMutexDatabase{
		data: map[string][]byte{},
	}
}

func (db *MemoryMapRWMutexDatabase) Get(key string) ([]byte, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	if value, found := db.data[key]; found {
		return value, nil
	}

	return nil, nil
}

func (db *MemoryMapRWMutexDatabase) Set(key string, value []byte) error {
	db.mutex.Lock()

	db.data[key] = value

	db.mutex.Unlock()

	return nil
}
