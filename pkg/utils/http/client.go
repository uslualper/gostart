package http

import (
	"time"

	"github.com/valyala/fasthttp"
)

type Client struct {
	req     *fasthttp.Request
	timeout time.Duration
}

func (c *Client) Init(url string) {
	c.req = fasthttp.AcquireRequest()
	c.req.SetRequestURI(url)
	c.SetTimeout(5)
}

func (c *Client) SetTimeout(timeout int) {
	c.timeout = time.Duration(timeout) * time.Second
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

func (c *Client) executeRequest(method string, body []byte) (response []byte, status int) {
	c.req.Header.SetMethod(method)
	c.req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{}
	err := client.DoTimeout(c.req, resp, c.timeout)

	if err != nil {
		return nil, fasthttp.StatusInternalServerError
	}

	return resp.Body(), resp.StatusCode()
}

func (c *Client) Post(body []byte) (response []byte, status int) {
	return c.executeRequest("POST", body)
}

func (c *Client) Get() (response []byte, status int) {
	return c.executeRequest("GET", nil)
}

func (c *Client) Put(body []byte) (response []byte, status int) {
	return c.executeRequest("PUT", body)
}
