package models

import (
	"github.com/brijeshshah13/url-shortener-service/models/dbs"
	"github.com/brijeshshah13/url-shortener-service/models/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type URLSchema struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	IsActive       bool               `bson:"is_active,omitempty"`
	IsUsed         bool               `bson:"is_used,omitempty"`
	OriginalURL    string             `bson:"original_url,omitempty"`
	CreationDate   primitive.DateTime `bson:"creation_date,omitempty"`
	ExpirationDate primitive.DateTime `bson:"expiration_date,omitempty"`
}

var URL *mongo.Collection

func init() {
	if url, err := dbs.GetModel(utils.CollectionNames["url"]); err != nil {
		log.Fatal(err)
	} else {
		URL = url
	}
}
