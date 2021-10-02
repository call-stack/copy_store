package server

import (
	"context"

	"github.com/call-stack/copy_store.git/internal/store"
	pb "github.com/call-stack/copy_store.git/rpc/copystore"
)

type CopyStoreServer struct{}

func (s *CopyStoreServer) PasteContent(ctx context.Context, req *pb.PasteReq) (*pb.PasteResp, error) {
	core := &store.StoreCore{}

	return core.StoreContent(ctx, req)
}

func (s *CopyStoreServer) GetContent(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	return &pb.GetResp{Content: "some content"}, nil
}
