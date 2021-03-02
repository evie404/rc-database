package database

import "sync"

type Database struct {
	// TODO: support on-disk persistence
	data  map[string][]byte
	mutex sync.Mutex
}

func NewDatabase() *Database {
	return &Database{
		data: map[string][]byte{},
	}
}

func (db *Database) Get(key string) ([]byte, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if value, found := db.data[key]; found {
		return value, nil
	}

	return nil, nil
}

func (db *Database) Set(key string, value []byte) error {
	db.mutex.Lock()

	db.data[key] = value

	db.mutex.Unlock()

	return nil
}
