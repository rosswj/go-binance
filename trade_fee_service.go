package binance

import (
	"context"
	"encoding/json"
	"github.com/adshao/go-binance/common"
)

// Fetch trade fee, values
type TradeFeeTickersService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *TradeFeeTickersService) Symbol(symbol string) *TradeFeeTickersService {
	s.symbol = &symbol
	return s
}

// TradeFee define trade fee info
type TradeFee struct {
	Symbol string  `json:"symbol"`
	Maker  float64 `json:"maker"`
	Taker  float64 `json:"taker"`
}

// Do send request
func (s *TradeFeeTickersService) Do(ctx context.Context, opts ...RequestOption) (res []*TradeFee, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/tradeFee.html",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*TradeFee{}, err
	}
	data = common.ToJSONList(data)
	res = make([]*TradeFee, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*TradeFee{}, err
	}
	return res, nil
}
