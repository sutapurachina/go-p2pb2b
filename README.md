# Golang p2pb2b API library

## API Docs

p2pb2b API docs can be found [here](https://documenter.getpostman.com/view/6288660/SVYxnEmD?version=latest#b4c28f20-582e-4a71-ab13-ce2de8c9151d)

## Usage

```
package main

import (
	"fmt"
	"os"

	"github.com/krinklesaurus/go-p2pb2b"
)

func main() {
	// create client with api key and api secret
	client, err := p2pb2b.NewClient("API_KEY", "API_SECRET")
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
	fmt.Println(fmt.Sprintf("ticker %+v", ticker.Result))
}
```

## Testing

Tests are run with `make test`. It uses a Docker container to run a sticky Golang version. Coverage can be checked with running
`make test` first and then run `make cover`.

## Contributions

Contributions are welcome. Just open a PR and I will review.