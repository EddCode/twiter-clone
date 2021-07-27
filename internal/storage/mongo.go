package storage

import (
	"context"
	"log"

	"github.com/EddCode/twitter-clone/cmd/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

var db *Database

func (db Database) getConnection() (*mongo.Client, error) {
	setting, err := config.GetConfig()

	if err != nil {
		panic("fail load config")
	}

	options := options.Client().ApplyURI(setting.Database.URL)
	client, mongoErr := mongo.Connect(context.TODO(), options)

	if mongoErr != nil {
		log.Fatal(mongoErr.Error())
		return nil, mongoErr
	}

	mongoErr = client.Ping(context.TODO(), nil)

	if mongoErr != nil {
		log.Fatal(mongoErr.Error())
		return nil, mongoErr
	}

	log.Println("Succesful conextion into DB")
	return client, nil

}

func NewMongoClient() (*Database, error) {
	log.Println("Creting db conecction")

	if db == nil {
		db = &Database{}
		client, err := db.getConnection()

		db.Client = client

		if err != nil {
			return nil, err
		}

		log.Printf("creating new db conecction %+v \n", db.Client)
		return db, nil
	}

	return db, nil
}
