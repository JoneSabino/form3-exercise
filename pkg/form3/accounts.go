package form3

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/JoneSabino/form3-exercise/pkg/model"
	uuid "github.com/satori/go.uuid"
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
	err := checkReqFields(accData)
	if err != nil {
		return model.Account{}, err
	}

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
			respBody, err = readResponseBody(resp.Body)
			if err != nil {
				return model.Account{}, err
			}
			return model.Account{}, errors.New(resp.Status + "\n" +
				string(respBody))
		} else {
			break
		}
	}

	respBody, err = readResponseBody(resp.Body)
	if err != nil {
		return model.Account{}, err
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

	// respBody, err := ioutil.ReadAll(resp.Body)
	respBody, err := readResponseBody(resp.Body)
	if err != nil {
		return model.Account{}, err
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
func Delete(accountId string, version string) (int, error) {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", uRL+accountId, nil)
	if err != nil {
		return 0, errors.New(err.Error() +
			"\n Error while creating the request")
	}

	q := req.URL.Query()
	q.Add("version", version)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return 0, errors.New(err.Error() +
			"\n Failed sending the DELETE request")
	}
	if resp.StatusCode != http.StatusNoContent {
		respBody, err := readResponseBody(resp.Body)
		if err != nil {
			return 0, err
		}
		log.Println("ERROR - Status Code:", resp.StatusCode)
		return resp.StatusCode, errors.New(string(respBody))
	}
	return resp.StatusCode, nil
}

func readResponseBody(responseBody io.ReadCloser)([]byte, error){
	respBody, err := ioutil.ReadAll(responseBody)
		if err != nil {
			return nil, errors.New(err.Error() +
				"\nThere was a problem when trying to read the response body")
		}
	return respBody, err
}

// Maps json response to a structure
func mapResponse(respBody []byte) (model.Account, error) {
	var accModel model.Account
	err := json.Unmarshal(respBody, &accModel)
	if err != nil {
		return model.Account{}, errors.New(err.Error() +
			"\nFailed when trying map the response body to the Account structure")
	}
	return accModel, nil
}

// Check if the required fields are present in the structure
func checkReqFields(accData model.AccountData) (error){
	if accData.OrganisationID == "" ||
	accData.Attributes.Country == nil ||
	*accData.Attributes.Country == "" ||
	accData.Attributes.Name == nil || 
	len(accData.Attributes.Name) == 0 ||
	contains(accData.Attributes.Name, "") {
        return errors.New("Organisation ID, Country and all Names must be filled")
    }
	return nil
}

// Check if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}