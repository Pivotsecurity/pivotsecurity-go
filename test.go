package main

import "github.com/pivtosecurity/pivtosecurity-go"
import "fmt"

func main() {
	const host = "http://localhost:8080/api"
	endpoint := "/account/info"
	baseURL := host + endpoint
	Headers := make(map[string]string)
	key := os.Getenv("API_KEY")
	Headers["Authorization"] = "Basic " + os.Getenv("PRIVATE_KEY") +":"
	Headers["X-Test"] = "Test"
	var Body = []byte(`{"uid": "A13", "emial":""}`)
	queryParams := make(map[string]string)
	method := rest.Post
	request = rest.Request{
		Method:      method,
		BaseURL:     baseURL,
		Headers:     Headers,
		QueryParams: queryParams,
		Body:        Body,
	}
	response, err := rest.Send(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}