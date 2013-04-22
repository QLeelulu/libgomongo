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
        fmt.Println(conn.Error())
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

    p2 := M{
        "name": "GoLang",
        "age":  18,
    }
    coll := conn.Db("libgomongo-test").C("people")
    status, err := coll.Insert(p2, nil)
    assert.Equals(t, err, nil)
    assert.Equals(t, status, MONGO_OK)

    conn.Destroy()
}

func TestFind(t *testing.T) {
    conn, status := newClient()
    assert.Equals(t, status, MONGO_OK)
    defer conn.Destroy()

    db := conn.Db("libgomongo-test")
    col := db.C("people")
    q := M{
        "name": "Joe",
        "age": M{
            "$gt": 31,
        },
    }

    query := col.Find(q).Fields(M{"name": 1, "_id": 0})
    cur, err := query.Cursor()
    assert.Equals(t, err, nil)
    assert.NotEquals(t, cur, nil)
    if cur != nil {
        defer cur.Destroy()
    }
    assert.Equals(t, cur.Next(), MONGO_OK)
}

func TestCount(t *testing.T) {
    conn, status := newClient()
    assert.Equals(t, status, MONGO_OK)
    defer conn.Destroy()

    db := conn.Db("libgomongo-test")
    col := db.C("people")
    q := M{
        "name": "Joe",
        "age": M{
            "$gt": 31,
        },
    }

    count, err := col.Count(q)
    assert.Equals(t, err, nil)
    assert.Equals(t, count, int64(1))
}

func TestRemove(t *testing.T) {
    conn, status := newClient()
    assert.Equals(t, status, MONGO_OK)
    defer conn.Destroy()

    db := conn.Db("libgomongo-test")
    col := db.C("people")
    q := M{
        "name": "Joe",
    }

    status, err := col.Remove(q, nil)
    assert.Equals(t, err, nil)
    assert.Equals(t, status, MONGO_OK)

    count, err := col.Count(nil)
    assert.Equals(t, err, nil)
    assert.Equals(t, count, int64(1))

    status, err = col.Remove(nil, nil)
    assert.Equals(t, err, nil)
    assert.Equals(t, status, MONGO_OK)

    count, err = col.Count(nil)
    assert.Equals(t, err, nil)
    assert.Equals(t, count, int64(0))
}
