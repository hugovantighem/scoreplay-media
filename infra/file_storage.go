package infra

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"myproject/app"
	"os"
	"path/filepath"
)

var _ app.FileStorage = &FileStorage{}

type FileStorage struct {
	dir string
}

func NewFileStorage(dir string) *FileStorage {
	return &FileStorage{
		dir: dir,
	}
}

func (x *FileStorage) WriteFile(ctx context.Context, content []byte, filename string) error {
	dst, err := os.Create(filepath.Join(x.dir, filepath.Base(filename)))
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, bytes.NewReader(content)); err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}

	return nil
}
