package app_test

import (
	"myproject/app"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTag(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		t.Run("invalid ID", func(t *testing.T) {
			// GIVEN an invalid ID
			id := "   "
			name := "my-tag"

			// WHEN NewTag
			_, err := app.NewTag(id, name)

			// THEN an error is raised
			assert.Error(t, err)
		})
		t.Run("invalid name", func(t *testing.T) {
			id := uuid.NewString()
			// GIVEN an invalid name
			name := "   "

			// WHEN NewTag
			_, err := app.NewTag(id, name)

			// THEN an error is raised
			assert.Error(t, err)
		})
	})
	t.Run("success", func(t *testing.T) {
		// GIVEN valid properties
		id := uuid.NewString()
		name := "my-tag"

		// WHEN NewTag
		result, err := app.NewTag(id, name)

		// THEN no error is raised
		assert.NoError(t, err)
		// AND result attributes are populates
		assert.Equal(t, id, result.Id)
		assert.Equal(t, name, result.Name)
	})
}

func TestNewMedia(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		t.Run("invalid ID", func(t *testing.T) {
			// GIVEN valid properties
			id := "  "
			name := "my-madia"
			tags := []app.Tag{
				{
					Id:   uuid.NewString(),
					Name: "my-tag",
				},
			}

			// WHEN NewMedia
			_, err := app.NewMedia(id, name, tags)

			// THEN an error is raised
			assert.Error(t, err)
		})
		t.Run("invalid name", func(t *testing.T) {
			// GIVEN invalid name
			id := uuid.NewString()
			name := "  "
			tags := []app.Tag{
				{
					Id:   uuid.NewString(),
					Name: "my-tag",
				},
			}

			// WHEN NewMedia
			_, err := app.NewMedia(id, name, tags)

			// THEN an error is raised
			assert.Error(t, err)
		})
		t.Run("invalid tags", func(t *testing.T) {
			// GIVEN invalid tags
			id := uuid.NewString()
			name := "my-tag"
			tags := []app.Tag{}

			// WHEN NewMedia
			_, err := app.NewMedia(id, name, tags)

			// THEN an error is raised
			assert.Error(t, err)
		})
	})
	t.Run("success", func(t *testing.T) {
		// GIVEN valid properties
		id := uuid.NewString()
		name := "my-tag"
		tags := []app.Tag{
			{
				Id:   uuid.NewString(),
				Name: "my-tag",
			},
		}

		// WHEN NewMedia
		result, err := app.NewMedia(id, name, tags)

		// THEN no error is raised
		assert.NoError(t, err)
		// AND result attributes are populates
		assert.Equal(t, id, result.Id)
		assert.Equal(t, name, result.Name)
		assert.Len(t, result.Tags, 1)
	})
}
