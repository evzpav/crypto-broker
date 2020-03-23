package telegram

import (
	"github.com/evzpav/crypto-broker/internal/domain"
	"github.com/evzpav/crypto-broker/pkg/config"
	"github.com/evzpav/crypto-broker/pkg/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type client struct {
	updatesChannel tgbotapi.UpdatesChannel
	bot            *tgbotapi.BotAPI
	timeout        int
	keyboard       tgbotapi.ReplyKeyboardMarkup
	log            log.Logger
}

func New(log log.Logger, config *config.Config) *client {
	if config.Telegram.Token == "" {
		log.Fatal().Sendf("Telegram Token is required")
	}

	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		log.Fatal().Sendf("could not instantiate telegram bot: %v", err)
	}

	bot.Debug = config.Telegram.Debug
	err = tgbotapi.SetLogger(log)
	if err != nil {
		log.Error().Err(err).Sendf("could not instantiate telegram bot: %v", err)
	}

	return &client{
		bot:     bot,
		timeout: config.Telegram.Timeout,
		log:     log,
	}

}

func (tg *client) Start() {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = tg.timeout

	updates, err := tg.bot.GetUpdatesChan(u)
	if err != nil {
		tg.log.Error().Err(err).Sendf("could not get updates channel: %v", err)
		return
	}

	tg.updatesChannel = updates

	tg.log.Info().Sendf("Telegram bot started at channel: %s ", tg.bot.Self.FirstName)

}

func (tg *client) SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = tg.keyboard
	msg.ParseMode = "HTML"
	_, err := tg.bot.Send(msg)
	return err
}

func (tg *client) CheckTelegramCommands(msg chan domain.TelegramMessage) {
	for update := range tg.updatesChannel {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		var tgMsg domain.TelegramMessage

		if update.Message.IsCommand() {
			tg.log.Debug().Sendf("Command %s %s", update.Message.Command(), update.Message.CommandArguments())

			tgMsg.CommandName = update.Message.Command()
			tgMsg.Args = update.Message.CommandArguments()
			tgMsg.ChatID = update.Message.Chat.ID
			msg <- tgMsg
			continue
		}
	}
}

func (tg *client) AddKeyboard(commands []string) {

	var buttons []tgbotapi.KeyboardButton
	for _, cmdName := range commands {
		btn := tgbotapi.NewKeyboardButton("/" + cmdName)
		buttons = append(buttons, btn)
	}

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(buttons...),
	)
	tg.keyboard = keyboard
}
