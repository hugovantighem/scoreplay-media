package infra

import (
	"context"
	"fmt"
	"myproject/app"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ app.MediaStore = &MediaStoreMongo{}

type MediaStoreMongo struct {
	client *mongo.Client
}

func NewMediaStore(
	client *mongo.Client,
) *MediaStoreMongo {
	return &MediaStoreMongo{
		client: client,
	}
}

func (x *MediaStoreMongo) Collection() *mongo.Collection {
	return x.client.Database("infra_poc").Collection("media")
}

func (x *MediaStoreMongo) CreateIndexes() error {
	_, err := x.Collection().Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "tags.name", Value: 1}},
		Options: &options.IndexOptions{Name: ToPtr("media.tag.id")},
	})
	if err != nil {
		return fmt.Errorf("create index failed: %w", err)
	}

	return nil
}

func (x *MediaStoreMongo) Search(ctx context.Context, criteria app.SearchMediaCriteria) ([]app.Media, error) {
	logrus.Infof("MediaStoreMongo.Search: %+v", criteria)
	coll := x.Collection()
	filter := bson.D{{Key: "tags.id", Value: criteria.TagID}}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []app.Media
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (x *MediaStoreMongo) Save(ctx context.Context, item app.Media) error {
	logrus.Infof("MediaStoreMongo.Save: %+v", item)
	coll := x.Collection()
	update := bson.D{{Key: "$set", Value: item}}

	_, err := coll.UpdateOne(ctx, bson.D{{Key: "_id", Value: item.Id}}, update, &options.UpdateOptions{Upsert: ToPtr(true)})
	if err != nil {
		return err
	}

	return nil
}
