package vision

type LastQuoteResponse struct {
	Status string `json:"status"`
	Symbol string `json:"symbol"`
	Last LastQuote `json:"last"`
}

type LastQuote struct {
	AskPrice    float32 `json:"askprice"`
	AskSize     int32   `json:"asksize"`
	AskExchange int     `json:"askexchange"`
	BidPrice    float32 `json:"bidprice"`
	BidSize     int32   `json:"bidsize"`
	BidExchange int     `json:"bidexchange"`
	Timestamp   int64   `json:"timestamp"`
}