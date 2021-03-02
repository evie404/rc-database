package database

import (
	"testing"

	"github.com/rickypai/rc-database/testhelpers"
)

func TestDatabase_GetSet(t *testing.T) {
	db := NewDatabase()

	testhelpers.TestGetSet(db, 1000, t)
}

func TestDatabase_ConcurrentGetSet(t *testing.T) {
	db := NewDatabase()

	testhelpers.TestConcurrentGetSet(db, 1000, t)
}
