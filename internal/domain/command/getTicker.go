package command

import (
	"fmt"

	"github.com/evzpav/crypto-broker/internal/domain"
)

type getTicker struct{}

func GetTicker() *getTicker {
	return &getTicker{}
}

func (cmd *getTicker) Execute(broker domain.Broker) error {
	ticker, err := broker.Exchange.GetTicker(broker.Message.Args)
	if err != nil {
		return err
	}

	err = broker.Telegram.SendMessage(broker.Message.ChatID, fmt.Sprintf("%+v", ticker))
	if err != nil {
		return err
	}

	return nil
}
