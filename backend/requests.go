package backend

func GetLatestRequest(currency string) (map[string]interface{}, error) {
	url := Latest(currency)
	return ProcessRequest(url)
}

func GetExchangedRequest(from, to string) (map[string]interface{}, error) {
	url := Exchange(from, to)
	return ProcessRequest(url)
}

func GetEnrichedDataRequest(from, to string) (map[string]interface{}, error) {
	url := EnrichedData(from, to)
	return ProcessRequest(url)
}

func GetHistoryRequest(currency, year, month, day string) (map[string]interface{}, error) {
	url := History(currency, year, month, day)
	return ProcessRequest(url)
}
