package command

import (
	"fmt"

	"github.com/evzpav/crypto-broker/internal/domain"
)

type getBalance struct{}

func GetBalance() *getBalance {
	return &getBalance{}
}

func (cmd *getBalance) Execute(broker domain.Broker) error {
	balance, err := broker.Exchange.GetBalance()
	if err != nil {
		return err
	}

	err = broker.Telegram.SendMessage(broker.Message.ChatID, fmt.Sprintf("%+v", balance))
	if err != nil {
		return err
	}

	return nil
}
