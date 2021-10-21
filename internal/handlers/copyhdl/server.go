package copyhdl

import (
	"context"
	"fmt"

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
	if err := sr.copystore.StoreContent(ctx, req.Content); err != nil {
		return &pb.PasteResp{}, err
	}

	return &pb.PasteResp{Url: "some random url"}, nil
}

func (sr *server) GetContent(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	fmt.Print("Get some stateement")
	return nil, nil
}
