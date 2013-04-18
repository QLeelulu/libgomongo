package libgomongo

import (
    "fmt"
    "github.com/sdegutis/go.assert"
    "testing"
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
        switch conn.ErrNo() {
        case MONGO_CONN_NO_SOCKET:
            fmt.Println("no socket")
        case MONGO_CONN_FAIL:
            fmt.Println("connection failed")
        case MONGO_CONN_NOT_MASTER:
            fmt.Println("not master")
        default:
            fmt.Println("not know error")
        }
    }
    conn.Destroy()

    assert.Equals(t, status, MONGO_OK)
}

func TestInsert(t *testing.T) {
    bson := NewBson()
    bson.Init()
    bson.AppendNewOid("_id")
    bson.AppendString("name", "Joe")
    bson.AppendInt("age", 33)
    bson.Finish()

    conn, status := newClient()
    assert.Equals(t, status, MONGO_OK)

    conn.Insert("libgomongo-test.people", bson, nil)

    bson.Destroy()
    conn.Destroy()
}
