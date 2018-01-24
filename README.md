# Bitcoin Core - RPC client

[![Build Status](https://travis-ci.org/fishbullet/btcrpc.svg?branch=master)](https://travis-ci.org/fishbullet/btcrpc)

## BTCRPC

The `btcrpc` is the http Bitcoin JSON-RPC client package written in Go.
That package was written for usage in internal projects.
Package doesn't contains [all RPC methods](https://bitcoin.org/en/developer-reference#rpcs). 

The implemented methods:

- [x] [getinfo](https://bitcoin.org/en/developer-reference#getinfo)
- [x] [getbalance](https://bitcoin.org/en/developer-reference#getbalance)
- [x] [getnewaddress](https://bitcoin.org/en/developer-reference#getnewaddress)
- [ ] [estimatefee](https://bitcoin.org/en/developer-reference#estimatefee)
- [ ] [sendtoaddress](https://bitcoin.org/en/developer-reference#sendtoaddress)
- [ ] [getreceivedbyaddress](https://bitcoin.org/en/developer-reference#getreceivedbyaddress)

Feel free to submit a pull request with new RPC method.

## Usage

Make sure you're running bitcoin node. Setup and run JSON-RPC client:

```go
package main

import (
	"encoding/json"
	"github.com/fishbullet/btcrpc"
	"log"
)

func main() {
	btcClient := btcrpc.NewClient(&btcrpc.Options{
		Login:    "admin",
		Password: "admin",
		Host:     "127.0.0.1", // Localhost
		Port:     8334,        // Testnet port
		TSL:      false,       // If you're using https instead of http
	})

	// Get balance across all accounts
	minConf := 0
	account := ""
	balance, err := btcClient.GetBalance(account, minConf)
	if err != nil {
		log.Fatalf("%s", err)
	}
	result := []byte(balance.Result())
	var b float64
	json.Unmarshal(result, &b)
	log.Printf("%f", b)
}
```
Will output:
```
2018/01/24 16:15:25 0.000000
```


### Disclaimer
> :exclamation: Package provided “as is," and you use it at your own risk.. 

### Buy me a beer

BTC: 19SYMA2hqRZHRSL4di35Uf7jV87KBKc9bf

ETH: 0x9c00577856BcBDf87Fea58404FaEC5A2034BD86F
