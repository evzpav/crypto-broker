package command

import (
	"fmt"

	"github.com/evzpav/crypto-broker/internal/domain"
)

type commandHandler struct {
	commands map[string]domain.Commander
}

func New() *commandHandler {
	handler := &commandHandler{
		commands: make(map[string]domain.Commander),
	}

	handler.Add("exchange", SetExchange())
	handler.Add("ticker", GetTicker())
	handler.Add("balance", GetBalance())

	return handler
}

func (ch *commandHandler) Add(name string, command domain.Commander) {
	ch.commands[name] = command
}

func (ch *commandHandler) Get(name string) (domain.Commander, error) {
	cmd, ok := ch.commands[name]
	if !ok {
		return nil, fmt.Errorf("Invalid command %s", name)
	}

	return cmd, nil
}

func (ch *commandHandler) GetAllCommandsNames() []string {
	var commandsNames []string
	for cmdName := range ch.commands {
		commandsNames = append(commandsNames, cmdName)
	}

	return commandsNames
}
