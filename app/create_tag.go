package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateTagCmd struct {
	Name string
}

// CreateTag stores a newly created tag.
func CreateTag(ctx context.Context, store TagStore, cmd CreateTagCmd) error {
	logrus.Infof("CreateTag: %+v", cmd)
	item, err := NewTag(uuid.NewString(), cmd.Name)
	if err != nil {
		return fmt.Errorf("invalid value: %w", err)
	}

	err = store.Save(ctx, item)
	if err != nil {
		return fmt.Errorf("save failed: %w", err)
	}

	return nil
}
