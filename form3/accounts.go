package form3

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/JoneSabino/form3-exercise/model"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	id := uuid.NewV4().String()
	// id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	accData.Type = "accounts"
	accData.ID = id

	payload := model.Account{
		Data: &accData,
	}
	
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return model.Account{}, errors.New(err.Error() +
		"\nFailed to marshal the payload")
	}

	var resp *http.Response
	var respBody []byte
	for i := 0; i < 3; i++ {
		reqBody := bytes.NewBuffer(payloadJson)
		resp, err = http.Post(uRL, "application/json", reqBody)
		if err != nil {
			return model.Account{}, errors.New(err.Error() +
				"\nFailed sending the POST request")
		}

		if resp.StatusCode == http.StatusConflict {
			log.Println(resp.Status + ". Probably the value of the `id` field alredy exists. Retrying...")
			if i == 2 {
				return model.Account{}, errors.New(strconv.Itoa(resp.StatusCode) + ": Duplicated Account ID")
			}
			continue
		} else if resp.StatusCode != http.StatusCreated {
			respBody, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return model.Account{}, errors.New(err.Error() +
					"\nThere was a problem when trying to read the response body")
			}
			return model.Account{}, errors.New(resp.Status + "\n" +
				string(respBody))
		} else {
			break
		}
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Account{}, errors.New(err.Error() +
			"\nThere was a problem when trying to read the response body")
	}

	accModel, err := mapResponse(respBody)

	return accModel, nil
}

// Receives the account id and returns the response body content struct
//
// Parameters:
//  - `accountId` : string
// Request: GET /v1/organisation/accounts/{account_id}
func Fetch(accountId string) (model.Account, error) {

	resp, err := http.Get(uRL + accountId)
	if err != nil {
		return model.Account{}, errors.New(err.Error() +
			"\nFailed sending the GET request")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Account{}, errors.New(err.Error() +
			"\nThere was a problem when trying to read the response body")
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("ERROR - Status Code:", resp.StatusCode)
		return model.Account{}, errors.New(string(respBody))
	}

	respModel, err := mapResponse(respBody)

	return respModel, nil
}

// Deletes the account specified.
//
// Returns a string with the response status.
//
// A successul request will return '204 No Content'
//Parameters:
//	- `AccountId` : string
//	- `Version` : string
// Request: DELETE /v1/organisation/accounts/{account_id}?version={version}
func Delete(accountId string, version string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", uRL+accountId, nil)
	if err != nil {
		return "", errors.New(err.Error() +
			"\n Error while creating the request")
	}

	q := req.URL.Query()
	q.Add("version", version)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(err.Error() +
			"\n Failed sending the DELETE request")
	}
	return resp.Status, nil
}

func mapResponse(respBody []byte) (model.Account, error) {
	var accModel model.Account
	err := json.Unmarshal(respBody, &accModel)
	if err != nil {
		return model.Account{}, errors.New(err.Error() +
			"\nFailed when trying map the response body to the Account structure")
	}
	return accModel, nil
}
