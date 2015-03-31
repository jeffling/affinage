package main

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

var Db *leveldb.DB

func init() {
	var err error
	dbPath := "./db"
	Db, err = leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatal("Could not open db in", dbPath)
	}
}
