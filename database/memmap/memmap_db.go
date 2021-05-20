package memmap

type MemoryMapDatabase struct {
	data map[string][]byte
}

func NewMemoryMapDatabase() *MemoryMapDatabase {
	return &MemoryMapDatabase{
		data: map[string][]byte{},
	}
}

func (db *MemoryMapDatabase) Get(key string) ([]byte, error) {
	if value, found := db.data[key]; found {
		return value, nil
	}

	return nil, nil
}

func (db *MemoryMapDatabase) Set(key string, value []byte) error {
	db.data[key] = value

	return nil
}
