package libgomongo

import (
    "github.com/couchbaselabs/go.assert"
    "testing"
)

func TestConnPool(t *testing.T) {
    size := 3
    pool := NewPool(host, port, size)
    assert.Equals(t, pool.Host, host)
    assert.Equals(t, pool.Port, port)
    assert.Equals(t, pool.Size, size)

    conn1, err := pool.Get()
    assert.Equals(t, err, nil)
    conn2, err := pool.Get()
    assert.Equals(t, err, nil)
    conn3, err := pool.Get()
    assert.Equals(t, err, nil)
    conn4, err := pool.Get()
    assert.Equals(t, err, nil)
    assert.Equals(t, len(pool.conns), 0)

    conn1.Close()
    assert.Equals(t, len(pool.conns), 1)
    conn2.Close()
    assert.Equals(t, len(pool.conns), 2)
    conn3.Close()
    assert.Equals(t, len(pool.conns), 3)
    conn4.Close()
    assert.Equals(t, len(pool.conns), size)
}
