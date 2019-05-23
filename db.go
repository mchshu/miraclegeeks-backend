package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

var session *dbr.Session

func initDB(driver, dsn string) {
	conn, err := dbr.Open(driver, dsn, &DBLogger{})
	if err != nil {
		log.Fatal(err)
	}
	session = conn.NewSession(nil)
}

func DB() *dbr.Session {
	return session
}

type DBLogger struct {
	dbr.NullEventReceiver
	MaxOutLen int
}

func (d DBLogger) EventErr(eventName string, err error) error {
	d.output(fmt.Sprintf("dbr event: %s, err: %s", eventName, err))
	return err
}

func (d DBLogger) EventErrKv(eventName string, err error, kvs map[string]string) error {
	d.output(fmt.Sprintf("dbr event: %s, err: %s, kvs: %+v", eventName, err, kvs))
	return err
}

func (d DBLogger) Timing(eventName string, nanoseconds int64) {
	dur := time.Duration(nanoseconds) * time.Nanosecond
	d.output(fmt.Sprintf("dbr timing: %s, duration: %v", eventName, dur))
}

func (d DBLogger) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	dur := time.Duration(nanoseconds) * time.Nanosecond
	d.output(fmt.Sprintf("dbr timing: %s, duration: %v, kvs: %+v", eventName, dur, kvs))
}

func (d DBLogger) output(s string) {
	l := len(s)
	if d.MaxOutLen > 0 && l > d.MaxOutLen {
		l = d.MaxOutLen
	}
	log.Printf(s[:l])
}
