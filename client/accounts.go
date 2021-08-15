package accounts

import (
	"bytes"
	"encoding/json"
	// "io/ioutil"
	// "log"
	"net/http"

	"../model"
)

const (
	URL = "http://localhost:8080/v1/organisation/accounts"
)

// func CreatePayload()

// POST /v1/organisation/accounts
func Create(data model.AccountData) (*http.Response, error) {
	body, _ := json.Marshal(data)
	reqBody := bytes.NewReader(body)

	resp, err := http.Post(URL,"application/json", reqBody)
	// respBody, err := ioutil.ReadAll(resp.Body)
	// respCode := string(resp.StatusCode)

	return resp, err
}

// // GET /v1/organisation/accounts/{account_id}
// func Fetch(fiat string, crypto string) (string, error) {
// 	//Build The URL string

// 	//We make HTTP request using the Get function
// 	resp, err := http.Get(URL)
// 	if err != nil {
// 		log.Fatal("ooopsss an error occurred, please try again")
// 	}
// 	defer resp.Body.Close()
// 	//Create a variable of the same type as our model
// 	var cResp model.Cryptoresponse
// 	//Decode the data
// 	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
// 		log.Fatal("ooopsss! an error occurred, please try again")
// 	}
// 	//Invoke the text output function & return it with nil as the error value
// 	return cResp.TextOutput(), nil
// }

// // DELETE /v1/organisation/accounts/{account_id}?version={version}
// func Delete() {}
