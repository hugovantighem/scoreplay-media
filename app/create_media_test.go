package app_test

import (
	"context"
	"fmt"
	"myproject/app"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateMedia(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		t.Run("fileStorage", func(t *testing.T) {
			// GIVEN an error from fileStorage WriteFile
			ctrl := gomock.NewController(t)
			fileStorage := app.NewMockFileStorage(ctrl)
			fileStorage.EXPECT().WriteFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))
			tagID := uuid.NewString()

			// AND a create media command
			cmd := app.CreateMediaCmd{
				Name:   "my-media",
				TagIDs: []string{tagID},
			}

			// WHEN create media
			err := app.CreateMedia(context.Background(), nil, nil, fileStorage, cmd)

			// THEN an error is raised
			require.Error(t, err)
		})
		t.Run("tagStore", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			fileStorage := app.NewMockFileStorage(ctrl)
			fileStorage.EXPECT().WriteFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			tagStore := app.NewMockTagStore(ctrl)
			tagID := uuid.NewString()
			// GIVEN an error from tagStrore Search
			tagStore.EXPECT().Search(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("error"))

			// AND a create media command
			cmd := app.CreateMediaCmd{
				Name:   "my-media",
				TagIDs: []string{tagID},
			}

			// WHEN create media
			err := app.CreateMedia(context.Background(), nil, tagStore, fileStorage, cmd)

			// THEN an error is raised
			require.Error(t, err)
		})
		t.Run("mediaStore", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			fileStorage := app.NewMockFileStorage(ctrl)
			fileStorage.EXPECT().WriteFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			tagStore := app.NewMockTagStore(ctrl)
			tagID := uuid.NewString()
			tagStore.EXPECT().Search(gomock.Any(), gomock.Any()).Return([]app.Tag{
				{
					Id:   tagID,
					Name: "foobar",
				},
			}, nil)
			// GIVEN an error from mediaStore Save
			mediaStore := app.NewMockMediaStore(ctrl)
			mediaStore.EXPECT().Save(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))

			// AND a create media command
			cmd := app.CreateMediaCmd{
				Name:   "my-media",
				TagIDs: []string{tagID},
			}

			// WHEN create media
			err := app.CreateMedia(context.Background(), mediaStore, tagStore, fileStorage, cmd)

			// THEN an error is raised
			require.Error(t, err)
		})
	})
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		fileStorage := app.NewMockFileStorage(ctrl)
		fileStorage.EXPECT().WriteFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		tagStore := app.NewMockTagStore(ctrl)
		tagID := uuid.NewString()
		tagStore.EXPECT().Search(gomock.Any(), gomock.Any()).Return([]app.Tag{
			{
				Id:   tagID,
				Name: "foobar",
			},
		}, nil)
		mediaStore := app.NewMockMediaStore(ctrl)
		mediaStore.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

		// GIVEN a create media command
		cmd := app.CreateMediaCmd{
			Name:   "my-media",
			TagIDs: []string{tagID},
		}

		// WHEN create media
		err := app.CreateMedia(context.Background(), mediaStore, tagStore, fileStorage, cmd)

		// THEN no error is raised
		require.NoError(t, err)
	})

}
