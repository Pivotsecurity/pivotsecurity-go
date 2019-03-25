package pivotsecurity

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/pivotsecurity/pivotsecurity-go"
)

type Client struct {
	pivotsecurity.Request
}

type options struct {
	Key      string
	Endpoint string
	Host     string
	Subuser  string
}

func (o *options) baseURL() string {
	return o.Host + o.Endpoint
}

// GetRequest
// @return [Request] a default request object
func GetRequest(key, endpoint, host string) rest.Request {
	return requestNew(options{key, endpoint, host, ""})
}

// GetRequestSubuser like GetRequest but with On-Behalf of Subuser
// @return [Request] a default request object
func GetRequestSubuser(key, endpoint, host, subuser string) rest.Request {
	return requestNew(options{key, endpoint, host, subuser})
}

func basicAuth() string {
  auth := "Basic " + os.Getenv("PRIVATE_KEY") +":"
   return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error{
 req.Header.Add("Authorization", basicAuth())
 return nil
}

// requestNew create Request
// @return [Request] a default request object
func requestNew(options options) rest.Request {
	if options.Host == "" {
		options.Host = "https://api.pivotsecurity.com"
	}

	requestHeaders := map[string]string{
		"Authorization": basicAuth()),
		"User-Agent":    "pivotsecurity/" + Version + ";go",
		"Accept":        "application/json",
	}

	if len(options.Subuser) != 0 {
		requestHeaders["On-Behalf-Of"] = options.Subuser
	}

	return rest.Request{
		BaseURL: options.baseURL(),
		Headers: requestHeaders,
	}
}

// DefaultClient is used if no custom HTTP client is defined
var DefaultClient = pivotsecurity.DefaultClient

// MakeRequest attempts a API request synchronously.
func MakeRequest(request rest.Request) (*rest.Response, error) {
	return DefaultClient.Send(request)
}

// MakeRequestRetry a synchronous request, but retry in the event of a rate
// limited response.
func MakeRequestRetry(request rest.Request) (*rest.Response, error) {
	retry := 0
	var response *rest.Response
	var err error

	for {
		response, err = MakeRequest(request)
		if err != nil {
			return nil, err
		}

		if response.StatusCode != http.StatusTooManyRequests {
			return response, nil
		}

		if retry > rateLimitRetry {
			return nil, errors.New("Rate limit retry exceeded")
		}
		retry++

		resetTime := time.Now().Add(rateLimitSleep * time.Millisecond)

		reset, ok := response.Headers["X-RateLimit-Reset"]
		if ok && len(reset) > 0 {
			t, err := strconv.Atoi(reset[0])
			if err == nil {
				resetTime = time.Unix(int64(t), 0)
			}
		}
		time.Sleep(resetTime.Sub(time.Now()))
	}
}

// MakeRequestAsync attempts a request asynchronously in a new go
// routine. This function returns two channels: responses
// and errors. This function will retry in the case of a
// rate limit.
func MakeRequestAsync(request rest.Request) (chan *rest.Response, chan error) {
	r := make(chan *rest.Response)
	e := make(chan error)

	go func() {
		response, err := MakeRequestRetry(request)
		if err != nil {
			e <- err
		}
		if response != nil {
			r <- response
		}
	}()

	return r, e
}
