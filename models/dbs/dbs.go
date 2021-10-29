package dbs

import (
	"context"
	"fmt"
	"github.com/brijeshshah13/url-shortener-service/config/environments"
	"github.com/brijeshshah13/url-shortener-service/models/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	DBNames = map[string]string{
		"main": "urlshortener",
	}
)

var allowedDBNames = make(map[string]struct{})

var collectionRegistry = map[string]string{
	utils.CollectionNames["url"]: DBNames["main"],
}

type dbRegistryConfig struct {
	connected  bool
	connection mongo.Client
	uri        string
	opts       environments.OptionsConfig
}

var dbRegistry = map[string]dbRegistryConfig{
	DBNames["main"]: {
		connected:  false,
		connection: mongo.Client{},
		uri:        environments.Mongo["main"].URI,
		opts:       environments.Mongo["main"].Options,
	},
}

func init() {
	// initialize list of allowed db names
	for _, v := range DBNames {
		allowedDBNames[v] = struct{}{}
	}
}

func ConnectDB(dbName string) error {
	if _, ok := allowedDBNames[dbName]; !ok {
		return fmt.Errorf("invalid db name: %s", dbName)
	}
	if db, ok := dbRegistry[dbName]; ok {
		if db.connected {
			log.Printf("db %s already connected", dbName)
			return nil
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.uri))
		if err != nil {
			log.Printf("error connecting to db: %s", err)
			return err
		}
		db.connected = true
		db.connection = *client
		return nil
	}
	return fmt.Errorf("db %s not found", dbName)
}

func GetModel(collectionName string) (*mongo.Collection, error) {
	dbName, ok := collectionRegistry[collectionName]
	if !ok {
		return nil, fmt.Errorf("collection %s not found", collectionName)
	}
	if _, ok := allowedDBNames[dbName]; !ok {
		return nil, fmt.Errorf("invalid db name: %s", dbName)
	}
	if db, ok := dbRegistry[dbName]; ok {
		if !db.connected {
			err := ConnectDB(dbName)
			if err != nil {
				return nil, err
			}
		}
		return db.connection.Database(dbName).Collection(collectionName), nil
	}
	return nil, fmt.Errorf("db %s not found", dbName)
}
