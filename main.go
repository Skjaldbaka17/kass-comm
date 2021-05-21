package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var authToken string = "kass_test_auth_token"

func main() {

	//Encode the data
	postBody, _ := json.Marshal(&Request{
		Name:  "Toby",
		Email: "Toby@example.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
