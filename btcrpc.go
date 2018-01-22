package btcrpc

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Options struct {
	Host string
	Port int
}

type Client struct {
	client *rpc.Client
}

func NewDefaultClient() (*Client, error) {
	client, err := NewClient(&Options{"localhost", 8334})
	if err != nil {
		return &Client{}, err
	}
	return client, nil
}

func NewClient(opt *Options) (*Client, error) {
	client, err := jsonrpc.Dial("tcp", fmt.Sprintf("%s:%d", opt.Host, opt.Port))
	if err != nil {
		fmt.Printf("%s", err)
		return &Client{}, err
	}
	return &Client{client}, nil
}
