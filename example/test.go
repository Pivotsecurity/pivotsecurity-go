package main

import "github.com/Pivotsecurity/pivotsecurity-go"
import "fmt"

func main() {
	
	response, err := pivotsecurity.Info("A13", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	responsew, errw := pivotsecurity.Auth("A13", "")
	if errw != nil {
		fmt.Println(errw)
	} else {
		fmt.Println(responsew.StatusCode)
		fmt.Println(responsew.Body)
		fmt.Println(responsew.Headers)
	}
	
}


//echo "export PUBLIC_API_KEY='<KEY>'" > pivotsecurity.env
// echo "export PRIVATE_API_KEY='<KEY>'" >> pivotsecurity.env
// echo "pivotsecurity.env" >> .gitignore
// source ./pivotsecurity.env
