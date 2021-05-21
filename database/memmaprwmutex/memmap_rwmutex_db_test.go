package memmaprwmutex

import (
	"testing"

	"github.com/rickypai/rc-database/testhelpers"
)

func TestMemoryMapRWMutexDatabase_GetSet(t *testing.T) {
	db := NewMemoryMapRWMutexDatabase()

	testhelpers.TestGetSet(db, 1000, t)
}

func TestMemoryMapRWMutexDatabase_ConcurrentGetSet(t *testing.T) {
	db := NewMemoryMapRWMutexDatabase()

	testhelpers.TestConcurrentGetSet(db, 1000, t)
}

func TestMemoryMapRWMutexDatabase_SetMultipleVersions(t *testing.T) {
	db := NewMemoryMapRWMutexDatabase()

	testhelpers.TestSetMultipleVersions(db, 1000, t)
}

func TestMemoryMapRWMutexDatabase_ConcurrentSetRace(t *testing.T) {
	db := NewMemoryMapRWMutexDatabase()

	testhelpers.TestConcurrentSetRace(db, 1000, t)
}
