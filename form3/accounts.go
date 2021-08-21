package accounts

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"

	// "io/ioutil"
	// "log"
	"net/http"
)

const (
	uRL = "http://localhost:8080/v1/organisation/accounts/"
)

// POST /v1/organisation/accounts
func Create(payload string) (*http.Response, error) {
	reqBody := strings.NewReader(payload)
	
	res, err := http.Post(uRL, "application/json", reqBody)
	if err != nil {
		return nil, err
	}
	// body, _ := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	return res, nil
}

// GET /v1/organisation/accounts/{account_id}
func Fetch(accountId string) (*http.Response, error) {
	//Build The URL string

	//We make HTTP request using the Get function
	res, err := http.Get(uRL + accountId)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	//Create a variable of the same type as our model
	
	//Invoke the text output function & return it with nil as the error value
	return res, nil
}

// DELETE /v1/organisation/accounts/{account_id}?version={version}
func Delete(accountId string, version string) (*http.Response, error){
	// Create client
    client := &http.Client{}

    // Create request
    req, err := http.NewRequest("DELETE", uRL + accountId, nil)
    if err != nil {
        return nil, err
    }

    // Fetch Request
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Read Response Body
    // respBody, err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //     return nil, err
    // }

    return resp, nil
}
