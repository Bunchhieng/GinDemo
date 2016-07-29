package main

import (
	"gopkg.in/mgo.v2"
)

var (
	mgoSession   *mgo.Session
	databaseName = "test"
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("mongodb://localhost:27017")
		if err != nil {
			panic(err) // no, not really
		}
	}
	return mgoSession.Clone()
}

func withCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(databaseName).C(collection)
	return s(c)
}
