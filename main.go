package main

import (
	. "backend"
	"log"
	"net/http"

	"github.com/gorilla/rpc"
	jsonrpc "github.com/gorilla/rpc/json"
)

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(jsonrpc.NewCodec(), "application/json")
	s.RegisterService(new(CurrencyService), "")
	http.Handle("/rpc", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
