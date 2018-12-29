# cmc
a CLI tool for pulling coinmarketcap api data

## Install
```go get github.com/apaz037/cmc && cd $GOROOT/src/github.com/apaz037/cmc && go install```

## Current functionality

GET:

```cmc get TKR```

returns the current price for a ticker defaulting to USD.

## Planned functionality
```cmc get TKR1 TKR2 TKR3```

will return a list of current prices for specified tickers

```cmc get TKRSET.csv```

will return a list of current prices for tickers specified in a .csv

```cmc get 10```

will return a list of current prices for the top 10 cryptocurrencies by market cap

```cmc get --exchange= --convert=```

the exchange flag will set the exchange to pull price from

the convert flag will set the final currency to display prices in (USD, BTC, EUR, etc)

these flags will be compatible with all of the above commands
