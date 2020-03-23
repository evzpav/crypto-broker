# Crypto Broker

It receives commands from Telegram and integrates to crypto exchanges create/cancel an order or to retrieve some information.
To get ticker from crypto exchanges, for example, a "/ticker" command will make a request to the exchange selected and retrieve the data and send it to Telegram.



#### Pre-requisites:
- [Make](https://www.gnu.org/software/make/)
- [Golang > 1.11](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/install/)

## Get Started

```bash
# Clone repository

# Create config.yaml based on the example
make config
# Create a bot on Telegram via @botfather and take note of the token
# Fill the Telegram bot token and exchanges api key and secret

# Run project locally (must have Go installed)
make run

# Or run it on Docker
make run-docker 
```


## Available commands (to be used on the Telegram bot):
```bash
  # Set the exchange to be used for all the commands (default is bitfinex.com), which show be the same name as specified on config.yaml
  /exchange <exchangeName>

  # Ticker: pair param should be in the format ETH/USD (default is BTC/USD)
  /ticker <cryptopair>

  # Account info (needs api key and secret on config.yaml)
  /balance
```
## Available exchanges
This is easily expandable due to use of lib [goex](github.com/nntaoli-project/goex) which supports multiple exchanges
- bitfinex.com
- binance.com