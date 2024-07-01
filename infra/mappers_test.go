package infra_test

import (
	"myproject/app"
	"myproject/infra"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToTagDto(t *testing.T) {
	// GIVEN a tag model
	item := app.Tag{
		Id:   uuid.NewString(),
		Name: "a-tag",
	}

	// WHEN mapping to DTO
	result, err := infra.ToTagDto(item)
	require.NoError(t, err)
	dto, err := result.AsTag()
	require.NoError(t, err)

	// THEN dto fields are populated
	assert.Equal(t, item.Id, dto.Id)
	assert.Equal(t, item.Name, dto.Name)
}

func TestToTagDtos(t *testing.T) {
	// GIVEN list of tag models
	item := app.Tag{
		Id:   uuid.NewString(),
		Name: "a-tag",
	}
	items := []app.Tag{item}

	// WHEN mapping to DTO list
	result, err := infra.ToTagDtos(items)
	require.NoError(t, err)

	// THEN dto list contains 1 item
	assert.Len(t, result, 1)
}

func TestToMediaDto(t *testing.T) {
	// GIVEN a media model
	item := app.Media{
		Id:   uuid.NewString(),
		Name: "a-madia",
		Tags: []app.Tag{
			{
				Id:   uuid.NewString(),
				Name: "my-tag",
			},
		},
	}

	// WHEN mapping to DTO
	result, err := infra.ToMediaDto(item)
	require.NoError(t, err)
	dto, err := result.AsMedia()
	require.NoError(t, err)

	// THEN dto fields are populated
	assert.Equal(t, item.Id, dto.Id)
	assert.Equal(t, item.Name, dto.Name)
	assert.Equal(t, item.FileURL, dto.FileUrl)
	require.Len(t, dto.Tags, 1)
	assert.Equal(t, item.Tags[0].Name, dto.Tags[0])
}

func TestToMediaDtos(t *testing.T) {
	// GIVEN list of media models
	item := app.Media{
		Id:   uuid.NewString(),
		Name: "a-madia",
		Tags: []app.Tag{
			{
				Id:   uuid.NewString(),
				Name: "my-tag",
			},
		},
	}
	items := []app.Media{item}

	// WHEN mapping to DTO list
	result, err := infra.ToMediaDtos(items)
	require.NoError(t, err)

	// THEN dto list contains 1 item
	assert.Len(t, result, 1)
}
