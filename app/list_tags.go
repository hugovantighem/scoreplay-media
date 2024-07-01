package app

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

// ListTags retrieve all tags.
func ListTags(ctx context.Context, store TagStore) ([]Tag, error) {
	logrus.Info("ListTags")
	items, err := store.Search(ctx, SearchTagCriteria{})
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	return items, nil
}
