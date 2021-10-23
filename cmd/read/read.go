package main

import (
	"fmt"
	"net/http"

	copyserv "github.com/call-stack/copy_store.git/internal/core/services/copyserv"
	"github.com/call-stack/copy_store.git/internal/handlers/copyhdl"
	"github.com/call-stack/copy_store.git/internal/repo/copyrepo"
	"github.com/call-stack/copy_store.git/rpc/readstore"
)

func main() {
	copyrepo := copyrepo.NewRepo()
	copyservice := copyserv.New(copyrepo)
	copyhadlr := copyhdl.NewReaderServer(copyservice)
	twip_handler := readstore.NewReadStoreServer(copyhadlr)
	fmt.Println("Running in port 5001...")
	http.ListenAndServe(":5001", twip_handler)
}
