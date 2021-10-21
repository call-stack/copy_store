package copyserv

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/call-stack/copy_store.git/internal/core/domain"
	"github.com/call-stack/copy_store.git/internal/core/ports"
)

type service struct {
	repo ports.Repository
}

func New(repo ports.Repository) *service {
	return &service{repo: repo}
}

func (srv *service) StoreContent(ctx context.Context, content string) (string, error) {
	clientIP := ctx.Value("remote-addr").(string)
	ct_nano := time.Now().UnixNano()
	data_to_hash := clientIP + fmt.Sprint(ct_nano)
	md5_encoded := md5.Sum([]byte(data_to_hash))
	hash := base64.URLEncoding.EncodeToString(md5_encoded[:])[:7]
	content_url := fmt.Sprintf("http://127.0.0.1:8080/twirp/GetContent/%s", hash)
	store := domain.Store{
		URL:     content_url,
		Content: content,
		TTL:     "100",
	}

	err := srv.repo.Create(store)
	if err != nil {
		log.Fatal()
	}

	return content_url, nil
}

func (srv *service) GetContent(ctx context.Context, url string) (*domain.Store, error) {

	content, err := srv.repo.Get(url)
	if err != nil {
		log.Fatal()
	}

	return &content, nil
}
