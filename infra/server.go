package infra

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"myproject/api"
	"myproject/app"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	client      *mongo.Client
	fileStorage app.FileStorage
	tagStore    app.TagStore
	mediaStore  app.MediaStore
}

func NewServer(
	client *mongo.Client,
	fileStorage app.FileStorage,
	tagStore app.TagStore,
	mediaStore app.MediaStore,
) *Server {
	return &Server{
		client:      client,
		fileStorage: fileStorage,
		tagStore:    tagStore,
		mediaStore:  mediaStore,
	}
}

func (x *Server) GetPing(ctx context.Context, request api.GetPingRequestObject) (api.GetPingResponseObject, error) {
	return api.GetPing200JSONResponse{
		Ping: "pong",
	}, nil
}

func (x *Server) FindMedia(ctx context.Context, request api.FindMediaRequestObject) (api.FindMediaResponseObject, error) {
	items, err := app.SearchMedia(ctx, x.mediaStore, app.SearchMediaCriteria{TagID: request.Params.Tag})
	if err != nil {
		return nil, fmt.Errorf("search media failed: %w", err)
	}

	result, err := ToMediaDtos(items)
	if err != nil {
		return nil, fmt.Errorf("mapping to media DTO failed: %w", err)
	}

	return result, nil
}

func (x *Server) PostMedia(ctx context.Context, request api.PostMediaRequestObject) (api.PostMediaResponseObject, error) {
	cmd := app.CreateMediaCmd{}
	var err error
	for err == nil {
		var part *multipart.Part
		part, err = request.Body.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			}
			return api.PostMedia201JSONResponse{}, fmt.Errorf("error reading multipart: %w", err)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(part)

		if part.FormName() == "name" {
			cmd.Name = buf.String()
		}

		if part.FormName() == "tags" {
			cmd.TagIDs = strings.Split(buf.String(), ",")
		}

		if part.FormName() == "file" {
			cmd.File = buf.Bytes()
		}
	}

	err = app.CreateMedia(ctx, x.mediaStore, x.tagStore, x.fileStorage, cmd)
	if err != nil {
		return nil, fmt.Errorf("create media failed: %w", err)
	}

	return api.PostMedia201JSONResponse{}, nil
}

func (x *Server) GetTags(ctx context.Context, request api.GetTagsRequestObject) (api.GetTagsResponseObject, error) {
	items, err := app.ListTags(ctx, x.tagStore)
	if err != nil {
		return nil, fmt.Errorf("list tags failed: %w", err)
	}

	resp, err := ToTagDtos(items)
	if err != nil {
		return nil, fmt.Errorf("mapping to tag DTO failed: %w", err)
	}

	return resp, nil
}

func (x *Server) CreateTag(ctx context.Context, request api.CreateTagRequestObject) (api.CreateTagResponseObject, error) {
	err := app.CreateTag(ctx, x.tagStore, app.CreateTagCmd{
		Name: request.Body.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("create tag failed: %w", err)
	}

	return api.CreateTag201JSONResponse{}, nil
}
