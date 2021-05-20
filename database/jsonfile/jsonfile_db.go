package jsonfile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type JSONFileDatabase struct {
	data  map[string][]byte
	mutex sync.Mutex

	filePath string
}

func NewJSONFileDatabase(filePath string) (*JSONFileDatabase, error) {
	var file *os.File
	var err error

	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			return nil, fmt.Errorf("creating file: %v", err)
		}

		file.Close()
	}

	jsonBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("reading file: %v", err)
	}

	data := map[string][]byte{}
	if len(jsonBytes) > 0 {
		err = json.Unmarshal(jsonBytes, &data)
		if err != nil {
			return nil, fmt.Errorf("unmarshalling json: %v", err)
		}
	}

	return &JSONFileDatabase{
		data:     data,
		filePath: filePath,
	}, nil
}

func (db *JSONFileDatabase) Get(key string) ([]byte, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if value, found := db.data[key]; found {
		return value, nil
	}

	return nil, nil
}

func (db *JSONFileDatabase) Set(key string, value []byte) error {
	db.mutex.Lock()

	db.data[key] = value

	// TODO: handle case where writes fail but memory succeeds

	jsonBytes, err := json.Marshal(db.data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(db.filePath, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("writing file: %v", err)
	}

	db.mutex.Unlock()

	return nil
}
