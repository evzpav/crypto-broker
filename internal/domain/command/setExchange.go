package command

import (
	"fmt"

	"github.com/evzpav/crypto-broker/internal/domain"
)

type setExchange struct{}

func SetExchange() *setExchange {
	return &setExchange{}
}

func (cmd *setExchange) Execute(broker domain.Broker) error {
	exchangeName := broker.Message.Args

	success := broker.Exchange.SetExchange(exchangeName)
	msg := fmt.Sprintf("Exchange is: <b>%s</b>", broker.Exchange.GetExchange())
	if !success && exchangeName != "" {
		msg = fmt.Sprintf("Invalid exchange: %s. Exchange available are: %v", exchangeName, broker.Exchange.AvailableExchanges())
	}

	err := broker.Telegram.SendMessage(broker.Message.ChatID, msg)
	if err != nil {
		return err
	}

	return nil
}
