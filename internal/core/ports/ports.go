package ports

import (
	"context"

	"github.com/call-stack/copy_store.git/internal/core/domain"
)

type CopyStore interface {
	StoreContent(ctx context.Context, content string) (string, error)
	GetContent(ctx context.Context, url string) (*domain.Store, error)
}

type Repository interface {
	Get(id string) (domain.Store, error)
	Create(item domain.Store) error
}
