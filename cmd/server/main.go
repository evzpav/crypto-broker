package main

import (
	"os"

	"github.com/evzpav/crypto-broker/internal/client/exchange"
	"github.com/evzpav/crypto-broker/internal/client/telegram"
	"github.com/evzpav/crypto-broker/internal/domain/broker"
	"github.com/evzpav/crypto-broker/internal/domain/command"

	"github.com/evzpav/crypto-broker/pkg/config"
	"github.com/evzpav/crypto-broker/pkg/log"
)

const configPath = "./config.yaml"

func main() {
	log := log.NewZeroLog("crypto-broker", os.Getenv("VERSION"), os.Getenv("LOGGER_LEVEL"))

	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal().Sendf("Could not load configs: %v", err)
	}

	telegramClient := telegram.New(log, config)
	exchangeClient := exchange.New(log, config)
	commandHandler := command.New()
	broker := broker.New(log, commandHandler, telegramClient, exchangeClient)
	broker.Run()

}
