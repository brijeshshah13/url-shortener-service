package main

import (
	"github.com/brijeshshah13/url-shortener-service/models/dbs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := initDBConn(dbs.DBNames["main"]); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/health-check", func (ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	r.Run(":" + httpPort)
}

func initDBConn(dbName string) error {
	if err := dbs.ConnectDB(dbName); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
