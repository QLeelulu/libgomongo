package libgomongo

// #cgo CFLAGS: -std=gnu99 -I./mongo-c-driver/src/
// #cgo LDFLAGS: -L./mongo-c-driver/src/ -lmongoc
// #include "bson.h"
import "C"

import (
    // "errors"
    // "fmt"
    // "time"
    "unsafe"
)

type Bson struct {
    _bson *C.bson
}

type BsonIterator struct {
    bsonIterator *C.bson_iterator
}

func NewBson() *Bson {
    b := &Bson{}
    b._bson = &C.bson{}
    return b
}

func (b *Bson) Print() {
    C.bson_print(b._bson)
}

func (b *Bson) Init() int {
    return int(C.bson_init(b._bson))
}

func (b *Bson) Finish() int {
    return int(C.bson_finish(b._bson))
}

func (b *Bson) Destroy() {
    C.bson_destroy(b._bson)
}

/**
 * Append a string to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the string.
 * @param str the string to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendString(key, val string) int {
    return int(C.bson_append_string(b._bson, C.CString(key), C.CString(val)))
}

/**
 * Append len bytes of a string to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the string.
 * @param str the string to append.
 * @param len the number of bytes from str to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendStringN(key, val string, _len uint) int {
    return int(C.bson_append_string_n(b._bson, C.CString(key), C.CString(val), C.size_t(_len)))
}

/**
 * Append a previously created bson_oid_t to a bson object.
 *
 * @param b the bson to append to.
 * @param name the key for the bson_oid_t.
 * @param oid the bson_oid_t to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
// MONGO_EXPORT int bson_append_oid( bson *b, const char *name, const bson_oid_t *oid );

/**
 * Append a bson_oid_t to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the bson_oid_t.
 *
 * @return BSON_OK or BSON_ERROR.
 */
// MONGO_EXPORT int bson_append_new_oid( bson *b, const char *name );
func (b *Bson) AppendNewOid(name string) int {
    return int(C.bson_append_new_oid(b._bson, C.CString(name)))
}

/**
 * Append an int to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the int.
 * @param i the int to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendInt(key string, val int) int {
    return int(C.bson_append_int(b._bson, C.CString(key), C.int(val)))
}

/**
 * Append an long to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the long.
 * @param i the long to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendLong(key string, val int64) int {
    return int(C.bson_append_long(b._bson, C.CString(key), C.int64_t(val)))
}

/**
 * Append an double to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the double.
 * @param d the double to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendDouble(key string, val float64) int {
    return int(C.bson_append_double(b._bson, C.CString(key), C.double(val)))
}

/**
 * Append a symbol to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the symbol.
 * @param str the symbol to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendSymbol(key, val string) int {
    return int(C.bson_append_symbol(b._bson, C.CString(key), C.CString(val)))
}

/**
 * Append code to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the code.
 * @param str the code to append.
 * @param len the number of bytes from str to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendCode(name, code string) int {
    return int(C.bson_append_code(b._bson, C.CString(name), C.CString(code)))
}

/**
 * Append len bytes of code to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the code.
 * @param str the code to append.
 * @param len the number of bytes from str to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendCodeN(name, code string, _len uint) int {
    return int(C.bson_append_code_n(b._bson, C.CString(name), C.CString(code), C.size_t(_len)))
}

/**
 * Append code to a bson with scope.
 *
 * @param b the bson to append to.
 * @param name the key for the code.
 * @param str the string to append.
 * @param scope a BSON object containing the scope.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendCodeWScope(name, code string, scope *Bson) int {
    return int(C.bson_append_code_w_scope(b._bson, C.CString(name), C.CString(code), scope._bson))
}

/**
 * Append len bytes of code to a bson with scope.
 *
 * @param b the bson to append to.
 * @param name the key for the code.
 * @param str the string to append.
 * @param len the number of bytes from str to append.
 * @param scope a BSON object containing the scope.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendCodeWScopeN(name, code string, _len uint, scope *Bson) int {
    return int(C.bson_append_code_w_scope_n(b._bson, C.CString(name), C.CString(code), C.size_t(_len), scope._bson))
}

/**
 * Append binary data to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the data.
 * @param type the binary data type.
 * @param str the binary data.
 * @param len the length of the data.
 *
 * @return BSON_OK or BSON_ERROR.
 */
// MONGO_EXPORT int bson_append_binary( bson *b, const char *name, char type, const char *str, size_t len );
func (b *Bson) AppendBinary(name string, _type byte, data []byte, _len uint) int {
    p := (**byte)(unsafe.Pointer(&data))
    p2 := (*C.char)(unsafe.Pointer(*p))
    return int(C.bson_append_binary(b._bson, C.CString(name), C.char(_type), p2, C.size_t(_len)))
}

/**
 * Append a bson_bool_t to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the boolean value.
 * @param v the bson_bool_t to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendBool(name string, v bool) int {
    i := 0
    if v {
        i = 1
    }
    return int(C.bson_append_bool(b._bson, C.CString(name), C.bson_bool_t(i)))
}

/**
 * Append a null value to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the null value.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendNull(name string) int {
    return int(C.bson_append_null(b._bson, C.CString(name)))
}

/**
 * Append an undefined value to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the undefined value.
 *
 * @return BSON_OK or BSON_ERROR.
 */
// MONGO_EXPORT int bson_append_undefined( bson *b, const char *name );
func (b *Bson) AppendUndefined(name string) int {
    return int(C.bson_append_undefined(b._bson, C.CString(name)))
}

/**
 * Append a regex value to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the regex value.
 * @param pattern the regex pattern to append.
 * @param the regex options.
 *
 * @return BSON_OK or BSON_ERROR.
 */
// MONGO_EXPORT int bson_append_regex( bson *b, const char *name, const char *pattern, const char *opts );
func (b *Bson) AppendRegex(name, pattern, opts string) int {
    return int(C.bson_append_regex(b._bson, C.CString(name), C.CString(pattern), C.CString(opts)))
}

/**
 * Append bson data to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the bson data.
 * @param bson the bson object to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendBson(name string, bson *Bson) int {
    return int(C.bson_append_bson(b._bson, C.CString(name), bson._bson))
}

/**
 * Append a BSON element to a bson from the current point of an iterator.
 *
 * @param b the bson to append to.
 * @param name_or_null the key for the BSON element, or NULL.
 * @param elem the bson_iterator.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendElement(name_or_null string, bson_iterator *BsonIterator) int {
    return int(C.bson_append_element(b._bson, C.CString(name_or_null), bson_iterator.bsonIterator))
}

/**
 * Append a bson_timestamp_t value to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the timestampe value.
 * @param ts the bson_timestamp_t value to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
// MONGO_EXPORT int bson_append_timestamp( bson *b, const char *name, bson_timestamp_t *ts );
// MONGO_EXPORT int bson_append_timestamp2( bson *b, const char *name, int time, int increment );

/* these both append a bson_date */
/**
 * Append a bson_date_t value to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the date value.
 * @param millis the bson_date_t to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
// MONGO_EXPORT int bson_append_date( bson *b, const char *name, bson_date_t millis );

/**
 * Append a time_t value to a bson.
 *
 * @param b the bson to append to.
 * @param name the key for the date value.
 * @param secs the time_t to append.
 *
 * @return BSON_OK or BSON_ERROR.
 */
// MONGO_EXPORT int bson_append_time_t( bson *b, const char *name, time_t secs );

/**
 * Start appending a new object to a bson.
 *
 * @param b the bson to append to.
 * @param name the name of the new object.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendStartObject(name string) int {
    return int(C.bson_append_start_object(b._bson, C.CString(name)))
}

/**
 * Start appending a new array to a bson.
 *
 * @param b the bson to append to.
 * @param name the name of the new array.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendStartArray(name string) int {
    return int(C.bson_append_start_array(b._bson, C.CString(name)))
}

/**
 * Finish appending a new object or array to a bson.
 *
 * @param b the bson to append to.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendFinishObject() int {
    return int(C.bson_append_finish_object(b._bson))
}

/**
 * Finish appending a new object or array to a bson. This
 * is simply an alias for bson_append_finish_object.
 *
 * @param b the bson to append to.
 *
 * @return BSON_OK or BSON_ERROR.
 */
func (b *Bson) AppendFinishArray() int {
    return int(C.bson_append_finish_array(b._bson))
}
