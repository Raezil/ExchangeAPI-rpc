package backend

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
