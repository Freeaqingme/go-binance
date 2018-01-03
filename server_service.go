package binance

import (
	"context"
	"encoding/json"
)

// PingService ping server
type PingService struct {
	c *Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/ping",
	}
	_, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	return
}

// ServerTimeService get server time
type ServerTimeService struct {
	c *Client
}

// Do send request
func (s *ServerTimeService) Do(ctx context.Context, opts ...RequestOption) (serverTime int64, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/time",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	j, err := newJSON(data)
	if err != nil {
		return
	}
	serverTime = j.Get("serverTime").MustInt64()
	return
}

type ExchangeInfoService struct {
	c *Client
}

type SymbolFilter struct {
	FilterType  string `json:"filterType"`
	MinPrice    string `json:"minPrice"`
	MaxPrice    string `json:"maxPrice"`
	TickSize    string `json:"tickSize"`
	MinQty      string `json:"minQty"`
	MaxQty      string `json:"maxQty"`
	StepSize    string `json:"stepSize"`
	MinNotional string `json:"minNotional"`
}

type SymbolPriceFilter struct {
	FilterType string `json:"filterType"`
	MinPrice   string `json:"minPrice"`
	MaxPrice   string `json:"maxPrice"`
	TickSize   string `json:"tickSize"`
}

type SymbolSizeFilter struct {
	FilterType string `json:"filterType"`
	MinQty     string `json:"minQty"`
	MaxQty     string `json:"maxQty"`
	StepSize   string `json:"stepSize"`
}

type SymbolNotionalFilter struct {
	FilterType  string `json:"filterType"`
	MinNotional string `json:"minNotional"`
}

type SymbolUnit struct {
	Symbol             string         `json:"symbol"`
	Status             string         `json:"status"`
	BaseAsset          string         `json:"baseAsset"`
	BaseAssetPrecision int            `json:"baseAssetPrecision"`
	QuoteAsset         string         `json:"quoteAsset"`
	QuotePrecision     int            `json:"quotePrecision"`
	Filters            []SymbolFilter `json:"filters"`
}

type RateLimitInterval string

const (
	RateLimitIntervalMinute RateLimitInterval = "MINUTE"
	RateLimitIntervalHour   RateLimitInterval = "HOUR"
	RateLimitIntervalDay    RateLimitInterval = "DAY"
)

type RateLimitType string

const (
	RateLimitTypeRequests RateLimitType = "REQUESTS"
	RateLimitTypeOrders   RateLimitType = "ORDERS"
)

type RateLimit struct {
	Type     RateLimitType     `json:"rateLimitType"`
	Interval RateLimitInterval `json:"interval"`
	Limit    int               `json:"limit"`
}

type ExchangeInfo struct {
	Timezone   string       `json:"timezone"`
	ServerTime int64        `json:"serverTime"`
	RateLimits []RateLimit  `json:"rateLimits"`
	Symbols    []SymbolUnit `json:"symbols"`
}

// Do send request
func (s *ExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) (res *ExchangeInfo, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/exchangeInfo",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = new(ExchangeInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}

	return
}
