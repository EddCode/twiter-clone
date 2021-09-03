package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	connection *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{connection: db}
}

func DBCollection(connection *mongo.Client, collectionName string) (collection *mongo.Collection, ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	db := connection.Database("twitterclone")
	collection = db.Collection(collectionName)

	return collection, ctx, cancel
}
