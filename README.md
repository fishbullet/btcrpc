# Bitcoin Core - RPC client

[![Build Status](https://travis-ci.org/fishbullet/btcrpc.svg?branch=master)](https://travis-ci.org/fishbullet/btcrpc)

## BTCRPC

The `btcrpc` library provide simple client for interaction with [bitcoin](https://github.com/bitcoin/bitcoin).

## Usage

Make sure you're running bitcoin node. Setup and run JSON-RPC client:

```go
package main

import (  
    "encoding/json"
    "github.com/fishbullet/btcrpc"
    "fmt"
)

func main() {  
  btcClient := btcrpc.NewClient(&btcrpc.Options{
    Login:    "RPC_LOGIN_HERE",
    Password: "RPC_PASSWORD_HERE",
    Host:     "127.0.0.1", // Localhost
    Port:     8334,        // Testnet port
    TSL:      true,        // If you're using https instead of http
  })

  // Get balance across all accounts
  balance, err := btcClient.GetBalance("", 0)
  result := []byte(balance.Response.Result)
  var balance float64
  json.Unmarshal(result, &balance)
  fmt.Printf("%s", balance) // => 0.034
}
```

## Development

Run docker container with testnet:

```bash
docker build -t btc_node
docker run --rm -v -p 8334:8334 $(pwd)/bitcoin:/root/.bitcoin btc_node
```

Test bitcoin node RPC api:

```bash
curl --data-binary '{"jsonrpc":"1.0","id":1,"method":"getinfo","params":[]}' -H 'content-type:text/plain;' http://admin:admin@127.0.0.1:8334/
```
Should return:

```json
{
  "result": {
    "deprecation-warning": "WARNING: getinfo is deprecated and will be fully removed in 0.16...",
    "version": 150100,
    "protocolversion": 70015,
    "walletversion": 139900,
    "balance": 0,
    "blocks": 90381,
    "timeoffset": 45,
    "connections": 8,
    "proxy": "",
    "difficulty": 16,
    "testnet": true,
    "keypoololdest": 1516618949,
    "keypoolsize": 2000,
    "paytxfee": 0,
    "relayfee": 1e-05,
    "errors": ""
  },
  "error": null,
  "id": 1
}
```

### Buy me a beer

BTC: 19SYMA2hqRZHRSL4di35Uf7jV87KBKc9bf

ETH: 0x9c00577856BcBDf87Fea58404FaEC5A2034BD86F
