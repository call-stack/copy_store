package copyrepo

import (
	"github.com/call-stack/copy_store.git/internal/core/domain"
)

type repo struct {
}

func NewRepo() *repo {
	return &repo{}
}

func (r *repo) Get(id string) (domain.Store, error) {
	return domain.Store{}, nil
}

func (r *repo) Create(item domain.Store) error {
	return nil
}
