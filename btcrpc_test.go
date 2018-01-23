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
	getNewAddress = `{
	"result": "mgzMtBQTFxi4v7gv3exg4MQTddHeDWPBVo",
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
)

func setupServer(payload string) *httptest.Server {
	l, _ := net.Listen("tcp", "127.0.0.1:8334")
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		Host:     "127.0.0.1",
		Port:     8334,
	})
}

func TestNewClient(t *testing.T) {
	assert := assert.New(t)
	client := setupClient()
	assert.NotNil(client)
}

func TestGetInfo(t *testing.T) {
	assert := assert.New(t)

	server := setupServer(getInfo)
	defer server.Close()
	client := setupClient()
	assert.NotNil(client)

	resp, err := client.GetInfo()
	assert.Nil(err)
	assert.Equal(1, resp.Response.ID)

	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestGetNewAddress(t *testing.T) {
	assert := assert.New(t)

	server := setupServer(getNewAddress)
	defer server.Close()
	client := setupClient()

	resp, err := client.GetNewAddress("test1")
	assert.Nil(err)

	result := []byte(resp.Response.Result)
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

	result := []byte(resp.Response.Result)
	var balance float64
	json.Unmarshal(result, &balance)

	assert.Equal(1.3, balance)
	assert.Equal(http.StatusOK, resp.StatusCode)
}
