package app_test

import (
	"context"
	"fmt"
	"myproject/app"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateTag(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		t.Run("tagStore", func(t *testing.T) {
			// GIVEN an error from tagStore Save
			ctrl := gomock.NewController(t)
			tagStore := app.NewMockTagStore(ctrl)
			tagStore.EXPECT().Save(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))

			// AND a create tag command
			cmd := app.CreateTagCmd{
				Name: "my-tag",
			}

			// WHEN create tag
			err := app.CreateTag(context.Background(), tagStore, cmd)

			// THEN an error is raised
			require.Error(t, err)
		})
	})
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tagStore := app.NewMockTagStore(ctrl)
		tagStore.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

		// GIVEN a create tag command
		cmd := app.CreateTagCmd{
			Name: "my-tag",
		}

		// WHEN create tag
		err := app.CreateTag(context.Background(), tagStore, cmd)

		// THEN no error is raised
		require.NoError(t, err)
	})

}
