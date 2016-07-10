package app

import (
	"labix.org/v2/mgo"
	_ "labix.org/v2/mgo/bson"
	"log"
	_ "sync"
	"time"
)

const (
	MongoDBHosts = "localhost:27027"
	AuthDatabase = "busapp"
	AuthUserName = ""
	AuthPassword = ""
	TestDatabase = "busapp_test"
)

type BusApp struct {
}

func (app *BusApp) GetDBSession() *mgo.Session {
	mongoDBDailInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDailInfo)
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	return mongoSession
}
