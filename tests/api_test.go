package tests

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"myproject/api"
	"myproject/common"
	"myproject/infra"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestScenario tests all enpoints
func TestScenario(t *testing.T) {
	conf := infra.Config{
		ServerAddr:   "0.0.0.0:8080",
		DbConnString: "mongodb://localhost:27017/default",
	}
	stop := infra.RunApplication(conf)
	defer stop()

	_, err := common.WithRetry(context.Background(),
		func(ctx context.Context) (any, error) {
			result, err := http.Get("http://" + conf.ServerAddr + "/ping")
			if err != nil {
				return nil, err
			}

			return result, nil
		}, func(_ context.Context, _ interface{}, err error) (bool, error) {
			if err != nil {
				return true, err
			}

			return false, nil
		}, 10, 100*time.Millisecond)
	require.NoError(t, err)

	client, err := api.NewClientWithResponses("http://" + conf.ServerAddr)
	require.NoError(t, err)

	ctx := context.Background()
	tagName := "my-tag-1" + uuid.NewString()
	_, err = client.CreateTag(ctx, api.CreateTagCommand{
		Name: tagName,
	})
	require.NoError(t, err)

	tags, err := client.GetTagsWithResponse(ctx)
	require.NoError(t, err)
	var selectedTag api.Tag
	for _, item := range *tags.JSON200 {
		tag, err := item.AsTag()
		require.NoError(t, err)
		if tag.Name == tagName {
			selectedTag = tag
		}
	}
	require.NotEmpty(t, selectedTag.Id)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part1, err := writer.CreateFormField("name")
	require.NoError(t, err)
	mediaName := "my-media" + uuid.NewString()
	io.Copy(part1, bytes.NewReader([]byte(mediaName)))

	part2, err := writer.CreateFormField("tags")
	require.NoError(t, err)
	io.Copy(part2, bytes.NewReader([]byte(selectedTag.Id)))

	part3, err := writer.CreateFormField("file")
	require.NoError(t, err)
	io.Copy(part3, bytes.NewReader([]byte("hello-world"+uuid.NewString())))

	writer.Close()
	_, err = client.PostMediaWithBody(ctx, writer.FormDataContentType(), body)
	require.NoError(t, err)

	mediaList, err := client.FindMediaWithResponse(ctx, &api.FindMediaParams{
		Tag: selectedTag.Id,
	})
	require.NoError(t, err)
	assert.Len(t, *mediaList.JSON200, 1)
}
