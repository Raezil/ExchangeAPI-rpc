package main

import (
	"encoding/json"
	"fmt"
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

func Exchange(from string, to string) string {
	return UrlWithAPIKey() + "/pair/" + from + "/" + to
}

func EnrichedData(from string, to string) string {
	return UrlWithAPIKey() + "/enriched/" + from + "/" + to
}

func History(currency, year, month, day string) string {
	return UrlWithAPIKey() + "/history/" + currency + "/" + year + "/" + month + "/" + day
}

func Request(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return body, err
}

func GetLatestRequest(currency string) (map[string]interface{}, error) {
	url := Latest(currency)
	val, err := ProcessRequest(url)
	if err != nil {
		log.Printf("Error in processing request: %v", err)
		return nil, err
	}
	return val, nil
}

func GetExchangedRequest(from, to string) (map[string]interface{}, error) {
	url := Exchange(from, to)
	val, err := ProcessRequest(url)
	if err != nil {
		log.Printf("Error in processing request: %v", err)
		return nil, err
	}
	return val, nil
}

func GetEnrichedDataRequest(from, to string) (map[string]interface{}, error) {
	url := EnrichedData(from, to)
	val, err := ProcessRequest(url)
	if err != nil {
		log.Printf("Error in processing request: %v", err)
		return nil, err
	}
	return val, nil
}

func ProcessRequest(url string) (map[string]interface{}, error) {
	body, err := Request(url)
	if err != nil {
		return nil, err
	}
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(body, &jsonMap); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return jsonMap, nil
}

func GetHistoryRequest(currency, year, month, day string) (map[string]interface{}, error) {
	url := History(currency, year, month, day)
	val, err := ProcessRequest(url)
	if err != nil {
		log.Printf("Error in processing request: %v", err)
		return nil, err
	}
	return val, nil
}

type CurrencyExchangeArgs struct {
	From string
	To   string
}

type CurrencyHistoryArgs struct {
	Currency string
	Year     string
	Month    string
	Day      string
}

type CurrencyExchangeListArgs struct {
	Currency string
}

type CurrencyReply struct {
	Message map[string]interface{}
}

type CurrencyService struct{}

func (h *CurrencyService) Latest(r *http.Request, args *CurrencyExchangeListArgs, reply *CurrencyReply) error {
	var err error
	reply.Message, err = GetLatestRequest(args.Currency)
	if err != nil {
		return err
	}
	return nil
}

func (h *CurrencyService) Exchange(r *http.Request, args *CurrencyExchangeArgs, reply *CurrencyReply) error {
	var err error
	reply.Message, err = GetExchangedRequest(args.From, args.To)
	if err != nil {
		return err
	}
	return nil
}

func (h *CurrencyService) EnrichedData(r *http.Request, args *CurrencyExchangeArgs, reply *CurrencyReply) error {
	var err error
	reply.Message, err = GetEnrichedDataRequest(args.From, args.To)
	if err != nil {
		return err
	}
	return nil
}

func (h *CurrencyService) History(r *http.Request, args *CurrencyHistoryArgs, reply *CurrencyReply) error {
	var err error
	reply.Message, err = GetHistoryRequest(args.Currency, args.Year, args.Month, args.Day)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(jsonrpc.NewCodec(), "application/json")
	s.RegisterService(new(CurrencyService), "")
	http.Handle("/rpc", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
