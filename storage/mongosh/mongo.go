package mongosh

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"question/config"
)

func Connect(ctx context.Context) (*mongo.Database, error) {
	conf := config.LoadConfig()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.MDB_ADDRESS).SetAuth(options.Credential{Username: "root", Password: "example"}))
	if err != nil {
		return nil, err
	}

	db := client.Database(conf.MDB_NAME)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return db, nil
}
