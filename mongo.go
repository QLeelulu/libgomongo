package libgomongo

// #cgo CFLAGS: -std=gnu99 -I./mongo-c-driver/src/
// #cgo LDFLAGS: -L./mongo-c-driver/src/ -lmongoc
// #include "mongo.h"
import "C"

// import (
// // "errors"
// // "fmt"
// // "tim
//)

type MongoError int8
type CursorError int8

const (
    MONGO_CONN_SUCCESS      MongoError = iota /**< Connection success! */
    MONGO_CONN_NO_SOCKET                      /**< Could not create a socket. */
    MONGO_CONN_FAIL                           /**< An error occured while calling connect(). */
    MONGO_CONN_ADDR_FAIL                      /**< An error occured while calling getaddrinfo(). */
    MONGO_CONN_NOT_MASTER                     /**< Warning: connected to a non-master node (read-only). */
    MONGO_CONN_BAD_SET_NAME                   /**< Given rs name doesn't match this replica set. */
    MONGO_CONN_NO_PRIMARY                     /**< Can't find primary in replica set. Connection closed. */

    MONGO_IO_ERROR              /**< An error occurred while reading or writing on the socket. */
    MONGO_SOCKET_ERROR          /**< Other socket error. */
    MONGO_READ_SIZE_ERROR       /**< The response is not the expected length. */
    MONGO_COMMAND_FAILED        /**< The command returned with 'ok' value of 0. */
    MONGO_WRITE_ERROR           /**< Write with given write_concern returned an error. */
    MONGO_NS_INVALID            /**< The name for the ns (database or collection) is invalid. */
    MONGO_BSON_INVALID          /**< BSON not valid for the specified op. */
    MONGO_BSON_NOT_FINISHED     /**< BSON object has not been finished. */
    MONGO_BSON_TOO_LARGE        /**< BSON object exceeds max BSON size. */
    MONGO_WRITE_CONCERN_INVALID /**< Supplied write concern object is invalid. */
)

const (
    MONGO_CURSOR_EXHAUSTED  CursorError = iota // The cursor has no more results.
    MONGO_CURSOR_INVALID                       // The cursor has timed out or is not recognized.
    MONGO_CURSOR_PENDING                       // Tailable cursor still alive but no data.
    MONGO_CURSOR_QUERY_FAIL                    // The server returned an '$err' object, indicating query failure.  See conn->lasterrcode and conn->lasterrstr for details.
    MONGO_CURSOR_BSON_ERROR                    // Something is wrong with the BSON provided. See conn->err for details.
)

type Mongo struct {
    conn *C.mongo
}

type MongoWriteConcern struct {
    writeConcern *C.mongo_write_concern
}
