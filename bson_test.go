package libgomongo

import (
    // "fmt"
    "github.com/sdegutis/go.assert"
    "testing"
)

const (
    BSON_OK    = 0
    BSON_ERROR = -1
)

func TestNewBson(t *testing.T) {
    b := NewBson()
    assert.NotEquals(t, b, nil)
    assert.NotEquals(t, b._bson, nil)
}

func TestBsonInit(t *testing.T) {
    b := NewBson()
    st := b.Init()
    assert.Equals(t, st, BSON_OK)
}

func TestBsonFinish(t *testing.T) {
    b := NewBson()

    st := b.Init()
    assert.Equals(t, st, BSON_OK)
    st = b.Finish()
    assert.Equals(t, st, BSON_OK)
}

func TestBson(t *testing.T) {
    b := NewBson()

    st := b.Init()
    assert.Equals(t, st, BSON_OK)

    st = b.AppendNewOid("_id")
    assert.Equals(t, st, BSON_OK)

    st = b.AppendString("name", "libgomongo")
    assert.Equals(t, st, BSON_OK)

    st = b.AppendStringN("name2", "libgomongo", 2)
    assert.Equals(t, st, BSON_OK)

    st = b.AppendInt("int", 11)
    assert.Equals(t, st, BSON_OK)

    st = b.AppendInt("long", 111111)
    assert.Equals(t, st, BSON_OK)

    st = b.AppendDouble("Double", 11.111)
    assert.Equals(t, st, BSON_OK)

    st = b.AppendSymbol("Symbol", "symbol")
    assert.Equals(t, st, BSON_OK)

    st = b.AppendCode("Code", "code")
    assert.Equals(t, st, BSON_OK)

    st = b.AppendCodeN("CodeN", "code", 2)
    assert.Equals(t, st, BSON_OK)

    scope := NewBson()
    scope.Init()
    b.AppendString("name", "scope")
    scope.Finish()
    st = b.AppendCodeWScope("CodeWScope", "code", scope)
    assert.Equals(t, st, BSON_OK)
    st = b.AppendCodeWScopeN("CodeWScopeN", "code", 2, scope)
    assert.Equals(t, st, BSON_OK)

    st = b.AppendBinary("Binary", 1, []byte("Binary"), 2)
    assert.Equals(t, st, BSON_OK)

    st = b.AppendBool("Bool", true)
    assert.Equals(t, st, BSON_OK)
    st = b.AppendBool("Bool2", false)
    assert.Equals(t, st, BSON_OK)

    st = b.AppendNull("NULL")
    assert.Equals(t, st, BSON_OK)

    st = b.AppendUndefined("Undefined")
    assert.Equals(t, st, BSON_OK)

    st = b.AppendRegex("Regex", "\\d+", "g")
    assert.Equals(t, st, BSON_OK)

    st = b.AppendBson("Bson", scope)
    assert.Equals(t, st, BSON_OK)

    // st = b.AppendElement(name_or_null, bson_iterator)
    // assert.Equals(t, st, BSON_OK)

    st = b.AppendStartObject("StartObject")
    assert.Equals(t, st, BSON_OK)
    st = b.AppendString("newkey", "sub-obj")
    assert.Equals(t, st, BSON_OK)
    st = b.AppendFinishObject()
    assert.Equals(t, st, BSON_OK)

    st = b.AppendStartArray("StartArray")
    assert.Equals(t, st, BSON_OK)
    st = b.AppendString("0", "item0")
    assert.Equals(t, st, BSON_OK)
    st = b.AppendString("1", "item1")
    assert.Equals(t, st, BSON_OK)
    st = b.AppendString("2", "item2")
    assert.Equals(t, st, BSON_OK)
    st = b.AppendFinishArray()
    assert.Equals(t, st, BSON_OK)

    st = b.Finish()
    assert.Equals(t, st, BSON_OK)

    b.Print()

    scope.Destroy()
    b.Destroy()
}
