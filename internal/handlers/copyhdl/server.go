package copyhdl

import (
	"context"
	"log"

	"github.com/call-stack/copy_store.git/internal/core/ports"
	pb "github.com/call-stack/copy_store.git/rpc/copystore"
)

type server struct {
	copystore ports.CopyStore
}

func Newserver(copystore ports.CopyStore) *server {
	return &server{
		copystore: copystore,
	}
}

func (sr *server) PasteContent(ctx context.Context, req *pb.PasteReq) (*pb.PasteResp, error) {
	url, err := sr.copystore.StoreContent(ctx, req.Content)
	if err != nil {
		return &pb.PasteResp{}, err
	}

	return &pb.PasteResp{Url: url}, nil
}

func (sr *server) GetContent(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	content, err := sr.copystore.GetContent(ctx, req.Url)

	if err != nil {
		log.Fatal()
	}

	return &pb.GetResp{Content: content.Content}, nil
}
