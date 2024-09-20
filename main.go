package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/rpc"
	jsonrpc "github.com/gorilla/rpc/json"
)

func UrlWithAPIKey() string {
	val, ok := os.LookupEnv("Exchange_API_KEY")
	if !ok {
		panic("api key not found")
	}
	return "https://v6.exchangerate-api.com/v6/" + val
}

func Latest(currency string) string {
	return UrlWithAPIKey() + "/latest/" + currency
}

func Exchange(left string, right string) string {
	return UrlWithAPIKey() + "/pair/" + left + "/" + right
}

func Request(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body, err
}

func GetLatestRequest(currency string) map[string]interface{} {
	url := Latest(currency)
	body, err := Request(url)
	if err != nil {
		panic(err)
	}
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonMap)
	return jsonMap
}

func GetExchangedRequest(from, to string) map[string]interface{} {
	url := Exchange(from, to)
	body, err := Request(url)
	if err != nil {
		panic(err)
	}
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonMap)
	return jsonMap
}

type CurrencyExchangeArgs struct {
	From string
	To   string
}

type CurrencyExchangeListArgs struct {
	Currency string
}

type CurrencyReply struct {
	Message map[string]interface{}
}

type CurrencyService struct{}

func (h *CurrencyService) Latest(r *http.Request, args *CurrencyExchangeListArgs, reply *CurrencyReply) error {
	reply.Message = GetLatestRequest(args.Currency)
	return nil
}

func (h *CurrencyService) Exchange(r *http.Request, args *CurrencyExchangeArgs, reply *CurrencyReply) error {
	reply.Message = GetExchangedRequest(args.From, args.To)
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(jsonrpc.NewCodec(), "application/json")
	s.RegisterService(new(CurrencyService), "")
	http.Handle("/rpc", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
