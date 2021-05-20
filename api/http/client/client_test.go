package client

import (
	"testing"
	"time"

	"github.com/rickypai/rc-database/api/http/server"
	"github.com/rickypai/rc-database/database/memmapmutex"
	"github.com/rickypai/rc-database/testhelpers"
	"github.com/stretchr/testify/require"
)

func TestClient_GetSet(t *testing.T) {
	db := memmapmutex.NewMemoryMapMutexDatabase()

	srv := server.NewDatabaseServer(db)
	port, err := testhelpers.GetUnusedPort()
	require.NoError(t, err)

	// start a server
	go func() {
		err := srv.Listen(port)
		require.NoError(t, err)
	}()

	// wait for server to start
	time.Sleep(time.Second)

	client, err := NewClient("127.0.0.1", port)
	require.NoError(t, err)

	testhelpers.TestGetSet(client, 1000, t)
}

func TestClient_ConcurrentGetSet(t *testing.T) {
	db := memmapmutex.NewMemoryMapMutexDatabase()

	srv := server.NewDatabaseServer(db)
	port, err := testhelpers.GetUnusedPort()
	require.NoError(t, err)

	// start a server
	go func() {
		err := srv.Listen(port)
		require.NoError(t, err)
	}()

	// wait for server to start
	time.Sleep(time.Second)

	client, err := NewClient("127.0.0.1", port)
	require.NoError(t, err)

	// limit concurrency due to server file descriptor limits
	testhelpers.TestConcurrentGetSet(client, 20, t)
}
