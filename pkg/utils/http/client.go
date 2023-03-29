package http

import (
	"time"

	"github.com/valyala/fasthttp"
)

type Client struct {
	req     *fasthttp.Request
	timeout int
}

func (c *Client) Init(url string) {
	c.req = fasthttp.AcquireRequest()
	c.req.SetRequestURI(url)
	c.SetTimeout(5)
}

func (c *Client) SetTimeout(timeout int) {
	c.timeout = timeout * int(time.Second)
}

func (c *Client) AddHeader(key, value string) {
	c.req.Header.Add(key, value)
}

func (c *Client) SetContentType(contentType string) {
	c.req.Header.SetContentType(contentType)
}

func (c *Client) SetAuthorization(authorization string) {
	c.req.Header.Set("Authorization", authorization)
}

func (c *Client) Post(body []byte) (response []byte, status int) {

	c.req.Header.SetMethod("POST")
	c.req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}

	err := client.DoTimeout(c.req, resp, time.Duration(c.timeout))

	if err != nil {
		return nil, resp.StatusCode()
	}

	return resp.Body(), resp.StatusCode()
}

func (c *Client) Get() (response []byte, status int) {

	c.req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}

	err := client.DoTimeout(c.req, resp, time.Duration(c.timeout))

	if err != nil {
		return nil, resp.StatusCode()
	}

	return resp.Body(), resp.StatusCode()
}

func (c *Client) Put(body []byte) (response []byte, status int) {

	c.req.Header.SetMethod("PUT")
	c.req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}

	err := client.DoTimeout(c.req, resp, time.Duration(c.timeout))

	if err != nil {
		return nil, resp.StatusCode()
	}

	return resp.Body(), resp.StatusCode()
}
