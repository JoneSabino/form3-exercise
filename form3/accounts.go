package accounts

import (
	// "bytes"
	// "encoding/json"
	// "io/ioutil"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	// "io/ioutil"
	// "log"
	"net/http"

	"github.com/JoneSabino/form3-exercise/model"
	"github.com/satori/go.uuid"
)

const (
	uRL = "http://localhost:8080/v1/organisation/accounts/"
)

// POST /v1/organisation/accounts
func Create(payload string) (*http.Response, error) {
	// id := uuid.NewV4().String()
	// organizationId := uuid.NewV4().String()

	
	reqBody := strings.NewReader(payload)
	
	res, err := http.Post(uRL, "application/json", reqBody)
	if err != nil {
		return nil, err
	}
	// body, _ := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	return res, nil
}

func Create2(payload model.AccountData) (*http.Response, error) {
	id := uuid.NewV4().String()
	organizationId := uuid.NewV4().String()

	payload.ID = id
	payload. OrganisationID = organizationId

	str, _ := json.MarshalIndent(payload,"","\t")
	fmt.Println(string(str))

	payloadJson, err := json.Marshal(id)
	responseBody := bytes.NewBuffer(payloadJson)

	if err != nil {	
		return  nil, err
	}

	// reqBody := strings.NewReader(payload)
	
	res, err := http.Post(uRL, "application/json", responseBody)
	if err != nil {
		return nil, err
	}
	// body, _ := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	return res, nil	
}

// GET /v1/organisation/accounts/{account_id}
func Fetch(accountId string) (model.AccountAttributes, error) {
	//Build The URL string

	//We make HTTP request using the Get function
	res, err := http.Get(uRL + accountId)

	if err != nil {
		return model.AccountAttributes{}, err
	}
	jsonRes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return model.AccountAttributes{}, err
	}

	var accResp model.AccountAttributes
	err = json.Unmarshal(jsonRes, &accResp)

	if err != nil {
		return model.AccountAttributes{}, err
	}
	// defer res.Body.Close()
	//Create a variable of the same type as our model
	
	//Invoke the text output function & return it with nil as the error value
	return accResp, nil
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
	q := req.URL.Query()
    q.Add("version", version)
    req.URL.RawQuery = q.Encode()

    
    // Fetch Request
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    // defer resp.Body.Close()

    // Read Response Body
    // respBody, err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //     return nil, err
    // }

    return resp, nil
}
