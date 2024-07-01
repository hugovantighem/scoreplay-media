package app

import (
	"fmt"
	"strings"
)

type Media struct {
	Id      string
	Name    string
	Tags    []Tag
	FileURL string
}

func NewMedia(id string, name string, tags []Tag) (Media, error) {
	if len(strings.TrimSpace(id)) == 0 {
		return Media{}, fmt.Errorf("invalid id: should not be empty")
	}

	if len(strings.TrimSpace(name)) == 0 {
		return Media{}, fmt.Errorf("invalid name: should not be empty")
	}

	if len(tags) == 0 {
		return Media{}, fmt.Errorf("invalid tag list: should not be empty")
	}

	// ensure non nil
	tagList := []Tag{}
	tagList = append(tagList, tags...)

	return Media{
		Id:      id,
		Name:    name,
		Tags:    tagList,
		FileURL: fmt.Sprintf("/assets/%s", id),
	}, nil
}

type Tag struct {
	Id   string
	Name string
}

func NewTag(id string, name string) (Tag, error) {
	if len(strings.TrimSpace(id)) == 0 {
		return Tag{}, fmt.Errorf("invalid id: should not be empty")
	}

	if len(strings.TrimSpace(name)) == 0 {
		return Tag{}, fmt.Errorf("invalid name: should not be empty")
	}

	return Tag{
		Id:   id,
		Name: name,
	}, nil
}
