package btcrpc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestBtcRpcNewDefaultClient(t *testing.T) {
	l, e := net.Listen("tcp", ":8334")
	if e != nil {
		fmt.Printf("net.Listen tcp :0: %v", e)
	}
	defer l.Close()
	assert := assert.New(t)
	client, err := NewDefaultClient()
	assert.Nil(err)
	assert.NotNil(client)
}

func TestBtcRpcNewClient(t *testing.T) {
	l, e := net.Listen("tcp", ":8334")
	if e != nil {
		fmt.Printf("net.Listen tcp :0: %v", e)
	}
	defer l.Close()
	assert := assert.New(t)
	client, err := NewClient(&Options{"127.0.0.1", 8334})
	assert.Nil(err)
	assert.NotNil(client)
}
