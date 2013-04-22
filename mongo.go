package libgomongo

// #cgo CFLAGS: -std=gnu99 -I./mongo-c-driver/src/
// #cgo LDFLAGS: -L./mongo-c-driver/src/ -lmongoc
// #include "mongo.h"
import "C"

import (
    "errors"
    // "fmt"
    // "tim
)

type MongoError int8
type CursorError int8

const (
    MONGO_OK    = 0
    MONGO_ERROR = -1
)

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

// M is a shortcut for writing map[string]interface{} in BSON literal
// expressions. The type M is encoded the same as the type
// map[string]interface{}.
type M map[string]interface{}

type Mongo struct {
    conn *C.mongo
}

type Cursor struct {
    Conn   *Mongo
    cursor *C.mongo_cursor
}

type MongoWriteConcern struct {
    writeConcern *C.mongo_write_concern
}

func NewMongo() *Mongo {
    m := &Mongo{}
    m.conn = &C.mongo{}
    return m
}

func (m *Mongo) Db(dbname string) *DB {
    db := &DB{Name: dbname, Conn: m}
    return db
}

/*********************************************************************
CRUD API
**********************************************************************/

/**
 * Insert a BSON document into a MongoDB server. This function
 * will fail if the supplied BSON struct is not UTF-8 or if
 * the keys are invalid for insert (contain '.' or start with '$').
 *
 * The default write concern set on the conn object will be used.
 *
 * @param conn a mongo object.
 * @param ns the namespace.
 * @param data the bson data.
 * @param custom_write_concern a write concern object that will
 *     override any write concern set on the conn object.
 *
 * @return MONGO_OK or MONGO_ERROR. If the conn->err
 *     field is MONGO_BSON_INVALID, check the err field
 *     on the bson struct for the reason.
 */
// MONGO_EXPORT int mongo_insert( mongo *conn, const char *ns, const bson *data,
//                                mongo_write_concern *custom_write_concern );
func (m *Mongo) Insert(ns string, data *Bson, writeConcern *MongoWriteConcern) int {
    if writeConcern == nil {
        return int(C.mongo_insert(m.conn, C.CString(ns), data._bson, nil))
    }
    return int(C.mongo_insert(m.conn, C.CString(ns), data._bson, writeConcern.writeConcern))
}

/**
 * Insert a batch of BSON documents into a MongoDB server. This function
 * will fail if any of the documents to be inserted is invalid.
 *
 * The default write concern set on the conn object will be used.
 *
 * @param conn a mongo object.
 * @param ns the namespace.
 * @param data the bson data.
 * @param num the number of documents in data.
 * @param custom_write_concern a write concern object that will
 *     override any write concern set on the conn object.
 * @param flags flags on this batch insert. Currently, this value
 *     may be 0 or MONGO_CONTINUE_ON_ERROR, which will cause the
 *     batch insert to continue even if a given insert in the batch fails.
 *
 * @return MONGO_OK or MONGO_ERROR.
 */
// MONGO_EXPORT int mongo_insert_batch( mongo *conn, const char *ns,
//                                      const bson **data, int num, mongo_write_concern *custom_write_concern,
//                                      int flags );

/**
 * Update a document in a MongoDB server.
 *
 * The default write concern set on the conn object will be used.
 *
 * @param conn a mongo object.
 * @param ns the namespace.
 * @param cond the bson update query.
 * @param op the bson update data.
 * @param flags flags for the update.
 * @param custom_write_concern a write concern object that will
 *     override any write concern set on the conn object.
 *
 * @return MONGO_OK or MONGO_ERROR with error stored in conn object.
 *
 */
// MONGO_EXPORT int mongo_update( mongo *conn, const char *ns, const bson *cond,
//                                const bson *op, int flags, mongo_write_concern *custom_write_concern );
func (m *Mongo) Update(ns string, cond, op *Bson, flags int, writeConcern *MongoWriteConcern) int {
    if writeConcern == nil {
        return int(C.mongo_update(m.conn, C.CString(ns), cond._bson, op._bson, C.int(flags), nil))
    }
    return int(C.mongo_update(m.conn, C.CString(ns), cond._bson,
        op._bson, C.int(flags), writeConcern.writeConcern))
}

/**
 * Remove a document from a MongoDB server.
 *
 * The default write concern set on the conn object will be used.
 *
 * @param conn a mongo object.
 * @param ns the namespace.
 * @param cond the bson query.
 * @param custom_write_concern a write concern object that will
 *     override any write concern set on the conn object.
 *
 * @return MONGO_OK or MONGO_ERROR with error stored in conn object.
 */
// MONGO_EXPORT int mongo_remove( mongo *conn, const char *ns, const bson *cond,
//                                mongo_write_concern *custom_write_concern );
func (m *Mongo) Remove(ns string, cond *Bson, writeConcern *MongoWriteConcern) int {
    if writeConcern == nil {
        return int(C.mongo_remove(m.conn, C.CString(ns), cond._bson, nil))
    }
    return int(C.mongo_remove(m.conn, C.CString(ns), cond._bson, writeConcern.writeConcern))
}

/*********************************************************************
Write Concern API
**********************************************************************/

// /**
//  * Initialize a mongo_write_concern object. Effectively zeroes out the struct.
//  *
//  */
// MONGO_EXPORT void mongo_write_concern_init( mongo_write_concern *write_concern );

// /**
//  * Finish this write concern object by serializing the literal getlasterror
//  * command that will be sent to the server.
//  *
//  * You must call mongo_write_concern_destroy() to free the serialized BSON.
//  *
//  */
// MONGO_EXPORT int mongo_write_concern_finish( mongo_write_concern *write_concern );

// /**
//  * Free the write_concern object (specifically, the BSON that it owns).
//  *
//  */
// MONGO_EXPORT void mongo_write_concern_destroy( mongo_write_concern *write_concern );

/*********************************************************************
Cursor API
**********************************************************************/

/**
 * Find documents in a MongoDB server.
 *
 * @param conn a mongo object.
 * @param ns the namespace.
 * @param query the bson query.
 * @param fields a bson document of fields to be returned.
 * @param limit the maximum number of documents to return.
 * @param skip the number of documents to skip.
 * @param options A bitfield containing cursor options.
 *
 * @return A cursor object allocated on the heap or NULL if
 *     an error has occurred. For finer-grained error checking,
 *     use the cursor builder API instead.
 */
// MONGO_EXPORT mongo_cursor *mongo_find( mongo *conn, const char *ns, const bson *query,
//                                        const bson *fields, int limit, int skip, int options );
func (m *Mongo) Find(ns string, query *Bson, fields *Bson, limit, skip, options int) (*Cursor, error) {
    if query == nil {
        query = &Bson{}
    }
    if fields == nil {
        fields = &Bson{}
    }
    c := C.mongo_find(m.conn, C.CString(ns), query._bson,
        fields._bson, C.int(limit), C.int(skip), C.int(options))
    if c == nil {
        // error need check
        return nil, errors.New("has error: " + m.Error().Error())
    }
    c2 := &Cursor{
        cursor: c,
    }
    return c2, nil
}

/**
 * Initalize a new cursor object.
 *
 * @param cursor
 * @param ns the namespace, represented as the the database
 *     name and collection name separated by a dot. e.g., "test.users"
 */
// MONGO_EXPORT void mongo_cursor_init( mongo_cursor *cursor, mongo *conn, const char *ns );
func (cur *Cursor) Init(conn *Mongo, ns string) {
    cur.Conn = conn
    C.mongo_cursor_init(cur.cursor, conn.conn, C.CString(ns))
}

/**
 * Set the bson object specifying this cursor's query spec. If
 * your query is the empty bson object "{}", then you need not
 * set this value.
 *
 * @param cursor
 * @param query a bson object representing the query spec. This may
 *   be either a simple query spec or a complex spec storing values for
 *   $query, $orderby, $hint, and/or $explain. See
 *   http://www.mongodb.org/display/DOCS/Mongo+Wire+Protocol for details.
 */
// MONGO_EXPORT void mongo_cursor_set_query( mongo_cursor *cursor, const bson *query );
func (cur *Cursor) SetQuery(query *Bson) {
    C.mongo_cursor_set_query(cur.cursor, query._bson)
}

// /**
//  * Set the fields to return for this cursor. If you want to return
//  * all fields, you need not set this value.
//  *
//  * @param cursor
//  * @param fields a bson object representing the fields to return.
//  *   See http://www.mongodb.org/display/DOCS/Retrieving+a+Subset+of+Fields.
//  */
// MONGO_EXPORT void mongo_cursor_set_fields( mongo_cursor *cursor, const bson *fields );
func (cur *Cursor) SetFields(fields *Bson) {
    C.mongo_cursor_set_fields(cur.cursor, fields._bson)
}

// /**
//  * Set the number of documents to skip.
//  *
//  * @param cursor
//  * @param skip
//  */
// MONGO_EXPORT void mongo_cursor_set_skip( mongo_cursor *cursor, int skip );
func (cur *Cursor) SetSkip(skip int) {
    C.mongo_cursor_set_skip(cur.cursor, C.int(skip))
}

// /**
//  * Set the number of documents to return.
//  *
//  * @param cursor
//  * @param limit
//  */
// MONGO_EXPORT void mongo_cursor_set_limit( mongo_cursor *cursor, int limit );
func (cur *Cursor) SetLimit(limit int) {
    C.mongo_cursor_set_limit(cur.cursor, C.int(limit))
}

// /**
//  * Set any of the available query options (e.g., MONGO_TAILABLE).
//  *
//  * @param cursor
//  * @param options a bitfield storing query options. See
//  *   mongo_cursor_bitfield_t for available constants.
//  */
// MONGO_EXPORT void mongo_cursor_set_options( mongo_cursor *cursor, int options );
func (cur *Cursor) SetOptions(options int) {
    C.mongo_cursor_set_options(cur.cursor, C.int(options))
}

// /**
//  * Return the current BSON object data as a const char*. This is useful
//  * for creating bson iterators with bson_iterator_init.
//  *
//  * @param cursor
//  */
// MONGO_EXPORT const char *mongo_cursor_data( mongo_cursor *cursor );

// /**
//  * Return the current BSON object data as a const char*. This is useful
//  * for creating bson iterators with bson_iterator_init.
//  *
//  * @param cursor
//  */
// MONGO_EXPORT const bson *mongo_cursor_bson( mongo_cursor *cursor );
func (cur *Cursor) Bson() *Bson {
    b := C.mongo_cursor_bson(cur.cursor)
    return &Bson{_bson: b}
}

// This cursor's current bson object. 
func (cur *Cursor) Current() *Bson {
    return &Bson{_bson: &cur.cursor.current}
}

// /**
//  * Iterate the cursor, returning the next item. When successful,
//  *   the returned object will be stored in cursor->current;
//  *
//  * @param cursor
//  *
//  * @return MONGO_OK. On error, returns MONGO_ERROR and sets
//  *   cursor->err with a value of mongo_error_t.
//  */
// MONGO_EXPORT int mongo_cursor_next( mongo_cursor *cursor );
func (cur *Cursor) Next() int {
    return int(C.mongo_cursor_next(cur.cursor))
}

// /**
//  * Destroy a cursor object. When finished with a cursor, you
//  * must pass it to this function.
//  *
//  * @param cursor the cursor to destroy.
//  *
//  * @return MONGO_OK or an error code. On error, check cursor->conn->err
//  *     for errors.
//  */
// MONGO_EXPORT int mongo_cursor_destroy( mongo_cursor *cursor );
func (cur *Cursor) Destroy() int {
    return int(C.mongo_cursor_destroy(cur.cursor))
}

func (c *Cursor) GetIterator() *BsonIterator {
    it := NewBsonIterator()
    it.Init(c.Current())
    return it
}

// /**
//  * Find a single document in a MongoDB server.
//  *
//  * @param conn a mongo object.
//  * @param ns the namespace.
//  * @param query the bson query.
//  * @param fields a bson document of the fields to be returned.
//  * @param out a bson document in which to put the query result.
//  *
//  */
// /* out can be NULL if you don't care about results. useful for commands */
// MONGO_EXPORT int mongo_find_one( mongo *conn, const char *ns, const bson *query,
//                                  const bson *fields, bson *out );
func (m *Mongo) FindOne(ns string, query, fields, out *Bson) int {
    return int(C.mongo_find_one(m.conn, C.CString(ns), query._bson, fields._bson, out._bson))
}

// /*********************************************************************
// Command API and Helpers
// **********************************************************************/

// /**
//  * Count the number of documents in a collection matching a query.
//  *
//  * @param conn a mongo object.
//  * @param db the db name.
//  * @param coll the collection name.
//  * @param query the BSON query.
//  *
//  * @return the number of matching documents. If the command fails,
//  *     MONGO_ERROR is returned.
//  */
// MONGO_EXPORT double mongo_count( mongo *conn, const char *db, const char *coll,
//                                  const bson *query );
func (m *Mongo) Count(db, coll string, query *Bson) int64 {
    if query == nil {
        query = &Bson{}
    }
    return int64(C.mongo_count(m.conn, C.CString(db), C.CString(coll), query._bson))
}

// /**
//  * Create a compound index.
//  *
//  * @param conn a mongo object.
//  * @param ns the namespace.
//  * @param key the bson index key.
//  * @param name the optional name, use NULL to generate a default name.
//  * @param options a bitfield for setting index options. Possibilities include
//  *   MONGO_INDEX_UNIQUE, MONGO_INDEX_DROP_DUPS, MONGO_INDEX_BACKGROUND,
//  *   and MONGO_INDEX_SPARSE.
//  * @param out a bson document containing errors, if any.
//  *
//  * @return MONGO_OK if index is created successfully; otherwise, MONGO_ERROR.
//  */
// MONGO_EXPORT int mongo_create_index( mongo *conn, const char *ns, const bson *key,
//                                      const char *name, int options, bson *out );

// *
//  * Create a capped collection.
//  *
//  * @param conn a mongo object.
//  * @param ns the namespace (e.g., "dbname.collectioname")
//  * @param size the size of the capped collection in bytes.
//  * @param max the max number of documents this collection is
//  *   allowed to contain. If zero, this argument will be ignored
//  *   and the server will use the collection's size to age document out.
//  *   If using this option, ensure that the total size can contain this
//  *   number of documents.

// MONGO_EXPORT int mongo_create_capped_collection( mongo *conn, const char *db,
//         const char *collection, int size, int max, bson *out );

// /**
//  * Create an index with a single key.
//  *
//  * @param conn a mongo object.
//  * @param ns the namespace.
//  * @param field the index key.
//  * @param options index options.
//  * @param out a BSON document containing errors, if any.
//  *
//  * @return true if the index was created.
//  */
// MONGO_EXPORT bson_bool_t mongo_create_simple_index( mongo *conn, const char *ns,
//         const char *field, int options, bson *out );

// /**
//  * Run a command on a MongoDB server.
//  *
//  * @param conn a mongo object.
//  * @param db the name of the database.
//  * @param command the BSON command to run.
//  * @param out the BSON result of the command.
//  *
//  * @return MONGO_OK if the command ran without error.
//  */
// MONGO_EXPORT int mongo_run_command( mongo *conn, const char *db,
//                                     const bson *command, bson *out );

// /**
//  * Run a command that accepts a simple string key and integer value.
//  *
//  * @param conn a mongo object.
//  * @param db the name of the database.
//  * @param cmd the command to run.
//  * @param arg the integer argument to the command.
//  * @param out the BSON result of the command.
//  *
//  * @return MONGO_OK or an error code.
//  *
//  */
// MONGO_EXPORT int mongo_simple_int_command( mongo *conn, const char *db,
//         const char *cmd, int arg, bson *out );

// /**
//  * Run a command that accepts a simple string key and value.
//  *
//  * @param conn a mongo object.
//  * @param db the name of the database.
//  * @param cmd the command to run.
//  * @param arg the string argument to the command.
//  * @param out the BSON result of the command.
//  *
//  * @return true if the command ran without error.
//  *
//  */
// MONGO_EXPORT int mongo_simple_str_command( mongo *conn, const char *db,
//         const char *cmd, const char *arg, bson *out );

// /**
//  * Drop a database.
//  *
//  * @param conn a mongo object.
//  * @param db the name of the database to drop.
//  *
//  * @return MONGO_OK or an error code.
//  */
// MONGO_EXPORT int mongo_cmd_drop_db( mongo *conn, const char *db );

// /**
//  * Drop a collection.
//  *
//  * @param conn a mongo object.
//  * @param db the name of the database.
//  * @param collection the name of the collection to drop.
//  * @param out a BSON document containing the result of the command.
//  *
//  * @return true if the collection drop was successful.
//  */
// MONGO_EXPORT int mongo_cmd_drop_collection( mongo *conn, const char *db,
//         const char *collection, bson *out );

// /**
//  * Add a database user.
//  *
//  * @param conn a mongo object.
//  * @param db the database in which to add the user.
//  * @param user the user name
//  * @param pass the user password
//  *
//  * @return MONGO_OK or MONGO_ERROR.
//   */
// MONGO_EXPORT int mongo_cmd_add_user( mongo *conn, const char *db,
//                                      const char *user, const char *pass );

// /**
//  * Authenticate a user.
//  *
//  * @param conn a mongo object.
//  * @param db the database to authenticate against.
//  * @param user the user name to authenticate.
//  * @param pass the user's password.
//  *
//  * @return MONGO_OK on sucess and MONGO_ERROR on failure.
//  */
// MONGO_EXPORT int mongo_cmd_authenticate( mongo *conn, const char *db,
//         const char *user, const char *pass );

// /**
//  * Check if the current server is a master.
//  *
//  * @param conn a mongo object.
//  * @param out a BSON result of the command.
//  *
//  * @return true if the server is a master.
//  */
// /* return value is master status */
// MONGO_EXPORT bson_bool_t mongo_cmd_ismaster( mongo *conn, bson *out );

// /**
//  * Get the error for the last command with the current connection.
//  *
//  * @param conn a mongo object.
//  * @param db the name of the database.
//  * @param out a BSON object containing the error details.
//  *
//  * @return MONGO_OK if no error and MONGO_ERROR on error. On error, check the values
//  *     of conn->lasterrcode and conn->lasterrstr for the error status.
//  */
// MONGO_EXPORT int mongo_cmd_get_last_error( mongo *conn, const char *db, bson *out );

// /**
//  * Get the most recent error with the current connection.
//  *
//  * @param conn a mongo object.
//  * @param db the name of the database.
//  * @param out a BSON object containing the error details.
//  *
//  * @return MONGO_OK if no error and MONGO_ERROR on error. On error, check the values
//  *     of conn->lasterrcode and conn->lasterrstr for the error status.
//  */
// MONGO_EXPORT int mongo_cmd_get_prev_error( mongo *conn, const char *db, bson *out );

// /**
//  * Reset the error state for the connection.
//  *
//  * @param conn a mongo object.
//  * @param db the name of the database.
//  */
// MONGO_EXPORT void mongo_cmd_reset_error( mongo *conn, const char *db );

// /*********************************************************************
// Utility API
// **********************************************************************/

// MONGO_EXPORT mongo* mongo_alloc( void );
// MONGO_EXPORT void mongo_dealloc(mongo* conn);
// MONGO_EXPORT int mongo_get_err(mongo* conn);
// MONGO_EXPORT int mongo_is_connected(mongo* conn);
// MONGO_EXPORT int mongo_get_op_timeout(mongo* conn);
// MONGO_EXPORT const char* mongo_get_primary(mongo* conn);
// MONGO_EXPORT SOCKET mongo_get_socket(mongo* conn) ;
// MONGO_EXPORT int mongo_get_host_count(mongo* conn);
// MONGO_EXPORT const char* mongo_get_host(mongo* conn, int i);
// MONGO_EXPORT mongo_write_concern* mongo_write_concern_alloc( void );
// MONGO_EXPORT void mongo_write_concern_dealloc(mongo_write_concern* write_concern);
// MONGO_EXPORT mongo_cursor* mongo_cursor_alloc( void );
// MONGO_EXPORT void mongo_cursor_dealloc(mongo_cursor* cursor);
// MONGO_EXPORT int  mongo_get_server_err(mongo* conn);
// MONGO_EXPORT const char*  mongo_get_server_err_string(mongo* conn);

// /**
//  * Set an error on a mongo connection object. Mostly for internal use.
//  *
//  * @param conn a mongo connection object.
//  * @param err a driver error code of mongo_error_t.
//  * @param errstr a string version of the error.
//  * @param errorcode Currently errno or WSAGetLastError().
//  */
// MONGO_EXPORT void __mongo_set_error( mongo *conn, mongo_error_t err,
//                                      const char *errstr, int errorcode );
// /**
//  * Clear all errors stored on a mongo connection object.
//  *
//  * @param conn a mongo connection object.
//  */
// MONGO_EXPORT void mongo_clear_errors( mongo *conn );
