package libgomongo

// #cgo CFLAGS: -std=gnu99 -I./mongo-c-driver/src/
// #cgo LDFLAGS: -L./mongo-c-driver/src/ -lmongoc
// #include "mongo.h"
import "C"

import (
    "errors"
    "fmt"
    // "tim
)

/*********************************************************************
Connection API
**********************************************************************/

func (c *Mongo) ErrNo() MongoError {
    return MongoError(c.conn.err)
}

func (c *Mongo) Error() error {
    var err error
    status := c.ErrNo()
    switch status {
    case MONGO_CONN_SUCCESS:
    case MONGO_CONN_NO_SOCKET:
        err = errors.New("MongoDB: Could not create a socket.")
    case MONGO_CONN_FAIL:
        err = errors.New("MongoDB: An error occured while calling connect(). ")
    case MONGO_CONN_ADDR_FAIL:
        err = errors.New("MongoDB: An error occured while calling getaddrinfo().")
    case MONGO_CONN_NOT_MASTER:
        err = errors.New("MongoDB [Warning]: connected to a non-master node (read-only).")
    case MONGO_CONN_BAD_SET_NAME:
        err = errors.New("MongoDB: Given rs name doesn't match this replica set.")
    case MONGO_CONN_NO_PRIMARY:
        err = errors.New("MongoDB: Can't find primary in replica set. Connection closed.")

    case MONGO_IO_ERROR:
        err = errors.New("MongoDB: An error occurred while reading or writing on the socket.")
    case MONGO_SOCKET_ERROR:
        err = errors.New("MongoDB: Other socket error.")
    case MONGO_READ_SIZE_ERROR:
        err = errors.New("MongoDB: The response is not the expected length.")
    case MONGO_COMMAND_FAILED:
        err = errors.New("MongoDB: The command returned with 'ok' value of 0.")
    case MONGO_WRITE_ERROR:
        err = errors.New("MongoDB: Write with given write_concern returned an error.")
    case MONGO_NS_INVALID:
        err = errors.New("MongoDB: The name for the ns (database or collection) is invalid.")
    case MONGO_BSON_INVALID:
        err = errors.New("MongoDB: BSON not valid for the specified op.")
    case MONGO_BSON_NOT_FINISHED:
        err = errors.New("MongoDB: BSON object has not been finished.")
    case MONGO_BSON_TOO_LARGE:
        err = errors.New("MongoDB: BSON object exceeds max BSON size.")
    case MONGO_WRITE_CONCERN_INVALID:
        err = errors.New("MongoDB: Supplied write concern object is invalid.")
    default:
        err = errors.New(fmt.Sprintf("MongoDB: Unkonw error[%d]", status))
    }
    return err
}

// /** Initialize sockets for Windows.
//  */
// MONGO_EXPORT void mongo_init_sockets( void );

/**
 * Initialize a new mongo connection object. You must initialize each mongo
 * object using this function.
 *
 *  @note When finished, you must pass this object to
 *      mongo_destroy( ).
 *
 *  @param conn a mongo connection object allocated on the stack
 *      or heap.
 */
func (c *Mongo) Init() {
    C.mongo_init(c.conn)
}

// /**
//  * Connect to a single MongoDB server.
//  *
//  * @param conn a mongo object.
//  * @param host a numerical network address or a network hostname.
//  * @param port the port to connect to.
//  *
//  * @return MONGO_OK or MONGO_ERROR on failure. On failure, a constant of type
//  *   mongo_error_t will be set on the conn->err field.
//  */
// MONGO_EXPORT int mongo_client( mongo *conn , const char *host, int port );
func (c *Mongo) Client(host string, port int) int {
    return int(C.mongo_client(c.conn, C.CString(host), C.int(port)))
}

// /**
//  * DEPRECATED - use mongo_client.
//  * Connect to a single MongoDB server.
//  *
//  * @param conn a mongo object.
//  * @param host a numerical network address or a network hostname.
//  * @param port the port to connect to.
//  *
//  * @return MONGO_OK or MONGO_ERROR on failure. On failure, a constant of type
//  *   mongo_error_t will be set on the conn->err field.
//  */
// MONGO_EXPORT int mongo_connect( mongo *conn , const char *host, int port );
func (c *Mongo) Connect(host string, port int) int {
    return c.Client(host, port)
}

// /**
//  * Set up this connection object for connecting to a replica set.
//  * To connect, pass the object to mongo_replica_set_client().
//  *
//  * @param conn a mongo object.
//  * @param name the name of the replica set to connect to.
//  * */
// MONGO_EXPORT void mongo_replica_set_init( mongo *conn, const char *name );
func (c *Mongo) ReplicaSetInit(name string) {
    C.mongo_replica_set_init(c.conn, C.CString(name))
}

// /**
//  * DEPRECATED - use mongo_replica_set_init.
//  * Set up this connection object for connecting to a replica set.
//  * To connect, pass the object to mongo_replset_connect().
//  *
//  * @param conn a mongo object.
//  * @param name the name of the replica set to connect to.
//  * */
// MONGO_EXPORT void mongo_replset_init( mongo *conn, const char *name );

// /**
//  * Add a seed node to the replica set connection object.
//  *
//  * You must specify at least one seed node before connecting to a replica set.
//  *
//  * @param conn a mongo object.
//  * @param host a numerical network address or a network hostname.
//  * @param port the port to connect to.
//  */
// MONGO_EXPORT void mongo_replica_set_add_seed( mongo *conn, const char *host, int port );

// /**
//  * DEPRECATED - use mongo_replica_set_add_seed.
//  * Add a seed node to the replica set connection object.
//  *
//  * You must specify at least one seed node before connecting to a replica set.
//  *
//  * @param conn a mongo object.
//  * @param host a numerical network address or a network hostname.
//  * @param port the port to connect to.
//  */
// MONGO_EXPORT void mongo_replset_add_seed( mongo *conn, const char *host, int port );

// /**
//  * Utility function for converting a host-port string to a mongo_host_port.
//  *
//  * @param host_string a string containing either a host or a host and port separated
//  *     by a colon.
//  * @param host_port the mongo_host_port object to write the result to.
//  */
// void mongo_parse_host( const char *host_string, mongo_host_port *host_port );

// /**
//  * Utility function for validation database and collection names.
//  *
//  * @param conn a mongo object.
//  *
//  * @return MONGO_OK or MONGO_ERROR on failure. On failure, a constant of type
//  *   mongo_conn_return_t will be set on the conn->err field.
//  *
//  */
// MONGO_EXPORT int mongo_validate_ns( mongo *conn, const char *ns );

// *
//  * Connect to a replica set.
//  *
//  * Before passing a connection object to this function, you must already have called
//  * mongo_set_replica_set and mongo_replica_set_add_seed.
//  *
//  * @param conn a mongo object.
//  *
//  * @return MONGO_OK or MONGO_ERROR on failure. On failure, a constant of type
//  *   mongo_conn_return_t will be set on the conn->err field.

// MONGO_EXPORT int mongo_replica_set_client( mongo *conn );

// /**
//  * DEPRECATED - use mongo_replica_set_client.
//  * Connect to a replica set.
//  *
//  * Before passing a connection object to this function, you must already have called
//  * mongo_set_replset and mongo_replset_add_seed.
//  *
//  * @param conn a mongo object.
//  *
//  * @return MONGO_OK or MONGO_ERROR on failure. On failure, a constant of type
//  *   mongo_conn_return_t will be set on the conn->err field.
//  */
// MONGO_EXPORT int mongo_replset_connect( mongo *conn );

// /** Set a timeout for operations on this connection. This
//  *  is a platform-specific feature, and only work on *nix
//  *  system. You must also compile for linux to support this.
//  *
//  *  @param conn a mongo object.
//  *  @param millis timeout time in milliseconds.
//  *
//  *  @return MONGO_OK. On error, return MONGO_ERROR and
//  *    set the conn->err field.
//  */
// MONGO_EXPORT int mongo_set_op_timeout( mongo *conn, int millis );
func (c *Mongo) SetOpTimeout(millis int) int {
    return int(C.mongo_set_op_timeout(c.conn, C.int(millis)))
}

// /**
//  * Ensure that this connection is healthy by performing
//  * a round-trip to the server.
//  *
//  * @param conn a mongo connection
//  *
//  * @return MONGO_OK if connected; otherwise, MONGO_ERROR.
//  */
// MONGO_EXPORT int mongo_check_connection( mongo *conn );
func (c *Mongo) CheckConnection() int {
    return int(C.mongo_check_connection(c.conn))
}

// /**
//  * Try reconnecting to the server using the existing connection settings.
//  *
//  * This function will disconnect the current socket. If you've authenticated,
//  * you'll need to re-authenticate after calling this function.
//  *
//  * @param conn a mongo object.
//  *
//  * @return MONGO_OK or MONGO_ERROR and
//  *   set the conn->err field.
//  */
// MONGO_EXPORT int mongo_reconnect( mongo *conn );
func (c *Mongo) Reconnect() int {
    return int(C.mongo_reconnect(c.conn))
}

// /**
//  * Close the current connection to the server. After calling
//  * this function, you may call mongo_reconnect with the same
//  * connection object.
//  *
//  * @param conn a mongo object.
//  */
// MONGO_EXPORT void mongo_disconnect( mongo *conn );
func (c *Mongo) Disconnect() {
    C.mongo_disconnect(c.conn)
}

// /**
//  * Close any existing connection to the server and free all allocated
//  * memory associated with the conn object.
//  *
//  * You must always call this function when finished with the connection object.
//  *
//  * @param conn a mongo object.
//  */
// MONGO_EXPORT void mongo_destroy( mongo *conn );
func (c *Mongo) Destroy() {
    C.mongo_destroy(c.conn)
}

// /**
//  * Specify the write concern object that this connection should use
//  * by default for all writes (inserts, updates, and deletes). This value
//  * can be overridden by passing a write_concern object to any write function.
//  *
//  * @param conn a mongo object.
//  * @param write_concern pointer to a write concern object.
//  *
//  */
// MONGO_EXPORT void mongo_set_write_concern( mongo *conn,
//         mongo_write_concern *write_concern );
func (c *Mongo) SetWriteConcern(mongo_write_concern *MongoWriteConcern) {
    C.mongo_set_write_concern(c.conn, mongo_write_concern.writeConcern)
}

// /**
//  * The following functions get the attributes of the write_concern object.
//  *
//  */
// MONGO_EXPORT int mongo_write_concern_get_w( mongo_write_concern *write_concern );
// MONGO_EXPORT int mongo_write_concern_get_wtimeout( mongo_write_concern *write_concern );
// MONGO_EXPORT int mongo_write_concern_get_j( mongo_write_concern *write_concern );
// MONGO_EXPORT int mongo_write_concern_get_fsync( mongo_write_concern *write_concern );
// MONGO_EXPORT const char* mongo_write_concern_get_mode( mongo_write_concern *write_concern );
// MONGO_EXPORT bson* mongo_write_concern_get_cmd( mongo_write_concern *write_concern );

// /**
//  * The following functions set the attributes of the write_concern object.
//  *
//  */
// MONGO_EXPORT void mongo_write_concern_set_w( mongo_write_concern *write_concern, int w );
// MONGO_EXPORT void mongo_write_concern_set_wtimeout( mongo_write_concern *write_concern, int wtimeout );
// MONGO_EXPORT void mongo_write_concern_set_j( mongo_write_concern *write_concern, int j );
// MONGO_EXPORT void mongo_write_concern_set_fsync( mongo_write_concern *write_concern, int fsync );
// MONGO_EXPORT void mongo_write_concern_set_mode( mongo_write_concern *write_concern, const char* mode );
