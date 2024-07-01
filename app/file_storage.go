package app

import (
	"context"
)

//go:generate mockgen -package=app -source=file_storage.go -destination=file_storage.mock.go
type FileStorage interface {
	WriteFile(ctx context.Context, content []byte, filename string) error
}
