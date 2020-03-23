package domain

import "github.com/evzpav/crypto-broker/pkg/config"

type Ticker interface{}

type Account interface{}

type Exchanger interface {
	SetExchange(string) bool
	GetExchange() string
	AvailableExchanges() []string
	GetTicker(exc string) (Ticker, error)
	GetBalance() (interface{}, error)
}

type Exchange struct {
	Config *config.Config
}
