package app_test

import (
	"context"
	"fmt"
	"myproject/app"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestSearchMedia(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		t.Run("mediaStore", func(t *testing.T) {
			tagID := uuid.NewString()
			// GIVEN an error from mediaStore Search
			ctrl := gomock.NewController(t)
			mediaStore := app.NewMockMediaStore(ctrl)
			mediaStore.EXPECT().Search(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("error"))

			// AND a search media criteria
			crit := app.SearchMediaCriteria{
				TagID: tagID,
			}

			// WHEN search media
			_, err := app.SearchMedia(context.Background(), mediaStore, crit)

			// THEN an error is raised
			require.Error(t, err)
		})
	})
	t.Run("success", func(t *testing.T) {
		// GIVEN a search media criteria
		tagID := uuid.NewString()
		crit := app.SearchMediaCriteria{
			TagID: tagID,
		}

		ctrl := gomock.NewController(t)
		mediaStore := app.NewMockMediaStore(ctrl)
		mediaStore.EXPECT().Search(gomock.Any(), gomock.Any()).Return(
			[]app.Media{
				{
					Id:   uuid.NewString(),
					Name: "my-media",
					Tags: []app.Tag{
						{
							Id:   tagID,
							Name: "my-tag",
						},
					},
					FileURL: "/file-url",
				},
			}, nil,
		)

		// WHEN search media
		result, err := app.SearchMedia(context.Background(), mediaStore, crit)

		// THEN no error is raised
		require.NoError(t, err)
		// AND result contains 1 entry
		assert.Len(t, result, 1)
	})

}
