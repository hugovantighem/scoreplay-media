package infra

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Should use one client by appliation
func InitDB(conf Config) (*mongo.Client, func()) {
	client, err := mongo.Connect(context.Background(), options.Client().
		ApplyURI(conf.DbConnString))
	if err != nil {
		log.Fatalln(err)
	}

	return client, func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}
}
