package pivotsecurity

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"encoding/base64"
)

const host = "https://api.pivotsecurity.com/api/"

const OP_CREATE= host + "account/create";
const OP_INFO = host + "account/info";
const OP_RISK_SCORE = host + "account/riskscore";
const OP_UPDATE_RISK_SCORE = host + "account/updateriskscore";
const OP_QRCODE = host + "account/qrcode";
const OP_AUTH_CODE = host + "account/authcode";
const OP_LOGS = host + "account/logs";
const OP_LOCK = host + "account/lock";
const OP_UNLOCK = host + "account/unlock";
const OP_TRAIN_ML = host + "account/trainml";
const OP_TEST_ML = host + "account/testml";
const OP_AUTH_META = host + "account/authwithmetadata";
const OP_SEND_AUTH_META = host + "account/sendauthwithmetadata";
const OP_VERIFY_META = host + "account/verifywithmetadata";
const OP_VERIFY_SESSION = host + "account/verifysession";

const OP_CUST_CREATE = host + "customer/create";
const OP_AUTH = host + "customer/auth";
const OP_VALIDATE = host + "customer/verify";

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
	request.Method = "POST"
	//if len(request.QueryParams) != 0 { // not used
	//	request.BaseURL = AddQueryParameters(request.BaseURL, request.QueryParams)
	//}
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

	res, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	return BuildResponse(res)
}

func privateAuth() string {
  auth := "Basic " + os.Getenv("PRIVATE_API_KEY") +":"
   return base64.StdEncoding.EncodeToString([]byte(auth))
}
func publicAuth() string {
  auth := "Basic " + os.Getenv("PUBLIC_API_KEY") +":"
   return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Account and Customer methods
func Create(uid string, email string, channel string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`", "channel":"` + channel + `"}`)
	
	request := Request{
		BaseURL:     OP_CREATE,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func Info(uid string, email string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`"}`)
	
	request := Request{
		BaseURL:     OP_INFO,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func Riskscore(uid string, email string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`"}`)
	
	request := Request{
		BaseURL:     OP_RISK_SCORE,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func UpdateRiskscore(uid string, email string, riskscore string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`", "riskscore":"` + riskscore + `"}`)
	
	request := Request{
		BaseURL:     OP_UPDATE_RISK_SCORE,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func QRCode(uid string, email string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`"}`)
	
	request := Request{
		BaseURL:     OP_QRCODE,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func AuthCode(uid string, email string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`"}`)
	
	request := Request{
		BaseURL:     OP_AUTH_CODE,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func Logs(uid string, email string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`"}`)
	
	request := Request{
		BaseURL:     OP_LOGS,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func Lock(uid string, email string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`"}`)
	
	request := Request{
		BaseURL:     OP_LOCK,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func Unlock(uid string, email string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`"}`)
	
	request := Request{
		BaseURL:     OP_UNLOCK,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func TrainMl(uid string, email string, data string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`", "data":"` + data + `"}`)
	
	request := Request{
		BaseURL:     OP_TRAIN_ML,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func TestMl(uid string, email string, data string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`", "data":"` + data + `"}`)
	
	request := Request{
		BaseURL:     OP_TEST_ML,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func AuthWithMetadata(uid string, email string, metadata string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`", "metadata":"` + metadata + `"}`)
	
	request := Request{
		BaseURL:     OP_AUTH_META,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func SendAuthWithMetadata(uid string, email string, metadata string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`", "metadata":"` + metadata + `"}`)
	
	request := Request{
		BaseURL:     OP_SEND_AUTH_META,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func VerifyWithMetadata(uid string, email string, code string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`", "code":"` + code + `"}`)
	
	request := Request{
		BaseURL:     OP_VERIFY_META,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func VerifySession(uid string, email string, sessionid string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PRIVATE_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`", "sessionid":"` + sessionid + `"}`)
	
	request := Request{
		BaseURL:     OP_VERIFY_SESSION,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func CustomerCreate(uid string, email string, channel string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PUBLIC_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`", "channel":"` + channel + `"}`)
	
	request := Request{
		BaseURL:     OP_CUST_CREATE,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func Auth(uid string, email string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PUBLIC_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`"}`)
	
	request := Request{
		BaseURL:     OP_AUTH,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}
func Validate(uid string, email string, code string) (*Response, error) {
	Headers := make(map[string]string)
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PUBLIC_API_KEY") +":"))
	Headers["Authorization"] = "Basic " + auth
	var Body = []byte(`{"uid":"`+ uid +`", "email":"` + email +`", "code":"` + code + `"}`)
	
	request := Request{
		BaseURL:     OP_VALIDATE,
		Headers:     Headers,
		Body:        Body,
	}
	
	return Send(request)
}

