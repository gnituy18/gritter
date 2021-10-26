package mission

import (
	"errors"

	"gritter/pkg/context"
)

var (
	ErrNotFound = errors.New("mission not found")
)

type Store interface {
	Create(ctx context.Context, m *Mission) (string, error)
	Get(ctx context.Context, id string) (*Mission, error)
	Update(ctx context.Context, m *Mission) error
	Delete(ctx context.Context, id string) error
}
