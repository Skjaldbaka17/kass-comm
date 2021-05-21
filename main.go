package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Example json
// {
//     "amount": 2199,
//     "description": "Kass bolur",
//     "image_url": "https://photos.kassapi.is/kass/kass-bolur.jpg",
//     "order": "ABC123",
//     "recipient": "7728440",
//     "terminal": 1,
//     "expires_in": 90,
//     "notify_url": "https://example.com/callbacks/kass"
// }

//Request object
type Request struct {
	Amount      int    `json:"amount"`
	Description string `json: "Description"`
	Image_Url   string `json: "image_url"`
	Order       string `json: "order"`
	Recipient   string `json: "recipient"`
	Terminal    int    `json: "terminal"`
	Expires_In  int    `json: "expires_in"`
	Notify_Url  string `json: "notify_url"`
}

var authToken string = "kass_test_auth_token"

func main() {

	//Encode the data
	postBody, _ := json.Marshal(&Request{
		Amount:      2199,
		Description: "Kass bolur",
		Image_Url:   "https://photos.kassapi.is/kass/kass-bolur.jpg",
		Order:       "ABC123",
		Recipient:   "7728440",
		Terminal:    1,
		Expires_In:  90,
		Notify_Url:  "https://example.com/callbacks/kass",
	})
	client := &http.Client{}
	requestBody := bytes.NewBuffer(postBody)
	username := authToken
	req, err := http.NewRequest("POST", "https://api.kass.is/v1/payments", requestBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return
	}

	req.SetBasicAuth(username, "")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf("HEllo!")
	log.Printf(sb)
	//Leverage Go's HTTP Post function to make request
	// resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	// //Handle Error
	// if err != nil {
	// 	log.Fatalf("An Error Occured %v", err)
	// }
	// defer resp.Body.Close()
	//Read the response body
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// sb := string(body)
	// log.Printf(sb)
}
