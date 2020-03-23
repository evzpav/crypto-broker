package domain

type Telegrammer interface {
	Start()
	AddKeyboard(commands []string)
	SendMessage(chatID int64, text string) error
	CheckTelegramCommands(msg chan TelegramMessage)
}

type Telegram struct{}

type TelegramMessage struct {
	CommandName string
	Args        string
	ChatID      int64
	Text        string
}
