package database

type Database struct {
	data map[string][]byte
}

func NewDatabase() *Database {
	return &Database{
		data: map[string][]byte{},
	}
}

func (db *Database) Get(key string) ([]byte, error) {
	if value, found := db.data[key]; found {
		return value, nil
	}

	return nil, nil
}

func (db *Database) Set(key string, value []byte) error {
	db.data[key] = value

	return nil
}
