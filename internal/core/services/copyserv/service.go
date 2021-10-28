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

const HOST_IP = "172.105.39.53"

type service struct {
	repo ports.Repository
}

func New(repo ports.Repository) *service {
	return &service{repo: repo}
}

func (srv *service) StoreContent(ctx context.Context, content string, remoteAddr string) (string, error) {
	fmt.Println(remoteAddr)
	ct_nano := time.Now().UnixNano()
	data_to_hash := remoteAddr + fmt.Sprint(ct_nano)
	md5_encoded := md5.Sum([]byte(data_to_hash))
	URLHash := base64.URLEncoding.EncodeToString(md5_encoded[:])[:7]

	store := domain.Store{
		Hash:    URLHash,
		Content: content,
		TTL:     "100",
	}

	err := srv.repo.Create(store)
	if err != nil {
		log.Fatal()
	}

	URL := fmt.Sprintf("%s/%s", HOST_IP, URLHash)
	return URL, nil
}

func (srv *service) GetContent(ctx context.Context, hash string) (*domain.Store, error) {

	content, err := srv.repo.Get(hash)
	if err != nil {
		log.Fatal()
	}

	return &content, nil
}
