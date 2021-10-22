package copyhdl

import (
	"context"
	"log"

	"github.com/call-stack/copy_store.git/internal/core/ports"
	rs "github.com/call-stack/copy_store.git/rpc/readstore"
)

type reader struct {
	copystore ports.CopyStore
}

func NewReaderServer(copystore ports.CopyStore) *reader {
	return &reader{
		copystore: copystore,
	}
}

func (sr *reader) GetContent(ctx context.Context, req *rs.GetReq) (*rs.GetResp, error) {
	content, err := sr.copystore.GetContent(ctx, req.Url)

	if err != nil {
		log.Fatal()
	}

	return &rs.GetResp{Content: content.Content}, nil
}
