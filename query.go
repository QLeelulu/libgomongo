package libgomongo

import (
// "fmt"
)

// FindOptions specifies options for the Conn.Find method.
type FindOptions struct {
    // Optional document that limits the fields in the returned documents.
    // Fields contains one or more elements, each of which is the name of a
    // field that should be returned, and the integer value 1.
    Fields M

    // Do not close the cursor when no more data is available on the server.
    Tailable bool

    // Allow query of replica slave.
    SlaveOk bool

    // Do not close the cursor on the server after a period of inactivity (10
    // minutes).
    NoCursorTimeout bool

    // Block at server for a short time if there's no data for a tailable cursor.
    AwaitData bool

    // Stream the data down from the server full blast. Normally the server
    // waits for a "get more" message before sending a batch of data to the
    // client. With this option set, the server sends batches of data without
    // waiting for the "get more" messages.
    Exhaust bool

    // Allow partial results in sharded environment. Normally the query
    // will fail if a shard is not available.
    PartialResults bool

    // Skip specifies the number of documents the server should skip at the
    // beginning of the result set.
    Skip int

    // Sets the number of documents to return.
    Limit int

    // Sets the batch size used for sending documents from the server to the
    // client.
    BatchSize int
}

// QuerySpec is a helper for specifying complex queries.
type QuerySpec struct {
    // The filter. This field is required.
    Query M   `bson:"$query"`

    // Sort order specified by (key, direction) pairs. The direction is 1 for
    // ascending order and -1 for descending order.
    Sort M   `bson:"$orderby"`

    // If set to true, then the query returns an explain plan record the query.
    // See http://www.mongodb.org/display/DOCS/Optimization#Optimization-Explain
    Explain bool `bson:"$explain,omitempty"`

    // Index hint specified by (key, direction) pairs.
    // See http://www.mongodb.org/display/DOCS/Optimization#Optimization-Hint
    Hint M   `bson:"$hint"`

    // Snapshot mode assures that objects which update during the lifetime of a
    // query are returned once and only once.
    // See http://www.mongodb.org/display/DOCS/How+to+do+Snapshotted+Queries+in+the+Mongo+Database
    Snapshot bool `bson:"$snapshot,omitempty"`

    // Min and Max constrain matches to those having index keys between the min
    // and max keys specified.The Min value is included in the range and the
    // Max value is excluded.
    // See http://www.mongodb.org/display/DOCS/min+and+max+Query+Specifiers
    Min interface{} `bson:"$min"`
    Max interface{} `bson:"$max"`
}

// Query represents a query to the database.
type Query struct {
    Conn      *Mongo // connection
    Namespace string
    Spec      QuerySpec
    Options   FindOptions
}

func NewQuery(conn *Mongo, namespace string) Query {
    q := Query{}
    q.Conn = conn
    q.Namespace = namespace
    q.Spec = QuerySpec{}
    q.Options = FindOptions{}
    return q
}

// Sort specifies the sort order for the result. The order is specified by
// (key, direction) pairs. Direction is 1 for ascending order and -1 for
// descending order.
func (q *Query) Sort(sort M) *Query {
    q.Spec.Sort = sort
    return q
}

// Hint specifies an index hint. The index is specified by (key, direction)
// pairs. Direction is 1 for ascending order and -1 for descending order.
//
// More information: http://www.mongodb.org/display/DOCS/Optimization#Optimization-Hint
func (q *Query) Hint(hint M) *Query {
    q.Spec.Hint = hint
    return q
}

// Limit specifies the number of documents to return from the query.
//
// More information: http://www.mongodb.org/display/DOCS/Advanced+Queries#AdvancedQueries-%7B%7Blimit%28%29%7D%7D
func (q *Query) Limit(limit int) *Query {
    q.Options.Limit = limit
    return q
}

// Skip specifies the number of documents the server should skip at the
// beginning of the result set.
//
// More information: http://www.mongodb.org/display/DOCS/Advanced+Queries#AdvancedQueries-%7B%7Bskip%28%29%7D%7D
func (q *Query) Skip(skip int) *Query {
    q.Options.Skip = skip
    return q
}

// BatchSize sets the batch sized used for sending documents from the server to
// the client.
func (q *Query) BatchSize(batchSize int) *Query {
    q.Options.BatchSize = batchSize
    return q
}

// Fields limits the fields in the returned documents. Fields contains one or
// more elements, each of which is the name of a field that should be returned,
// and the integer value 1.
//
// More information: http://www.mongodb.org/display/DOCS/Retrieving+a+Subset+of+Fields
func (q *Query) Fields(fields M) *Query {
    q.Options.Fields = fields
    return q
}

// SlaveOk specifies if query can be routed to a slave.
//
// More information: http://www.mongodb.org/display/DOCS/Querying#Querying-slaveOk
func (q *Query) SlaveOk(slaveOk bool) *Query {
    q.Options.SlaveOk = slaveOk
    return q
}

// PartialResults specifies if mongos can reply with partial results when a
// shard is missing.
func (q *Query) PartialResults(ok bool) *Query {
    q.Options.PartialResults = ok
    return q
}

// Exhaust specifies if the server should stream data to the client full blast.
// Normally the server waits for a "get more" message before sending a batch of
// data to the client.  With this option set, the server sends batches of data
// without waiting for the "get more" messages.
func (q *Query) Exhaust(exhaust bool) *Query {
    q.Options.Exhaust = exhaust
    return q
}

// Tailable specifies if the server should not close the cursor when no more
// data is available.
//
// More information: http://www.mongodb.org/display/DOCS/Tailable+Cursors
func (q *Query) Tailable(tailable bool) *Query {
    q.Options.Tailable = tailable
    return q
}

func (q *Query) bsonQuery() (*Bson, error) {
    if q.Spec.Query == nil {
        return nil, nil
    }
    b := NewBsonFromM(q.Spec.Query)
    return b, nil
}

func (q *Query) bsonFields() (*Bson, error) {
    if q.Options.Fields == nil {
        return nil, nil
    }
    b := NewBsonFromM(q.Options.Fields)
    return b, nil
}

// Cursor executes the query and returns a cursor over the results. Subsequent
// changes to the query object are ignored by the cursor.
func (q *Query) Cursor() (*Cursor, error) {
    query, err := q.bsonQuery()
    if err != nil {
        return nil, err
    }
    fields, err := q.bsonFields()
    if err != nil {
        return nil, err
    }
    return q.Conn.Find(q.Namespace, query, fields,
        q.Options.Limit, q.Options.Skip, 0)
}
