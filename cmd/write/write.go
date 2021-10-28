package main

import (
	"log"
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

func main() {
	copyrepo := copyrepo.NewRepo()
	copyservice := copyserv.New(copyrepo)
	copyhadlr := copyhdl.NewWriterServer(copyservice)
	twip_handler := writestore.NewWriteStoreServer(copyhadlr)

	log.Println("Running in port 5000")
	http.ListenAndServe(":5000", twip_handler)
}
