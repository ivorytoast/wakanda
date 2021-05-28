package vision

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AlpacaTestSuite struct {
	suite.Suite
}

func TestAlpacaTestSuite(t *testing.T) {
	suite.Run(t, new(AlpacaTestSuite))
}

func (s *AlpacaTestSuite) TestAlpaca() {
	actualLastQuote, err := GetLastQuote("AAPL")
	if err != nil {
		fmt.Println("Found an error!")
	} else {
		fmt.Println(actualLastQuote.Last.AskSize)
		fmt.Println(actualLastQuote.Last.BidSize)
		fmt.Println(actualLastQuote.Last.AskPrice)
		fmt.Println(actualLastQuote.Last.BidPrice)
		fmt.Println(actualLastQuote.Last.BidExchange)
		fmt.Println(actualLastQuote.Last.AskExchange)
		fmt.Println(actualLastQuote.Last.Timestamp)
		fmt.Println(actualLastQuote.Symbol)
		fmt.Println(actualLastQuote.Status)
	}
}