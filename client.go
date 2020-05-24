package jsonapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type URLs struct {
	BaseURL string
	Port    string
	Path    string
}

func main() {
	var operation int
	fmt.Println(" ************** Welcome API Client ************** ")
	fmt.Println("Please select an Operation")
	fmt.Println("1 - Get All Posts")
	fmt.Println("2 - Get specific post by ID")
	fmt.Println("3 - Insert a post")
	fmt.Println("4 - Update post")
	fmt.Println("5 - Delete a post")
	_, err := fmt.Scanf("%d", &operation)
	if err != nil {
		fmt.Println("Please enter a valid number!")
	}

	switch operation {
	case 1:
		var def = URLs{
			BaseURL: "http://192.168.1.106",
			Port:    "8080",
			Path:    "posts",
		}
		req, err := http.NewRequest(
			http.MethodGet,
			def.BaseURL+":"+def.Port+"/"+def.Path,
			nil,
		)
		if err != nil {
			log.Fatalf("error creating HTTP request: %v", err)
		}
		req.Header.Add("Accept", "application/json")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalf("error sending HTTP request: %v", err)
		}
		responseBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("error reading HTTP response body: %v", err)
		}
		log.Println("We got the response:", string(responseBytes))
	}

}
