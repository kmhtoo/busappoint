package app

import (
	"labix.org/v2/mgo"
	_ "labix.org/v2/mgo/bson"
	"log"
	_ "sync"
	"time"
)

type BusApp struct {
}

func (app *BusApp) GetDBSession() *mgo.Session {
	mongoDbHost := ParamString("db.host")
	confTimeout := ParamInt("db.timeout")
	timeout := time.Duration(confTimeout)
	authDatabase := ParamString("db.name")
	authUsername := ParamString("db.username")
	authPassword := ParamString("db.password")
	mongoDBDailInfo := &mgo.DialInfo{
		Addrs:    []string{mongoDbHost},
		Timeout:  timeout * time.Second,
		Database: authDatabase,
		Username: authUsername,
		Password: authPassword,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDailInfo)
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	return mongoSession
}
