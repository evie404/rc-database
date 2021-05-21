package memmapmutex

import (
	"testing"

	"github.com/rickypai/rc-database/testhelpers"
)

func TestMemoryMapMutexDatabase_GetSet(t *testing.T) {
	db := NewMemoryMapMutexDatabase()

	testhelpers.TestGetSet(db, 1000, t)
}

func TestMemoryMapMutexDatabase_ConcurrentGetSet(t *testing.T) {
	db := NewMemoryMapMutexDatabase()

	testhelpers.TestConcurrentGetSet(db, 1000, t)
}

func TestMemoryMapMutexDatabase_SetMultipleVersions(t *testing.T) {
	db := NewMemoryMapMutexDatabase()

	testhelpers.TestSetMultipleVersions(db, 1000, t)
}

func TestMemoryMapMutexDatabase_ConcurrentSetRace(t *testing.T) {
	db := NewMemoryMapMutexDatabase()

	testhelpers.TestConcurrentSetRace(db, 1000, t)
}
