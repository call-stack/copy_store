package store

import (
	"context"
	"fmt"

	"github.com/call-stack/copy_store.git/rpc/copystore"
)

type StoreCore struct{}

func (sc *StoreCore) StoreContent(ctx context.Context, req *copystore.PasteReq) (*copystore.PasteResp, error) {
	val := &Validator{}
	if err := val.ValidatePasteReq(req); err != nil {
		return &copystore.PasteResp{}, nil
	}
	client_ip := ctx.Value("remote-addr").(string)

	h := &Helper{}
	hash := h.getUniqueHash(client_ip)
	b, _ := h.isHashPresentInRedis(ctx, hash)
	if !b {
		h.setHashInRedis(ctx, hash)

	}

	content_url := fmt.Sprintf("http://127.0.0.1:8080/twirp/GetContent/%s", hash)

	return &copystore.PasteResp{Url: content_url}, nil
}
