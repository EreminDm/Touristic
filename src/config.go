package main

import (
	"gopkg.in/mgo.v2"
)

func Connect() (*mgo.Database, error) {
	host := "mongodb://localhost:27017"
	dbName := "demo"
	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}
	db := session.DB(dbName)
	return db, nil
}
