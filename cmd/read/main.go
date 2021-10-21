package main

import (
	"context"
	"net/http"

	mdb "github.com/call-stack/copy_store.git/internal/mongo"
	rdb "github.com/call-stack/copy_store.git/internal/redis"
	"github.com/call-stack/copy_store.git/internal/server"
	"github.com/call-stack/copy_store.git/rpc/copystore"
)

func withUserAgent(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ra := r.RemoteAddr
		ctx = context.WithValue(ctx, "remote-addr", ra)
		r = r.WithContext(ctx)

		base.ServeHTTP(w, r)
	})
}

func main() {
	twip_handler := copystore.NewCopyStoreServer(&server.CopyStoreServer{})
	//some value to set when we initialize.
	rclient := rdb.RedisCore{}
	rclient.SetRedisClient()

	mClient := mdb.MongoCore{}
	mClient.SetMongoClient()

	wrapped := withUserAgent(twip_handler)
	http.ListenAndServe(":8080", wrapped)
}
