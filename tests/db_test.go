package tests_test

import (
	"context"
	"myproject/app"
	"myproject/infra"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMongoDBMedia(t *testing.T) {
	uri := "mongodb://localhost:27017/default"

	client, close := infra.InitDB(infra.Config{
		DbConnString: uri,
	})

	defer close()

	tagStore := infra.NewTagStore(client)
	err := tagStore.Collection().Drop(context.Background())
	assert.NoError(t, err)

	mediaStore := infra.NewMediaStore(client)
	err = mediaStore.Collection().Drop(context.Background())
	assert.NoError(t, err)

	err = mediaStore.CreateIndexes()
	assert.NoError(t, err)

	err = tagStore.Save(context.Background(), app.Tag{Id: uuid.NewString(), Name: "foo"})
	assert.NoError(t, err)
	err = tagStore.Save(context.Background(), app.Tag{Id: uuid.NewString(), Name: "bar"})
	assert.NoError(t, err)

	id := uuid.NewString()
	tagId := uuid.NewString()
	item, err := app.NewMedia(
		id, "foo", []app.Tag{
			{
				Id:   tagId,
				Name: "bar",
			},
			{
				Id:   uuid.NewString(),
				Name: "foobar",
			},
		},
	)
	require.NoError(t, err)

	ctx := context.Background()
	err = mediaStore.Save(ctx, item)
	require.NoError(t, err)

	result, err := mediaStore.Search(ctx, app.SearchMediaCriteria{TagID: tagId})
	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, item.Id, result[0].Id)
	assert.Equal(t, item.Name, result[0].Name)
}

func TestMongoDBTags(t *testing.T) {
	uri := "mongodb://localhost:27017/default"

	client, close := infra.InitDB(infra.Config{
		DbConnString: uri,
	})

	defer close()

	tagStore := infra.NewTagStore(client)
	err := tagStore.Collection().Drop(context.Background())
	assert.NoError(t, err)

	id := uuid.NewString()

	err = tagStore.Save(context.Background(), app.Tag{Id: id, Name: "foo"})
	assert.NoError(t, err)
	err = tagStore.Save(context.Background(), app.Tag{Id: uuid.NewString(), Name: "bar"})
	assert.NoError(t, err)

	result, err := tagStore.Search(context.Background(), app.SearchTagCriteria{})
	require.NoError(t, err)
	assert.Len(t, result, 2)

	result, err = tagStore.Search(context.Background(), app.SearchTagCriteria{TagIDs: []string{id}})
	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "foo", result[0].Name)
}
