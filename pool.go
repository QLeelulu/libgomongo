package libgomongo

// Pool maintains a pool of database connections.
//
// The following example shows how to use a pool in a web application. The
// application creates a pool at application startup and makes it available to
// request handlers, possibly using a global variable:
//
//      var host string           // host of server
//      var port int              // port of server
//      var poolsize int          // max pool size
//
//      ...
//
//      pool = mongo.NewPool(host, port, 3)
//
// This pool has a maximum of three connections to the server specified by the
// variable "server". Each connection is logged into the "admin" database using
// the credentials specified by the variables "name" and "password".
//
// A request handler gets a connection from the pool and closes the connection
// when the handler is done:
//
//  conn, err := pool.Get()
//  if err != nil {
//      // handle the error
//  }
//  defer conn.Close()
//  // do something with the connection
//
// Close() returns the connection to the pool if there's room in the pool and
// the connection does not have a permanent error. Otherwise, Close() releases
// the resources used by the connection.
type Pool struct {
    Size int
    Host string
    Port int

    conns chan *Mongo
}

// NewPool returns a new connection pool. The pool create
// connections as needed and maintains a maximum of maxIdle idle connections.
func NewPool(host string, port int, maxIdle int) *Pool {
    return &Pool{Host: host, Port: port, Size: maxIdle,
        conns: make(chan *Mongo, maxIdle)}
}

// Get returns an idle connection from the pool if available or creates a new
// connection. The caller should Close() the connection to return the
// connection to the pool.
func (p *Pool) Get() (*Mongo, error) {
    var conn *Mongo
    select {
    case conn = <-p.conns:
    default:
        var err error
        conn = NewMongo()
        status := conn.Client(p.Host, p.Port)
        if status != MONGO_OK {
            err = conn.Error()
        }
        if err != nil {
            return nil, err
        }
    }
    conn.pool = p
    return conn, nil
}

func (p *Pool) Put(conn *Mongo) {
    select {
    case p.conns <- conn:
    default:
        conn.Destroy()
    }
}
