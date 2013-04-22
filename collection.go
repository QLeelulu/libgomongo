package libgomongo

import (
    "errors"
)

// db
type DB struct {
    Name string // db name
    Conn *Mongo // connection
}

type Collection struct {
    // String with the format "<database>.<collection>" where <database> is the
    // name of the database and <collection> is the name of the collection.
    Name      string
    Namespace string
    Db        *DB
}

// get collection instance
func (db *DB) C(name string) *Collection {
    c := &Collection{
        Name:      name,
        Namespace: db.Name + "." + name,
        Db:        db,
    }
    return c
}

func (c *Collection) Find(query M) *Query {
    q := NewQuery(c.Db.Conn, c.Namespace)
    q.Spec.Query = query
    return &q
}

/**
 * Count the number of documents in a collection matching a query.
 *
 * @param the query.
 *
 * @return the number of matching documents. If the command fails,
 *     MONGO_ERROR is returned.
 */
func (c *Collection) Count(query M) (int64, error) {
    b := NewBson()
    b.Init()
    b.FromMap(query)
    b.Finish()
    defer b.Destroy()
    r := c.Db.Conn.Count(c.Db.Name, c.Name, b)
    if r == MONGO_ERROR {
        var errs string
        err := c.Db.Conn.Error()
        if err != nil {
            errs = err.Error()
        } else {
            errs = "Unknow Error."
        }
        return r, errors.New("MongoDb Count error: " + errs)
    }
    return r, nil
}

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
func (c *Collection) Insert(data M, writeConcern *MongoWriteConcern) (int, error) {
    b := NewBsonFromM(data)
    r := c.Db.Conn.Insert(c.Namespace, b, writeConcern)
    if r == MONGO_OK {
        return r, nil
    }
    return r, c.Db.Conn.Error()
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
func (c *Collection) Remove(cond M, writeConcern *MongoWriteConcern) (int, error) {
    b_cond := NewBsonFromM(cond)
    r := c.Db.Conn.Remove(c.Namespace, b_cond, writeConcern)
    if r == MONGO_OK {
        return r, nil
    }
    return r, c.Db.Conn.Error()
}
