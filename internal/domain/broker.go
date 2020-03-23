package domain

type Broker struct {
	Exchange Exchanger
	Telegram Telegrammer
	CommandHandler Handler
	Message        TelegramMessage
}
