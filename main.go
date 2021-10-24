package main

import (
	"fmt"
	"github.com/brijeshshah13/url-shortener-service/models/dbs"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	initDBConn(dbs.DBNames["main"])
}

func initDBConn(dbName string) *mongo.Database {
	conn, err := dbs.ConnectDB(dbName)
	if err != nil {
		panic(fmt.Sprintf("ERROR: db conn error: %v", err))
	}
	return conn
}
