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

type BsonType int8

const (
    BSON_EOO BsonType = iota
    BSON_DOUBLE
    BSON_STRING
    BSON_OBJECT
    BSON_ARRAY
    BSON_BINDATA
    BSON_UNDEFINED
    BSON_OID
    BSON_BOOL
    BSON_DATE
    BSON_NULL
    BSON_REGEX
    BSON_DBREF /**< Deprecated. */
    BSON_CODE
    BSON_SYMBOL
    BSON_CODEWSCOPE
    BSON_INT
    BSON_TIMESTAMP
    BSON_LONG
)

type Bson struct {
    _bson *C.bson
}

type BsonIterator struct {
    iterator *C.bson_iterator
    // bson     *C.bson
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
    return int(C.bson_append_element(b._bson, C.CString(name_or_null), bson_iterator.iterator))
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

/***************************************
*
* bson_iterator
*
* **************************************/

/**
 * Advance a bson_iterator to the named field.
 *
 * @param it the bson_iterator to use.
 * @param obj the BSON object to use.
 * @param name the name of the field to find.
 *
 * @return the type of the found object or BSON_EOO if it is not found.
 */
// MONGO_EXPORT bson_type bson_find( bson_iterator *it, const bson *obj, const char *name );
func (it *BsonIterator) Find(bson *Bson, name string) BsonType {
    return BsonType(C.bson_find(it.iterator, bson._bson, C.CString(name)))
}

// MONGO_EXPORT bson_iterator* bson_iterator_alloc( void );
// MONGO_EXPORT void bson_iterator_dealloc(bson_iterator*);

// /**
//  * Initialize a bson_iterator.
//  *
//  * @param i the bson_iterator to initialize.
//  * @param bson the BSON object to associate with the iterator.
//  */
// MONGO_EXPORT void bson_iterator_init( bson_iterator *i , const bson *b );
func (it *BsonIterator) Init(bson *Bson) {
    C.bson_iterator_init(it.iterator, bson._bson)
}

// /**
//  * Initialize a bson iterator from a const char* buffer. Note
//  * that this is mostly used internally.
//  *
//  * @param i the bson_iterator to initialize.
//  * @param buffer the buffer to point to.
//  */
// MONGO_EXPORT void bson_iterator_from_buffer( bson_iterator *i, const char *buffer );

// /* more returns true for eoo. best to loop with bson_iterator_next(&it) */
// /**
//  * Check to see if the bson_iterator has more data.
//  *
//  * @param i the iterator.
//  *
//  * @return  returns true if there is more data.
//  */
// MONGO_EXPORT bson_bool_t bson_iterator_more( const bson_iterator *i );

// /**
//  * Point the iterator at the next BSON object.
//  *
//  * @param i the bson_iterator.
//  *
//  * @return the type of the next BSON object.
//  */
// MONGO_EXPORT bson_type bson_iterator_next( bson_iterator *i );
func (it *BsonIterator) Next() BsonType {
    return BsonType(C.bson_iterator_next(it.iterator))
}

// /**
//  * Get the type of the BSON object currently pointed to by the iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return  the type of the current BSON object.
//  */
// MONGO_EXPORT bson_type bson_iterator_type( const bson_iterator *i );
func (it *BsonIterator) Type() BsonType {
    return BsonType(C.bson_iterator_type(it.iterator))
}

// /**
//  * Get the key of the BSON object currently pointed to by the iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the key of the current BSON object.
//  */
// MONGO_EXPORT const char *bson_iterator_key( const bson_iterator *i );
func (it *BsonIterator) Key() string {
    return C.GoString(C.bson_iterator_key(it.iterator))
}

// /**
//  * Get the value of the BSON object currently pointed to by the iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return  the value of the current BSON object.
//  */
// MONGO_EXPORT const char *bson_iterator_value( const bson_iterator *i );
func (it *BsonIterator) Value() string {
    return C.GoString(C.bson_iterator_value(it.iterator))
}

// /* these convert to the right type (return 0 if non-numeric) */
// /**
//  * Get the double value of the BSON object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return  the value of the current BSON object.
//  */
// MONGO_EXPORT double bson_iterator_double( const bson_iterator *i );
func (it *BsonIterator) Double() float64 {
    return float64(C.bson_iterator_double(it.iterator))
}

// /**
//  * Get the int value of the BSON object currently pointed to by the iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return  the value of the current BSON object.
//  */
// MONGO_EXPORT int bson_iterator_int( const bson_iterator *i );
func (it *BsonIterator) Int() int {
    return int(C.bson_iterator_int(it.iterator))
}

// /**
//  * Get the long value of the BSON object currently pointed to by the iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the value of the current BSON object.
//  */
// MONGO_EXPORT int64_t bson_iterator_long( const bson_iterator *i );
func (it *BsonIterator) Long() int64 {
    return int64(C.bson_iterator_long(it.iterator))
}

// /* return the bson timestamp as a whole or in parts */
// /**
//  * Get the timestamp value of the BSON object currently pointed to by
//  * the iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the value of the current BSON object.
//  */
// MONGO_EXPORT bson_timestamp_t bson_iterator_timestamp( const bson_iterator *i );
// MONGO_EXPORT int bson_iterator_timestamp_time( const bson_iterator *i );
// MONGO_EXPORT int bson_iterator_timestamp_increment( const bson_iterator *i );
func (it *BsonIterator) TimestampTime() int {
    return int(C.bson_iterator_timestamp_time(it.iterator))
}
func (it *BsonIterator) TimestampTimeIncrement() int {
    return int(C.bson_iterator_timestamp_increment(it.iterator))
}

// /**
//  * Get the boolean value of the BSON object currently pointed to by
//  * the iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the value of the current BSON object.
//  */
// /* false: boolean false, 0 in any type, or null */
// /* true: anything else (even empty strings and objects) */
// MONGO_EXPORT bson_bool_t bson_iterator_bool( const bson_iterator *i );
func (it *BsonIterator) Bool() bool {
    var b bool
    bb := int(C.bson_iterator_bool(it.iterator))
    if bb == 1 {
        b = true
    }
    return b
}

// /**
//  * Get the double value of the BSON object currently pointed to by the
//  * iterator. Assumes the correct type is used.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the value of the current BSON object.
//  */
// /* these assume you are using the right type */
// double bson_iterator_double_raw( const bson_iterator *i );
func (it *BsonIterator) DoubleRaw() float64 {
    return float64(C.bson_iterator_double_raw(it.iterator))
}

// /**
//  * Get the int value of the BSON object currently pointed to by the
//  * iterator. Assumes the correct type is used.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the value of the current BSON object.
//  */
// int bson_iterator_int_raw( const bson_iterator *i );

// /**
//  * Get the long value of the BSON object currently pointed to by the
//  * iterator. Assumes the correct type is used.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the value of the current BSON object.
//  */
// int64_t bson_iterator_long_raw( const bson_iterator *i );

// *
//  * Get the bson_bool_t value of the BSON object currently pointed to by the
//  * iterator. Assumes the correct type is used.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the value of the current BSON object.

// bson_bool_t bson_iterator_bool_raw( const bson_iterator *i );

// /**
//  * Get the bson_oid_t value of the BSON object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the value of the current BSON object.
//  */
// MONGO_EXPORT bson_oid_t *bson_iterator_oid( const bson_iterator *i );

// /**
//  * Get the string value of the BSON object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return  the value of the current BSON object.
//  */
// /* these can also be used with bson_code and bson_symbol*/
// MONGO_EXPORT const char *bson_iterator_string( const bson_iterator *i );
func (it *BsonIterator) String() string {
    return C.GoString(C.bson_iterator_string(it.iterator))
}

// /**
//  * Get the string length of the BSON object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the length of the current BSON object.
//  */
// int bson_iterator_string_len( const bson_iterator *i );
func (it *BsonIterator) StringLen() int {
    return int(C.bson_iterator_string_len(it.iterator))
}

// /**
//  * Get the code value of the BSON object currently pointed to by the
//  * iterator. Works with bson_code, bson_codewscope, and BSON_STRING
//  * returns NULL for everything else.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the code value of the current BSON object.
//  */
// /* works with bson_code, bson_codewscope, and BSON_STRING */
// /* returns NULL for everything else */
// MONGO_EXPORT const char *bson_iterator_code( const bson_iterator *i );

// /**
//  * Get the code scope value of the BSON object currently pointed to
//  * by the iterator. Calls bson_init_empty on scope if current object is
//  * not BSON_CODEWSCOPE.
//  *
//  * @note When copyData is false, the scope becomes invalid when the
//  *       iterator's data buffer is deallocated. For either value of
//  *       copyData, you must pass the scope object to bson_destroy
//  *       when you are done using it.
//  *
//  * @param i the bson_iterator.
//  * @param scope an uninitialized BSON object to receive the scope.
//  * @param copyData when true, makes a copy of the scope data which will remain
//  *   valid when the iterator's data buffer is deallocated.
//  */
// MONGO_EXPORT void bson_iterator_code_scope_init( const bson_iterator *i, bson *scope, bson_bool_t copyData );

// /**
//  * Get the date value of the BSON object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the date value of the current BSON object.
//  */
// /* both of these only work with bson_date */
// MONGO_EXPORT bson_date_t bson_iterator_date( const bson_iterator *i );

// /**
//  * Get the time value of the BSON object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the time value of the current BSON object.
//  */
// MONGO_EXPORT time_t bson_iterator_time_t( const bson_iterator *i );

// /**
//  * Get the length of the BSON binary object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the length of the current BSON binary object.
//  */
// MONGO_EXPORT int bson_iterator_bin_len( const bson_iterator *i );

// /**
//  * Get the type of the BSON binary object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the type of the current BSON binary object.
//  */
// MONGO_EXPORT char bson_iterator_bin_type( const bson_iterator *i );

// /**
//  * Get the value of the BSON binary object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the value of the current BSON binary object.
//  */
// MONGO_EXPORT const char *bson_iterator_bin_data( const bson_iterator *i );

// /**
//  * Get the value of the BSON regex object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator
//  *
//  * @return the value of the current BSON regex object.
//  */
// MONGO_EXPORT const char *bson_iterator_regex( const bson_iterator *i );

// /**
//  * Get the options of the BSON regex object currently pointed to by the
//  * iterator.
//  *
//  * @param i the bson_iterator.
//  *
//  * @return the options of the current BSON regex object.
//  */
// MONGO_EXPORT const char *bson_iterator_regex_opts( const bson_iterator *i );

// /* these work with BSON_OBJECT and BSON_ARRAY */
// /**
//  * Get the BSON subobject currently pointed to by the
//  * iterator.
//  *
//  * @note When copyData is 0, the subobject becomes invalid when its parent's
//  *       data buffer is deallocated. For either value of copyData, you must
//  *       pass the subobject to bson_destroy when you are done using it.
//  *
//  * @param i the bson_iterator.
//  * @param sub an unitialized BSON object which will become the new subobject.
//  */
// MONGO_EXPORT void bson_iterator_subobject_init( const bson_iterator *i, bson *sub, bson_bool_t copyData );
func (it *BsonIterator) SubObjectInit(sub *Bson, copyData bool) {
    _copyData := 0
    if copyData {
        _copyData = 1
    }
    C.bson_iterator_subobject_init(it.iterator, sub._bson, C.bson_bool_t(_copyData))
}

// /**
//  * Get a bson_iterator that on the BSON subobject.
//  *
//  * @param i the bson_iterator.
//  * @param sub the iterator to point at the BSON subobject.
//  */
// MONGO_EXPORT void bson_iterator_subiterator( const bson_iterator *i, bson_iterator *sub );
func (it *BsonIterator) SubIterator(sub *BsonIterator) {
    C.bson_iterator_subiterator(it.iterator, sub.iterator)
}
