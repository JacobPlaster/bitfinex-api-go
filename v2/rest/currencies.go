package rest

import (
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"path"
	"strings"
)

// TODO

// TradeService manages the Trade endpoint.
type CurrenciesService struct {
	requestFactory
	Synchronous
}

// All returns all orders for the authenticated account.
func (cs *CurrenciesService) Conf(label, symbol, unit, explorer, pairs bool) ([]bitfinex.CurrencyConf, error) {
	segments := make([]string, 0)
	if label {
		segments = append(segments, bitfinex.CurrencyLabelMap)
	}
	if symbol {
		segments = append(segments, bitfinex.CurrencySymbolMap)
	}
	if unit {
		segments = append(segments, bitfinex.CurrencyUnitMap)
	}
	if explorer {
		segments = append(segments, bitfinex.CurrencyExplorerMap)
	}
	if pairs {
		segments = append(segments, bitfinex.CurrencyExchangeMap)
	}
	fmt.Println(path.Join("conf", strings.Join(segments,":")))
	req := NewRequestWithMethod(path.Join("conf", strings.Join(segments,",")), "GET")
	raw, err := cs.Request(req)
	if err != nil {
		return nil, err
	}
	// add mapping to raw data
	parsedRaw := make([]bitfinex.RawCurrencyConf, len(raw))
	for index, d := range raw {
		parsedRaw = append(parsedRaw, bitfinex.RawCurrencyConf{segments[index], d})
	}
	// parse to config object
	configs, err := bitfinex.NewCurrencyConfFromRaw(parsedRaw)
	if err != nil {
		return nil, err
	}
	return configs, nil
}

