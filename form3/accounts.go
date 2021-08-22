package accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/JoneSabino/form3-exercise/model"
)

const (
	uRL = "http://localhost:8080/v1/organisation/accounts/"
)

// Receives a structure containing the account creation request body
// Returns the response body content struct
//
// Parameters:
//  - `AccAttr`: model.AccountAttribute
// Request: POST /v1/organisation/accounts
func Create(accData model.AccountData) (model.Account, error) {
	// id := uuid.NewV4().String()
	id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	accData.Type = "accounts"
	accData.ID = id

	payload := model.Account{
		Data: &accData,
	}

	payloadJson, err := json.Marshal(payload)
	reqBody := bytes.NewBuffer(payloadJson)

	if err != nil {
		return model.Account{}, err
	}
	var resp *http.Response
	for i := 0; i < 3; i++ {
		resp, err := http.Post(uRL, "application/json", reqBody)
		if err != nil {
			return model.Account{}, err
		}

		if resp.StatusCode == 409 {
			log.Println(resp.Status + ". Probably the value of the `id` field alredy exists. Retrying...")
			time.Sleep(1)
		}
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	var accModel model.Account
	err = json.Unmarshal(respBody, &accModel)

	return accModel, nil
}

// Receives the account id and returns the response body content struct
//
// Parameters:
//  - `accountId` : string
// Request: GET /v1/organisation/accounts/{account_id}
func Fetch(accountId string) (model.Account, error) {

	resp, err := http.Get(uRL + accountId)

	respBody, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("ERROR - Status Code:", resp.StatusCode)
		return model.Account{}, err
	}

	if err != nil {
		panic(err)
	}

	var respModel model.Account
	err = json.Unmarshal(respBody, &respModel)

	if err != nil {
		panic(err)
	}

	return respModel, nil
}

// Deletes the account specified.
//
// Returns a string with the response status:
//  - Success: '204 No Content'
//Parameters:
//	- `AccountId` :
// Request: DELETE /v1/organisation/accounts/{account_id}?version={version}
func Delete(accountId string, version string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", uRL+accountId, nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("version", version)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	return resp.Status, nil
}
