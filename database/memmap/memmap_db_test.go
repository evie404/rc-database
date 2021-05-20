package memmap

import (
	"testing"

	"github.com/rickypai/rc-database/testhelpers"
)

func TestMemoryMapDatabase_GetSet(t *testing.T) {
	db := NewMemoryMapDatabase()

	testhelpers.TestGetSet(db, 1000, t)
}

func TestMemoryMapDatabase_ConcurrentGetSet(t *testing.T) {
	t.Skip("does not support concurrent get/set")
	// TODO: write test to show concurrency fails

	db := NewMemoryMapDatabase()

	testhelpers.TestConcurrentGetSet(db, 1000, t)
}
