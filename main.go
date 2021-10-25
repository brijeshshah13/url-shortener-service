package main

import (
	"github.com/brijeshshah13/url-shortener-service/models/dbs"
	"log"
)

func main() {
	if err := initDBConn(dbs.DBNames["main"]); err != nil {
		log.Fatal(err)
	}
}

func initDBConn(dbName string) error {
	if err := dbs.ConnectDB(dbName); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
