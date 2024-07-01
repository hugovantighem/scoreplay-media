package app

import (
	"context"
	"fmt"
)

//go:generate mockgen -package=app -source=tag_store.go -destination=tag_store.mock.go
type TagStore interface {
	// -- Read

	// Search for all tags.
	// Note: pagination is not implemented
	Search(ctx context.Context, criteria SearchTagCriteria) ([]Tag, error)

	// -- Write

	Save(ctx context.Context, item Tag) error
}

type SearchTagCriteria struct {
	TagIDs []string
}

func (x SearchTagCriteria) String() string {
	return fmt.Sprintf("tagID=%v", x.TagIDs)
}
