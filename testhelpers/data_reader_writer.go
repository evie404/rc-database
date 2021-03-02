package testhelpers

import (
	"fmt"
	"math/rand"
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
