package libgomongo

import (
    "fmt"
    "github.com/couchbaselabs/go.assert"
    "testing"
    "time"
)

var (
    host = "127.0.0.1"
    port = 27017
)

func newClient() (*Mongo, int) {
    conn := NewMongo()
    status := conn.Client(host, port)
    return conn, status
}

func TestConn(t *testing.T) {
    conn := NewMongo()
    status := conn.Client(host, port)
    if status != MONGO_OK {
        fmt.Println(conn.Error())
    }
    conn.Destroy()

    assert.Equals(t, status, MONGO_OK)
}

func TestConnPool(t *testing.T) {
    size := 20
    pool := NewConnPool(host, port, size)
    assert.Equals(t, pool.Host, host)
    assert.Equals(t, pool.Port, port)
    assert.Equals(t, pool.Size, size)

    conn := pool.Get()
    time.Sleep(1 * time.Second)
    assert.Equals(t, len(pool.freeConn), size/2-1)

    conn.Close()
    assert.Equals(t, len(pool.freeConn), size/2)
}
