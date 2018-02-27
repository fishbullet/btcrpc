package btcrpc

//
// This is simple JSON RPC client for bitcoin node
//

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	"strings"
	"sync"
)

type Request struct {
	Version string        `json:"jsonrpc"`
	ID      uint32        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type Response struct {
	ID     int             `json:"id"`
	Error  json.RawMessage `json:"error"`
	Result json.RawMessage `json:"result"`
}

type Options struct {
	Login       string
	Password    string
	Host        string
	Port        string
	ContentType string
	TSL         bool
}

type Client struct {
	Client *http.Client
	ID     uint32
	Path   string
	*Options
	l sync.Mutex
}

type RpcResponse struct {
	StatusCode int
	Response
}

func NewClient(opt *Options) Client {
	c := Client{Client: &http.Client{}}
	c.Options = &Options{
		Login:       opt.Login,
		Password:    opt.Password,
		Host:        opt.Host,
		Port:        opt.Port,
		ContentType: "text/plain",
	}
	var scheme string
	if opt.TSL {
		scheme = "https"
	}
	scheme = "http"
	c.Path = fmt.Sprintf("%s://%s:%s", scheme, c.Options.Host, c.Options.Port)
	return c
}

func (c *Client) GetInfo() (*RpcResponse, error) {
	c.l.Lock()
	c.ID += 1
	defer c.l.Unlock()
	buf := c.buildReqParams("getinfo", []interface{}{})
	resp, err := c.submitRequest(buf)
	if err != nil {
		// log.Printf("%s", err)
		return &RpcResponse{}, err
	}
	return resp, nil
}

func (c *Client) Move(from, to string, amount float64) (*RpcResponse, error) {
	c.l.Lock()
	c.ID += 1
	defer c.l.Unlock()
	var params []interface{}
	params = append(params, from)
	params = append(params, to)
	params = append(params, amount)
	buf := c.buildReqParams("move", params)
	resp, err := c.submitRequest(buf)
	if err != nil {
		// log.Printf("%s", err)
		return &RpcResponse{}, err
	}
	return resp, nil
}

func (c *Client) GetNewAddress(account string) (*RpcResponse, error) {
	c.l.Lock()
	c.ID += 1
	defer c.l.Unlock()
	var params []interface{}
	params = append(params, account)
	buf := c.buildReqParams("getnewaddress", params)
	resp, err := c.submitRequest(buf)
	if err != nil {
		// log.Printf("%s", err)
		return &RpcResponse{}, err
	}
	return resp, nil
}

func (c *Client) GetBalance(account string, minConf int) (*RpcResponse, error) {
	c.l.Lock()
	c.ID += 1
	defer c.l.Unlock()
	var params []interface{}
	params = append(params, account)
	params = append(params, minConf)
	buf := c.buildReqParams("getbalance", params)
	resp, err := c.submitRequest(buf)
	if err != nil {
		// log.Printf("%s", err)
		return &RpcResponse{}, err
	}
	return resp, nil
}

func (c *Client) ValidateAddress(address string) (*RpcResponse, error) {
	c.l.Lock()
	c.ID += 1
	defer c.l.Unlock()
	var params []interface{}
	params = append(params, address)
	buf := c.buildReqParams("validateaddress", params)
	resp, err := c.submitRequest(buf)
	if err != nil {
		// log.Printf("%s", err)
		return &RpcResponse{}, err
	}
	return resp, nil
}

func (c *Client) EstimateFee(block int) (*RpcResponse, error) {
	c.l.Lock()
	c.ID += 1
	defer c.l.Unlock()
	var params []interface{}
	params = append(params, block)
	buf := c.buildReqParams("estimatesmartfee", params)
	resp, err := c.submitRequest(buf)
	if err != nil {
		// log.Printf("%s", err)
		return &RpcResponse{}, err
	}
	return resp, nil
}

func (c *Client) SendToAddress(address string, amount float64) (*RpcResponse, error) {
	c.l.Lock()
	c.ID += 1
	defer c.l.Unlock()
	var params []interface{}
	params = append(params, address)
	params = append(params, amount)
	buf := c.buildReqParams("sendtoaddress", params)
	resp, err := c.submitRequest(buf)
	if err != nil {
		// log.Printf("%s", err)
		return &RpcResponse{}, err
	}
	return resp, nil
}

func (client *Client) submitRequest(params []byte) (*RpcResponse, error) {
	req, err := http.NewRequest("POST", client.Path, strings.NewReader(string(params)))
	if client.Options.Login != "" && client.Options.Password != "" {
		creds := fmt.Sprintf("%s:%s", client.Options.Login, client.Options.Password)
		header := b64.StdEncoding.EncodeToString([]byte(creds))
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", header))
	}
	req.Header.Add("Content-Type", client.Options.ContentType)
	resp, err := client.Client.Do(req)

	if err != nil {
		return &RpcResponse{}, err
	}

	r := Response{}
	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err := parseJSON(data, &r); err != nil {
		return &RpcResponse{}, err
	}
	// Debug print to output
	// bitcoin node response
	// log.Printf("%+v", r)
	return &RpcResponse{StatusCode: resp.StatusCode, Response: r}, nil
}

func (client *Client) buildReqParams(method string, p []interface{}) []byte {
	jsonPayload, _ := json.Marshal(Request{"0.1", client.ID, method, p})
	return jsonPayload
}

func parseJSON(data []byte, dest interface{}) error {
	err := json.Unmarshal(data, dest)
	if err != nil {
		return err
	}
	return nil
}

func (response *RpcResponse) Result() json.RawMessage {
	return response.Response.Result
}

func (response *RpcResponse) ID() int {
	return response.Response.ID
}

func (response *RpcResponse) Error() json.RawMessage {
	return response.Response.Error
}
