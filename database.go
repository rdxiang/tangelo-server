package main

import (
	"labix.org/v2/mgo"
	"log"
)

var db *mgo.Database

func init() {
	dbSession, err := mgo.Dial(DatabaseString)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	db = dbSession.DB(DatabaseName)
}
