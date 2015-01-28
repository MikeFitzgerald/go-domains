package domains

import (
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
)

/*
I want to use a different database for my tests,
so I'll embed *mgo.Session and store the database name.
*/
type DatabaseSession struct {
	*mgo.Session
	databaseName string
}

/*
Connect to the local MongoDB and set up the database.
*/
func NewSession(name string) *DatabaseSession {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	addIndexToDomainNames(session.DB(name))
	return &DatabaseSession{session, name}
}

/*
Add a unique index on the "domain_name" field.
*/
func addIndexToDomainNames(db *mgo.Database) {
	index := mgo.Index{
		Key:      []string{"domain_name"},
		Unique:   true,
		DropDups: true,
	}
	indexErr := db.C("domains").EnsureIndex(index)
	if indexErr != nil {
		panic(indexErr)
	}
}

/*
Martini lets you inject parameters for routing handlers
by using `context.Map()`. I'll pass each route handler
a instance of a *mgo.Database, so they can retrieve
and insert domains to and from that database.

For more information, check out:
http://blog.gopheracademy.com/day-11-martini
*/
func (session *DatabaseSession) Database() martini.Handler {
	return func(context martini.Context) {
		s := session.Clone()
		context.Map(s.DB(session.databaseName))
		defer s.Close()
		context.Next()
	}
}
