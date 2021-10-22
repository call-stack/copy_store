package main

import (
	"context"
	"net/http"

	copyserv "github.com/call-stack/copy_store.git/internal/core/services/copyserv"
	"github.com/call-stack/copy_store.git/internal/handlers/copyhdl"
	"github.com/call-stack/copy_store.git/internal/repo/copyrepo"
	"github.com/call-stack/copy_store.git/rpc/writestore"
)

var (
	ListenAddr = "localhost:8080"
	RedisAddr  = "localhost:6379"
	MongoURI   = "mongodb://localhost:27018"
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
	copyrepo := copyrepo.NewRepo()
	copyservice := copyserv.New(copyrepo)
	copyhadlr := copyhdl.NewWriterServer(copyservice)
	twip_handler := writestore.NewWriteStoreServer(copyhadlr)
	wrapped := withUserAgent(twip_handler)
	http.ListenAndServe(":5000", wrapped)
}
