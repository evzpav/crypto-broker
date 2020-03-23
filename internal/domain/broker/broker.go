package broker

import (
	"fmt"

	"github.com/evzpav/crypto-broker/internal/domain"
	"github.com/evzpav/crypto-broker/pkg/log"
)

type broker struct {
	b   *domain.Broker
	log log.Logger
}

func New(log log.Logger, handler domain.Handler, telegram domain.Telegrammer, exchange domain.Exchanger) *broker {
	return &broker{
		b: &domain.Broker{
			Telegram:       telegram,
			Exchange:       exchange,
			CommandHandler: handler,
		},
		log: log,
	}
}

func (b *broker) Run() {

	b.b.Telegram.Start()

	b.b.Telegram.AddKeyboard(b.b.CommandHandler.GetAllCommandsNames())

	msg := make(chan domain.TelegramMessage)

	go b.b.Telegram.CheckTelegramCommands(msg)

	for m := range msg {

		cmd, err := b.b.CommandHandler.Get(m.CommandName)
		if err != nil {
			b.log.Error().Err(err)
			b.b.Telegram.SendMessage(m.ChatID, err.Error())
			continue
		}

		var tgMsg = domain.TelegramMessage{
			ChatID:      m.ChatID,
			CommandName: m.CommandName,
			Args:        m.Args,
		}

		b.b.Message = tgMsg

		err = cmd.Execute(*b.b)
		if err != nil {
			b.log.Error().Err(err).Sendf("Failed to execute command: %s. Error: %+v", m.CommandName, err)
			b.b.Telegram.SendMessage(m.ChatID, fmt.Sprintf("Failed to execute command: %s. Error: %+v", m.CommandName, err))
			continue
		}
	}

	close(msg)

}
