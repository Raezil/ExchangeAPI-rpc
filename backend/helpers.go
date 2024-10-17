package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func UrlWithAPIKey() string {
	val, ok := os.LookupEnv("Exchange_API_KEY")
	if !ok {
		panic("api key not found")
	}
	return "https://v6.exchangerate-api.com/v6/" + val
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
