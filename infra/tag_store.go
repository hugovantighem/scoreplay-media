package infra

import (
	"context"
	"myproject/app"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ app.TagStore = &TagStoreMongo{}

type TagStoreMongo struct {
	client *mongo.Client
}

func NewTagStore(
	client *mongo.Client,
) *TagStoreMongo {
	return &TagStoreMongo{
		client: client,
	}
}

func (x *TagStoreMongo) Collection() *mongo.Collection {
	return x.client.Database("infra_poc").Collection("tags")
}

func (x *TagStoreMongo) Search(ctx context.Context, criteria app.SearchTagCriteria) ([]app.Tag, error) {
	logrus.Infof("TagStoreMongo.Search: %+v", criteria)
	coll := x.Collection()
	filter := bson.D{}

	if len(criteria.TagIDs) > 0 {
		filter = append(filter, bson.E{Key: "_id", Value: bson.M{"$in": criteria.TagIDs}})
	}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []app.Tag
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (x *TagStoreMongo) Save(ctx context.Context, item app.Tag) error {
	coll := x.Collection()
	update := bson.D{{Key: "$set", Value: item}}

	_, err := coll.UpdateOne(ctx, bson.D{{Key: "_id", Value: item.Id}}, update, &options.UpdateOptions{Upsert: ToPtr(true)})
	if err != nil {
		return err
	}

	return nil
}
