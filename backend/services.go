package backend

import "net/http"

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
