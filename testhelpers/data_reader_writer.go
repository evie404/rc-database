package testhelpers

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

type DataReaderWriter interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

func TestGetSet(db DataReaderWriter, times int, t *testing.T) {
	testKeyValues := map[string][]byte{}

	// seed with some random data and make sure they are accessible right after we write them
	for i := 0; i < times; i++ {
		// prefix with index to prevent the rare case of collision
		key := fmt.Sprintf("%v-%s", i, RandAlphanumericString(rand.Intn(100)+1))
		value := RandAlphanumericBytes(rand.Intn(100) + 1)

		testKeyValues[key] = value

		err := db.Set(key, value)
		require.NoError(t, err)

		gotValue, err := db.Get(key)
		require.NoError(t, err)
		require.Equal(t, value, gotValue)
	}

	// make sure individual keys are still accessible after we finish writing all data
	for key, value := range testKeyValues {
		gotValue, err := db.Get(key)
		require.NoError(t, err)
		require.Equal(t, value, gotValue)
	}
}

func TestConcurrentGetSet(db DataReaderWriter, times int, t *testing.T) {
	// TODO: customize concurrency

	testKeyValues := map[string][]byte{}
	mutex := sync.Mutex{}

	setWG := sync.WaitGroup{}

	// seed with some random data and make sure they are accessible right after we write them
	for i := 0; i < times; i++ {
		go func(i int) {
			setWG.Add(1)
			defer setWG.Done()

			// prefix with index to prevent the rare case of collision
			key := fmt.Sprintf("%v-%s", i, RandAlphanumericString(rand.Intn(100)+1))
			value := RandAlphanumericBytes(rand.Intn(100) + 1)

			mutex.Lock()
			testKeyValues[key] = value
			mutex.Unlock()

			err := db.Set(key, value)
			require.NoError(t, err)

			gotValue, err := db.Get(key)
			require.NoError(t, err)
			require.Equal(t, value, gotValue)
		}(i)
	}

	setWG.Wait()

	getWG := sync.WaitGroup{}

	// make sure individual keys are still accessible after we finish writing all data
	for key, value := range testKeyValues {
		go func(key string, value []byte) {
			getWG.Add(1)
			defer getWG.Done()

			gotValue, err := db.Get(key)
			require.NoError(t, err)
			require.Equal(t, value, gotValue)
		}(key, value)
	}

	getWG.Wait()
}
