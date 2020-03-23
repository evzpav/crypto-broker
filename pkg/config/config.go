package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type TelegramConfig struct {
	Token   string `yaml:"token"`
	Timeout int    `yaml:"timeout"`
	Debug   bool   `yaml:"debug"`
}

type ExchangeConfig struct {
	ExchangeName string `yaml:"exchange"`
	PublicKey    string `yaml:"public_key"`
	SecretKey    string `yaml:"secret_key"`
}

type Config struct {
	Telegram  TelegramConfig   `yaml:"telegram"`
	Exchanges []ExchangeConfig `yaml:"exchanges"`
}

func NewConfig(configFilePath string) (*Config, error) {
	var configs Config
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	contentToMarshal, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(contentToMarshal, &configs)
	if err != nil {
		return nil, err
	}
	return &configs, nil
}
