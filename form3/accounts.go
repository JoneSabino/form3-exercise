package accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/JoneSabino/form3-exercise/model"
	"github.com/satori/go.uuid"
)

const (
	uRL = "http://localhost:8080/v1/organisation/accounts/"
)

// // POST /v1/organisation/accounts
func Create(accAttr model.AccountAttributes) (model.Account, error) {
	id := uuid.NewV4().String()
	organizationId := uuid.NewV4().String()
	
	payload := model.Account{
		Data: &model.AccountData{
			Attributes:     &accAttr,
			ID:             id,
			OrganisationID: organizationId,
			Type:           "accounts",
		},
	}

	payloadJson, err := json.Marshal(payload)
	reqBody := bytes.NewBuffer(payloadJson)

	if err != nil {	
		return  model.Account{}, err
	}

	resp, err := http.Post(uRL, "application/json", reqBody)
	if err != nil {
		return model.Account{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	var accModel model.Account
    err = json.Unmarshal(respBody, &accModel)

	return accModel, nil	
}

// GET /v1/organisation/accounts/{account_id}
func Fetch(accountId string) (model.Account, error){
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

// DELETE /v1/organisation/accounts/{account_id}?version={version}
func Delete(accountId string, version string) (string, error){
    client := &http.Client{}

    req, err := http.NewRequest("DELETE", uRL + accountId, nil)
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
