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

func TestListTag(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		t.Run("tagStore", func(t *testing.T) {
			// GIVEN an error from tagstore Search
			ctrl := gomock.NewController(t)
			tagStore := app.NewMockTagStore(ctrl)
			tagStore.EXPECT().Search(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("error"))

			// WHEN listinig tags
			_, err := app.ListTags(context.Background(), tagStore)

			// THEN an error is raised
			require.Error(t, err)
		})
	})
	t.Run("success", func(t *testing.T) {
		// GIVEN one existing tag
		ctrl := gomock.NewController(t)
		tagStore := app.NewMockTagStore(ctrl)
		tagStore.EXPECT().Search(gomock.Any(), gomock.Any()).Return([]app.Tag{{Id: uuid.NewString(), Name: "my-tag"}}, nil)

		// WHEN listinig tags
		result, err := app.ListTags(context.Background(), tagStore)

		// THEN no error is raised
		require.NoError(t, err)
		//AND the result contains 1 entry
		assert.Len(t, result, 1)
	})

}
