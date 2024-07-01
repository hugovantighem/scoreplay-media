package app

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

// SearchMedia finds media according to criteria.
func SearchMedia(ctx context.Context, store MediaStore, criteria SearchMediaCriteria) ([]Media, error) {
	logrus.Infof("SearchMedia: %+v", criteria)
	items, err := store.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	return items, nil
}
