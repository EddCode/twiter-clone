package repository

import (
	"context"
	"log"

	"github.com/EddCode/twitter-clone/cmd/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newMongoClient() (*mongo.Client, error) {
    setting, err := config.GetConfig("../cmd/config/config.yml")

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
