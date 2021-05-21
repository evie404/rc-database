package jsonfile

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/rickypai/rc-database/testhelpers"
	"github.com/stretchr/testify/require"
)

func TestJSONFileDatabase_GetSet(t *testing.T) {
	db, err := NewJSONFileDatabase("/tmp/test.json")
	require.NoError(t, err)

	defer os.Remove("/tmp/test.json")

	testhelpers.TestGetSet(db, 1000, t)
}

func TestJSONFileDatabase_ConcurrentGetSet(t *testing.T) {
	db, err := NewJSONFileDatabase("/tmp/test.json")
	require.NoError(t, err)

	defer os.Remove("/tmp/test.json")

	testhelpers.TestConcurrentGetSet(db, 1000, t)
}

func TestJSONFileDatabase_ConcurrentSetRace(t *testing.T) {
	db, err := NewJSONFileDatabase("/tmp/test.json")
	require.NoError(t, err)

	defer os.Remove("/tmp/test.json")

	testhelpers.TestConcurrentSetRace(db, 1000, t)
}

func TestJSONFileDatabase_SetMultipleVersions(t *testing.T) {
	db, err := NewJSONFileDatabase("/tmp/test.json")
	require.NoError(t, err)

	defer os.Remove("/tmp/test.json")

	testhelpers.TestSetMultipleVersions(db, 1000, t)
}

func TestJSONFileDatabase_FilePersistence(t *testing.T) {
	db, err := NewJSONFileDatabase("/tmp/test.json")
	require.NoError(t, err)

	defer os.Remove("/tmp/test.json")

	testKeyValues := map[string][]byte{}

	times := 1000

	// seed with some random data and make sure they are accessible right after we write them
	for i := 0; i < times; i++ {
		// prefix with index to prevent the rare case of collision
		key := fmt.Sprintf("%v-%s", i, testhelpers.RandAlphanumericString(rand.Intn(100)+1))
		value := testhelpers.RandAlphanumericBytes(rand.Intn(100) + 1)

		testKeyValues[key] = value

		err := db.Set(key, value)
		require.NoError(t, err)

		gotValue, err := db.Get(key)
		require.NoError(t, err)
		require.Equal(t, value, gotValue)
	}

	// TODO: properly close the database

	// create a new instance of the database
	db1, err := NewJSONFileDatabase("/tmp/test.json")
	require.NoError(t, err)

	// make sure individual keys are still accessible after we finish writing all data
	for key, value := range testKeyValues {
		gotValue, err := db1.Get(key)
		require.NoError(t, err)
		require.Equal(t, value, gotValue)
	}
}
