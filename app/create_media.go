package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateMediaCmd struct {
	Name   string
	TagIDs []string
	File   []byte
}

// CreateMedia stores a newly created media.
func CreateMedia(ctx context.Context, store MediaStore, tagStore TagStore, fileStorage FileStorage, cmd CreateMediaCmd) error {
	logrus.Infof("CreateMedia: name= %q tagIDs= %+v", cmd.Name, cmd.TagIDs)
	id := uuid.NewString()

	err := fileStorage.WriteFile(ctx, cmd.File, id)
	if err != nil {
		return err
	}

	tags, err := tagStore.Search(ctx, SearchTagCriteria{
		TagIDs: cmd.TagIDs,
	})
	if err != nil {
		return fmt.Errorf("search failed: %w", err)
	}

	item, err := NewMedia(
		id,
		cmd.Name,
		tags,
	)
	if err != nil {
		return fmt.Errorf("invalid value: %w", err)
	}

	err = store.Save(ctx, item)
	if err != nil {
		return fmt.Errorf("save failed: %w", err)
	}

	return nil
}
