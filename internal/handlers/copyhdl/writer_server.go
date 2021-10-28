package copyhdl

import (
	"context"

	"github.com/call-stack/copy_store.git/internal/core/ports"
	ws "github.com/call-stack/copy_store.git/rpc/writestore"
)

type writer struct {
	copystore ports.CopyStore
}

func NewWriterServer(copystore ports.CopyStore) *writer {
	return &writer{
		copystore: copystore,
	}
}

func (sr *writer) PasteContent(ctx context.Context, req *ws.PasteReq) (*ws.PasteResp, error) {
	url, err := sr.copystore.StoreContent(ctx, req.Content, req.RemoteAddress)
	if err != nil {
		return &ws.PasteResp{}, err
	}

	return &ws.PasteResp{Url: url}, nil
}
