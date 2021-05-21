package kasscomm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var auth_token string
var is_prod = false //Defaults to false, i.e. sandbox/test-env

const real_path, test_path string = "https://api.kass.is/v1/", "https://api.testing.kass.is/v1/"

var base_url string = test_path

//Request is structured according to the json that the kass-api expects
type Request struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Image_Url   string `json:"image_url"`
	Order       string `json:"order"`
	Recipient   string `json:"recipient"`
	Terminal    int    `json:"terminal"`
	Expires_In  int    `json:"expires_in"`
	Notify_Url  string `json:"notify_url"`
}

//Response is structured according to the expected json received from the kass-api
type Response struct {
	Success bool   `json:"success"`
	Id      string `json:"id,omitempty"`
	Created int    `json:"created,omitempty"`
	Error   Error  `json:"error,omitempty"`
}

//Error is structured according to the expected json received from the kass-api on a success==false response
type Error struct {
	Code    string `json:"code"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

// SetAuthToken takes in the authToken-string to be used in the api-calls
func SetAuthToken(token string) {
	auth_token = token
}

// SetDev sets the environment to dev, i.e. use sandbox
func SetDev() {
	is_prod = false
	base_url = test_path
}

// SetProd sets the environment to production, i.e. use the real api
func SetProd() {
	is_prod = true
	base_url = real_path
}

func GetAuthToken() string {
	return auth_token
}

func GetIsProd() bool {
	return is_prod
}

/*
ValidateInputRequest does minor validation on the request object to be sent to the api. It returns an error if the request
object supplied fails basic validation. One Could also just use error codes from the Api if performance is not an issue.
If the basic validation finds nothing odd with the request it returns nil.
*/
func validateInputRequest(request *Request) error {
	if *request == (Request{}) {
		return errors.New("request struct is empty")
	}

	if request.Amount <= 0 {
		return errors.New("request.Amount is an invalid value")
	}

	if request.Recipient == "" {
		return errors.New("request.Recipient is empty")
	}

	return nil
}

/*
InitiatePayment initiates a payment request through the kass-api and returns the Reponse. If something fails
it returns an empty Response and an error.
*/
func InitiatePayment(request *Request) (Response, error) {
	//Validate before making the request

	if auth_token == "" {
		return Response{}, errors.New("authToken can not be an empty string")
	}
	err := validateInputRequest(request)
	if err != nil {
		return Response{}, err
	}

	//Encode the data
	post_body, _ := json.Marshal(request)

	//Create the request
	client := &http.Client{}
	request_body := bytes.NewBuffer(post_body)
	req, err := http.NewRequest("POST", base_url+"payments", request_body)

	if err != nil {
		return Response{}, err
	}

	//Set the authentication token and call the api
	req.SetBasicAuth(auth_token, "")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return Response{}, err
	}

	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return Response{}, err
	}

	// Decode json into Response struct
	var json_response Response
	err = json.Unmarshal(body, &json_response)

	if err != nil {
		log.Fatalln(err)
		return Response{}, err
	}

	return json_response, nil
}
