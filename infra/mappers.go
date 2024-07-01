package infra

import (
	"fmt"
	"myproject/api"
	"myproject/app"
)

func ToTagDtos(items []app.Tag) (api.GetTags200JSONResponse, error) {
	result := api.GetTags200JSONResponse{}
	for _, item := range items {
		dto, err := ToTagDto(item)
		if err != nil {
			return nil, fmt.Errorf("list tags failed (tag mapping): %w", err)
		}

		result = append(result, dto)
	}

	return result, nil
}

func ToTagDto(item app.Tag) (api.TagListResult_Item, error) {
	result := api.TagListResult_Item{}

	err := result.FromTag(api.Tag{
		Id:   item.Id,
		Name: item.Name,
	})
	if err != nil {
		return api.TagListResult_Item{}, fmt.Errorf("list tags failed (tag mapping): %w", err)
	}
	return result, nil
}

func ToMediaDtos(items []app.Media) (api.FindMedia200JSONResponse, error) {
	result := api.FindMedia200JSONResponse{}
	for _, item := range items {
		dto, err := ToMediaDto(item)
		if err != nil {
			return nil, fmt.Errorf("list media failed (media mapping): %w", err)
		}

		result = append(result, dto)
	}

	return result, nil
}

func ToMediaDto(item app.Media) (api.MediaListResult_Item, error) {
	result := api.MediaListResult_Item{}

	tags := []string{}
	for _, tag := range item.Tags {
		tags = append(tags, tag.Name)
	}

	err := result.FromMedia(api.Media{
		Id:      item.Id,
		Name:    item.Name,
		FileUrl: item.FileURL,
		Tags:    tags,
	})
	if err != nil {
		return api.MediaListResult_Item{}, fmt.Errorf("list media failed (media mapping): %w", err)
	}
	return result, nil
}
