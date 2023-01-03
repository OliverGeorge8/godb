package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Mongo struct {
	client *mongo.Client
	db     *mongo.Database
}

var instance *Mongo

func (m *Mongo) Db(db string) *mongo.Database {
	return m.client.Database(db)
}

func Initialize(ctx context.Context) *Mongo {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic("Error connecting to mongo")
	}

	return &Mongo{client: client, db: client.Database(dbName)}
}

func Get() *Mongo {
	if instance == nil {
		instance = Initialize(context.Background())
	}
	return instance
}
