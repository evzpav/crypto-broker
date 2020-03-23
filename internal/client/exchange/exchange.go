package exchange

import (
	"fmt"
	"strings"

	"time"

	"github.com/evzpav/crypto-broker/internal/domain"
	"github.com/evzpav/crypto-broker/pkg/config"
	"github.com/evzpav/crypto-broker/pkg/log"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/builder"
)

type client struct {
	exchanges        map[string]goex.API
	exchangeSelected goex.API
	log              log.Logger
	config           *config.Config
}

func New(log log.Logger, config *config.Config) *client {
	apiBuilder := builder.NewAPIBuilder().HttpTimeout(time.Duration(30) * time.Second)
	exchanges := make(map[string]goex.API)

	exc := config.Exchanges
	for _, cfg := range exc {
		api := apiBuilder.APIKey(cfg.PublicKey).APISecretkey(cfg.SecretKey).Build(cfg.ExchangeName)
		exchanges[cfg.ExchangeName] = api
	}

	return &client{
		exchanges:        exchanges,
		log:              log,
		exchangeSelected: exchanges[string(goex.BITFINEX)],
		config:           config,
	}
}

func (c *client) SetExchange(exchangeName string) bool {
	c.log.Debug().Sendf("setting exchange: %s", exchangeName)

	exc, ok := c.exchanges[exchangeName]
	if ok {
		c.exchangeSelected = exc
	}

	return ok
}

func (c *client) GetExchange() string {
	return c.exchangeSelected.GetExchangeName()
}

func (c *client) AvailableExchanges() []string {
	var exchanges []string
	for _, cfg := range c.config.Exchanges {
		exchanges = append(exchanges, cfg.ExchangeName)
	}

	return exchanges
}

func (c *client) GetTicker(symbol string) (domain.Ticker, error) {
	c.log.Debug().Sendf("get ticker for pair: %s", symbol)

	pair := assembleCurrencyPair("BTC", "USD")
	if symbol != "" {
		symbolSplit := strings.Split(symbol, "/")
		if len(symbolSplit) > 1 {
			pair = assembleCurrencyPair(symbolSplit[0], symbolSplit[1])
		}
	}

	ticker, err := c.exchangeSelected.GetTicker(pair)
	if err != nil {
		return nil, fmt.Errorf("Failed to get ticker: %+v", err)
	}

	return ticker, nil
}

func (c *client) GetBalance() (interface{}, error) {
	c.log.Debug().Sendf("get balance for: %s", c.exchangeSelected.GetExchangeName())

	account, err := c.exchangeSelected.GetAccount()
	if err != nil {
		return nil, err
	}

	return account, nil
}

func assembleCurrencyPair(baseCurrency, quoteCurrency string) goex.CurrencyPair {
	return goex.NewCurrencyPair2(fmt.Sprintf("%s_%s", baseCurrency, quoteCurrency))
}
func assembleCurrencyPairs(coins []string, quoteCurrency string) (pairs []goex.CurrencyPair) {
	for _, coin := range coins {
		pairs = append(pairs, goex.NewCurrencyPair2(fmt.Sprintf("%s_%s", coin, quoteCurrency)))
	}
	return pairs
}
