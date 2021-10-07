package store

import (
	"context"
	"fmt"

	"github.com/call-stack/copy_store.git/rpc/copystore"
	"go.mongodb.org/mongo-driver/bson"
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
	content := DataStore{URL: content_url, Content: req.Content}
	bson_content, _ := bson.Marshal(content)
	h.storeInMongo(ctx, bson_content)

	return &copystore.PasteResp{Url: content_url}, nil
}
