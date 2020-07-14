package binance

import (
	"context"
	"encoding/json"
	"github.com/rosswj/go-binance/common"
)

// Fetch trade fee, values
type TradeFeeService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *TradeFeeService) Symbol(symbol string) *TradeFeeService {
	s.symbol = &symbol
	return s
}

// TradeFee define trade fee info
type Fee struct {
	Symbol string  `json:"symbol"`
	Maker  float64 `json:"maker"`
	Taker  float64 `json:"taker"`
}

type TradeFeeResult struct {
	Success  bool  `json:"success"`
	TradeFee []Fee `json:"tradeFee"`
}

// Do send request
func (s *TradeFeeService) Do(ctx context.Context, opts ...RequestOption) (res []*TradeFeeResult, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/tradeFee.html",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*TradeFeeResult{}, err
	}
	data = common.ToJSONList(data)
	res = make([]*TradeFeeResult, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*TradeFeeResult{}, err
	}
	return res, nil
}
