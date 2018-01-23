# Bitcoin Core - RPC client

[![Build Status](https://travis-ci.org/fishbullet/btcrpc.svg?branch=master)](https://travis-ci.org/fishbullet/btcrpc)

## BTCRPC

The `btcrpc` is the http Bitcoin JSON-RPC client package written in Go.
That package was written for usage in internal projects.
Package doesn't contains [all RPC methods](https://bitcoin.org/en/developer-reference#rpcs). 

The implemented methods:

- [x] getinfo
- [x] getbalance
- [x] getnewaddress
- [ ] ...

Feel free to submit a pull request with new RPC method.

### Disclaimer
> :exclamation: Package provided “as is," and you use it at your own risk.. 

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

### Commands

addMultiSigAddress: 'addmultisigaddress',
addNode: 'addnode', // bitcoind v0.8.0+
backupWallet: 'backupwallet',
createMultiSig: 'createmultisig',
createRawTransaction: 'createrawtransaction', // bitcoind v0.7.0+
decodeRawTransaction: 'decoderawtransaction', // bitcoind v0.7.0+
decodeScript: 'decodescript',
dumpPrivKey: 'dumpprivkey',
dumpWallet: 'dumpwallet', // bitcoind v0.9.0+
encryptWallet: 'encryptwallet',
estimateFee: 'estimatefee', // bitcoind v0.10.0x
estimatePriority: 'estimatepriority', // bitcoind v0.10.0+
generate: 'generate', // bitcoind v0.11.0+
getAccount: 'getaccount',
getAccountAddress: 'getaccountaddress',
getAddedNodeInfo: 'getaddednodeinfo', // bitcoind v0.8.0+
getAddressesByAccount: 'getaddressesbyaccount',
getBalance: 'getbalance',
getBestBlockHash: 'getbestblockhash', // bitcoind v0.9.0+
getBlock: 'getblock',
getBlockchainInfo: 'getblockchaininfo', // bitcoind v0.9.2+
getBlockCount: 'getblockcount',
getBlockHash: 'getblockhash',
getBlockTemplate: 'getblocktemplate', // bitcoind v0.7.0+
getChainTips: 'getchaintips', // bitcoind v0.10.0+
getConnectionCount: 'getconnectioncount',
getDifficulty: 'getdifficulty',
getGenerate: 'getgenerate',
getInfo: 'getinfo',
getMempoolInfo: 'getmempoolinfo', // bitcoind v0.10+
getMiningInfo: 'getmininginfo',
getNetTotals: 'getnettotals',
getNetworkInfo: 'getnetworkinfo', // bitcoind v0.9.2+
getNetworkHashPs: 'getnetworkhashps', // bitcoind v0.9.0+
getNewAddress: 'getnewaddress',
getPeerInfo: 'getpeerinfo', // bitcoind v0.7.0+
getRawChangeAddress: 'getrawchangeaddress', // bitcoin v0.9+
getRawMemPool: 'getrawmempool', // bitcoind v0.7.0+
getRawTransaction: 'getrawtransaction', // bitcoind v0.7.0+
getReceivedByAccount: 'getreceivedbyaccount',
getReceivedByAddress: 'getreceivedbyaddress',
getTransaction: 'gettransaction',
getTxOut: 'gettxout', // bitcoind v0.7.0+
getTxOutProof: 'gettxoutproof', // bitcoind v0.11.0+
getTxOutSetInfo: 'gettxoutsetinfo', // bitcoind v0.7.0+
getUnconfirmedBalance: 'getunconfirmedbalance', // bitcoind v0.9.0+
getWalletInfo: 'getwalletinfo', // bitcoind v0.9.2+
help: 'help',
importAddress: 'importaddress', // bitcoind v0.10.0+
importPrivKey: 'importprivkey',
importWallet: 'importwallet', // bitcoind v0.9.0+
keypoolRefill: 'keypoolrefill',
keyPoolRefill: 'keypoolrefill',
listAccounts: 'listaccounts',
listAddressGroupings: 'listaddressgroupings', // bitcoind v0.7.0+
listLockUnspent: 'listlockunspent', // bitcoind v0.8.0+
listReceivedByAccount: 'listreceivedbyaccount',
listReceivedByAddress: 'listreceivedbyaddress',
listSinceBlock: 'listsinceblock',
listTransactions: 'listtransactions',
listUnspent: 'listunspent', // bitcoind v0.7.0+
lockUnspent: 'lockunspent', // bitcoind v0.8.0+
move: 'move',
ping: 'ping', // bitcoind v0.9.0+
prioritiseTransaction: 'prioritisetransaction', // bitcoind v0.10.0+
sendFrom: 'sendfrom',
sendMany: 'sendmany',
sendRawTransaction: 'sendrawtransaction', // bitcoind v0.7.0+
sendToAddress: 'sendtoaddress',
setAccount: 'setaccount',
setGenerate: 'setgenerate',
setTxFee: 'settxfee',
signMessage: 'signmessage',
signRawTransaction: 'signrawtransaction', // bitcoind v0.7.0+
stop: 'stop',
submitBlock: 'submitblock', // bitcoind v0.7.0+
validateAddress: 'validateaddress',
verifyChain: 'verifychain', // bitcoind v0.9.0+
verifyMessage: 'verifymessage',
verifyTxOutProof: 'verifytxoutproof', // bitcoind v0.11.0+
walletLock: 'walletlock',
walletPassphrase: 'walletpassphrase',
walletPassphraseChange: 'walletpassphrasechange'

### Buy me a beer

BTC: 19SYMA2hqRZHRSL4di35Uf7jV87KBKc9bf

ETH: 0x9c00577856BcBDf87Fea58404FaEC5A2034BD86F
