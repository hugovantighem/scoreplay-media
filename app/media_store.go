package app

import (
	"context"
	"fmt"
)

//go:generate mockgen -package=app -source=media_store.go -destination=media_store.mock.go
type MediaStore interface {
	// -- Read

	// Search for media by criteria.
	// Note: pagination is not implemented
	Search(ctx context.Context, criteria SearchMediaCriteria) ([]Media, error)

	// -- Write

	Save(ctx context.Context, item Media) error
}

type SearchMediaCriteria struct {
	TagID string
}

func (x SearchMediaCriteria) String() string {
	return fmt.Sprintf("tagID=%s", x.TagID)
}
