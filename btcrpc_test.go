package btcrpc

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	// "log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	sendToAddress = `{
	"result": ["ef2ca42bcbf7b877dea40ebd8124360903183207340d62ae1a93dbb87896628c"],
	"error": null,
	"id": 1
}`
	getNewAddress = `{
	"result": "mgzMtBQTFxi4v7gv3exg4MQTddHeDWPBVo",
	"error": null,
	"id": 1
}`

	estimateFee = `{
	"result": {
	  "feerate": 0.001,
	  "blocks": 6
	},
	"error": null,
	"id": 1
}`

	validateAddress = `{
  "result": {
	  "isvalid": true,
	  "address": "moj9RtYj4Anwhzgkwxs1Ldaw3TPxXLSDig",
	  "scriptPubKey": "76a9145a0f4afce3a99168131114f12430e03a30983bd388ac",
	  "ismine": false,
	  "iswatchonly": false,
	  "isscript": false
	},
	"error": null,
	"id": 1
}`

	getBalance = `{
  "result": 1.30000000,
  "error": null,
  "id": 1
}`
	getInfo = `{
  "result": {
    "deprecation-warning": "WARNING: getinfo is deprecated and will be fully removed in 0.16. Projects should transition to using getblockchaininfo, getnetworkinfo, and getwalletinfo before upgrading to 0.16",
    "version": 150100,
    "protocolversion": 70015,
    "walletversion": 139900,
    "balance": 0,
    "blocks": 1259883,
    "timeoffset": 44,
    "connections": 8,
    "proxy": "",
    "difficulty": 1,
    "testnet": true,
    "keypoololdest": 1516618949,
    "keypoolsize": 2000,
    "paytxfee": 0,
    "relayfee": 1e-05,
    "errors": "Warning: unknown new rules activated (versionbit 28)"
  },
  "error": null,
  "id": 1
}`
	host    = "127.0.0.1"
	port    = "8334"
	address = fmt.Sprintf("%s:%s", host, port)
)

func setupServer(payload string) *httptest.Server {
	l, _ := net.Listen("tcp", address)
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			// Authorization header
			// missing basic auth
			w.WriteHeader(http.StatusUnauthorized)
		}
		fmt.Fprintln(w, payload)
	}))
	ts.Listener.Close()
	ts.Listener = l
	ts.Start()
	return ts
}

func setupClient() Client {
	return NewClient(&Options{
		Login:    "admin",
		Password: "admin",
		Host:     host,
		Port:     port,
	})
}

func TestNewClient(t *testing.T) {
	assert := assert.New(t)
	client := setupClient()
	assert.NotNil(client)
}

func TestClientCredentials(t *testing.T) {
	assert := assert.New(t)
	client := NewClient(&Options{
		Login:    "",
		Password: "",
		Host:     host,
		Port:     port,
	})
	assert.NotNil(client)
	server := setupServer("{}")
	defer server.Close()
	resp, err := client.GetInfo()
	assert.Nil(err)
	assert.Equal(http.StatusUnauthorized, resp.StatusCode)
}

func TestGetInfo(t *testing.T) {
	assert := assert.New(t)

	server := setupServer(getInfo)
	defer server.Close()
	client := setupClient()
	assert.NotNil(client)

	resp, err := client.GetInfo()
	assert.Nil(err)
	assert.Equal(1, resp.ID())

	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestGetNewAddress(t *testing.T) {
	assert := assert.New(t)

	server := setupServer(getNewAddress)
	defer server.Close()
	client := setupClient()

	resp, err := client.GetNewAddress("test1")
	assert.Nil(err)

	result := []byte(resp.Result())
	var address string
	json.Unmarshal(result, &address)

	assert.Equal("mgzMtBQTFxi4v7gv3exg4MQTddHeDWPBVo", address)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestGetBalance(t *testing.T) {
	assert := assert.New(t)

	server := setupServer(getBalance)
	defer server.Close()
	client := setupClient()

	resp, err := client.GetBalance("test1", 1)
	assert.Nil(err)

	result := []byte(resp.Result())
	var balance float64
	json.Unmarshal(result, &balance)

	assert.Equal(1.3, balance)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestValidateAddress(t *testing.T) {
	assert := assert.New(t)

	server := setupServer(validateAddress)
	defer server.Close()
	client := setupClient()

	resp, err := client.ValidateAddress("mgzMtBQTFxi4v7gv3exg4MQTddHeDWPBVo")
	assert.Nil(err)

	result := []byte(resp.Result())
	var adder map[string]interface{}
	json.Unmarshal(result, &adder)

	assert.Equal(true, adder["isvalid"])
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestEstimateFee(t *testing.T) {
	assert := assert.New(t)

	server := setupServer(estimateFee)
	defer server.Close()
	client := setupClient()

	resp, err := client.EstimateFee(6)
	assert.Nil(err)

	result := []byte(resp.Result())
	var adder map[string]interface{}
	json.Unmarshal(result, &adder)

	assert.Equal(0.001, adder["feerate"])
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestSendToAddress(t *testing.T) {
	assert := assert.New(t)

	server := setupServer(sendToAddress)
	defer server.Close()
	client := setupClient()

	resp, err := client.SendToAddress("mgzMtBQTFxi4v7gv3exg4MQTddHeDWPBVo", 0.01)
	assert.Nil(err)

	result := []byte(resp.Result())
	var tx []interface{}
	json.Unmarshal(result, &tx)

	assert.Equal(1, len(tx))
	assert.Equal(http.StatusOK, resp.StatusCode)
}
