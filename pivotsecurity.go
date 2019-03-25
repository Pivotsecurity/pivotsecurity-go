package pivotsecurity

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Method string

const (
	Get    Method = "GET"
	Post   Method = "POST"
	Put    Method = "PUT"
	Patch  Method = "PATCH"
	Delete Method = "DELETE"
)

type Request struct {
	Method      Method
	BaseURL     string // e.g. https://api.pivotsecurity.com
	Headers     map[string]string
	QueryParams map[string]string
	Body        []byte
}

type RestError struct {
	Response *Response
}

func (e *RestError) Error() string {
	return e.Response.Body
}

var DefaultClient = &Client{HTTPClient: http.DefaultClient}

type Client struct {
	HTTPClient *http.Client
}

type Response struct {
	StatusCode int                 // e.g. 200
	Body       string              // e.g. {"result: success"}
	Headers    map[string][]string // e.g. map[X-Ratelimit-Limit:[600]]
}

func AddQueryParameters(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}

func BuildRequestObject(request Request) (*http.Request, error) {
	if len(request.QueryParams) != 0 {
		request.BaseURL = AddQueryParameters(request.BaseURL, request.QueryParams)
	}
	req, err := http.NewRequest(string(request.Method), request.BaseURL, bytes.NewBuffer(request.Body))
	if err != nil {
		return req, err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	_, exists := req.Header["Content-Type"]
	if len(request.Body) > 0 && !exists {
		req.Header.Set("Content-Type", "application/json")
	}
	req.method := "POST"

	return req, err
}

func MakeRequest(req *http.Request) (*http.Response, error) {
	return DefaultClient.HTTPClient.Do(req)
}

func BuildResponse(res *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(res.Body)
	response := Response{
		StatusCode: res.StatusCode,
		Body:       string(body),
		Headers:    res.Header,
	}
	res.Body.Close() // nolint
	return &response, err
}

func API(request Request) (*Response, error) {
	return Send(request)
}

func Send(request Request) (*Response, error) {
	return DefaultClient.Send(request)
}

func (c *Client) MakeRequest(req *http.Request) (*http.Response, error) {
	return c.HTTPClient.Do(req)
}

func (c *Client) API(request Request) (*Response, error) {
	return c.Send(request)
}

func (c *Client) Send(request Request) (*Response, error) {
	req, err := BuildRequestObject(request)
	if err != nil {
		return nil, err
	}

	// Build the HTTP client and make the request.
	res, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	// Build Response object.
	return BuildResponse(res)
}

