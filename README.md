# Golang P2pb2b API library

## GET Markets

```
GetMarkets() (*MarketsResult, error)
```

## Usage


```

import (
    p2pb2b "gitlab.com/krinklesaurus/go-p2pb2b"
)

// create client with api key and api secret
client, err := p2pb2b.NewClient([API_KEY], [API_SECRET])
if err != nil {
	fmt.Println(fmt.Sprintf("damn, %v", err))
    os.Exit(1)
}

// get ticker
ticker, err := client.GetTicker("ETH_BTC")
if err != nil {
	fmt.Println(fmt.Sprintf("damn, %v", err))
    os.Exit(1)
}
fmt.Println(fmt.Sprintf("ticker %+v", ticker.Ticker))
```