package form3

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
	"github.com/JoneSabino/form3-exercise/pkg/model"
	uuid "github.com/satori/go.uuid"
)

var (
	accountId string
	version   string
)

// TestCreate calls accounts.Create, checking
// for a valid return value.
func TestCreate(t *testing.T) {
	class, ct := new(string), new(string)
	*class = "Personal"
	*ct = "GB"

	accAttrs := model.AccountAttributes{
		AccountClassification: class,
		AlternativeNames:      []string{"bibi"},
		BankID:                "100000",
		BankIDCode:            "GBDSC",
		Bic:                   "NWBKGB42",
		Country:               ct,
		Name:                  []string{"bianca"},
	}

	accData := model.AccountData{
		Attributes:     &accAttrs,
		OrganisationID: uuid.NewV4().String(),
	}

	resp, err := Create(accData) // should return the response body as a struct
	accountId = resp.Data.ID
	if (model.Account{}) != resp {
		resp1, _ := json.Marshal(resp)
		t.Log(string(resp1))
	} else {
		t.Error("Expected a successful Create() return, found:\n" + err.Error())
	}
}

// TestCreateMissingReqFields calls accounts.Create, checking
// for a valid error message.
//
// Test Cases:
//  - Missing Country Value
//  - Missing Country Field
//  - Missing Name Field
//  - Missing Name Value 1
//  - Missing Name Value 2
//  - Missing OrganisationID Field
//  - Missing OrganisationID Value
func TestCreateMissingReqFields(t *testing.T) {
	// Missing Country Value
	ct := new(string)
	*ct = ""

	accAttrs := model.AccountAttributes{
		Country: ct,
		Name:    []string{"bianca"},
	}

	accData := model.AccountData{
		Attributes:     &accAttrs,
		OrganisationID: uuid.NewV4().String(),
	}

	_, err := Create(accData) // should return the expected error message
	if err.Error() == "Organisation ID, Country and all Names must be filled" {
		t.Log(err.Error())
	} else {
		t.Error("Error message differs from expected:\n" + err.Error())
	}

	// Missing Country Field
	accAttrs1 := model.AccountAttributes{
		Name: []string{"bianca"},
	}

	accData1 := model.AccountData{
		Attributes:     &accAttrs1,
		OrganisationID: uuid.NewV4().String(),
	}

	_, err = Create(accData1) // should return the expected error message
	if err.Error() == "Organisation ID, Country and all Names must be filled" {
		t.Log(err.Error())
	} else {
		t.Error("Error message differs from expected:\n" + err.Error())
	}

	// Missing Name Field
	*ct = "GB"
	accAttrs2 := model.AccountAttributes{
		Country: ct,
	}

	accData2 := model.AccountData{
		Attributes:     &accAttrs2,
		OrganisationID: uuid.NewV4().String(),
	}

	_, err = Create(accData2) // should return the expected error message
	if err.Error() == "Organisation ID, Country and all Names must be filled" {
		t.Log(err.Error())
	} else {
		t.Error("Error message differs from expected:\n" + err.Error())
	}

	// Missing Name Value 1
	accAttrs3 := model.AccountAttributes{
		Country: ct,
		Name:    []string{""},
	}

	accData3 := model.AccountData{
		Attributes:     &accAttrs3,
		OrganisationID: uuid.NewV4().String(),
	}

	_, err = Create(accData3) // should return the expected error message
	if err.Error() == "Organisation ID, Country and all Names must be filled" {
		t.Log(err.Error())
	} else {
		t.Error("Error message differs from expected:\n" + err.Error())
	}

	// Missing Name Value 2
	accAttrs4 := model.AccountAttributes{
		Country: ct,
		Name:    []string{},
	}

	accData4 := model.AccountData{
		Attributes:     &accAttrs4,
		OrganisationID: uuid.NewV4().String(),
	}

	_, err = Create(accData4) // should return the expected error message
	if err.Error() == "Organisation ID, Country and all Names must be filled" {
		t.Log(err.Error())
	} else {
		t.Error("Error message differs from expected:\n" + err.Error())
	}

	// Missing OrganisationID Field
	accAttrs5 := model.AccountAttributes{
		Country: ct,
		Name:    []string{"Bianca"},
	}

	accData5 := model.AccountData{
		Attributes: &accAttrs5,
	}

	_, err = Create(accData5) // should return the expected error message
	if err.Error() == "Organisation ID, Country and all Names must be filled" {
		t.Log(err.Error())
	} else {
		t.Error("Error message differs from expected:\n" + err.Error())
	}

	// Missing OrganisationID Value
	accAttrs6 := model.AccountAttributes{
		Country: ct,
		Name:    []string{"bianca sabino"},
	}

	accData6 := model.AccountData{
		Attributes:     &accAttrs6,
		OrganisationID: "",
	}

	_, err = Create(accData6) // should return the expected error message
	if err.Error() == "Organisation ID, Country and all Names must be filled" {
		t.Log(err.Error())
	} else {
		t.Error("Error message differs from expected:\n" + err.Error())
	}
}

// TestFetch calls accounts.Fetch, checking
// for a successful response.
func TestFetch(t *testing.T) {
	resp, err := Fetch(accountId)
	version = strconv.FormatInt(*resp.Data.Version, 10)
	if (model.Account{}) != resp {
		resp1, _ := json.Marshal(resp)
		t.Log(string(resp1))
	} else {
		t.Error("Expected a successful Create() return, found:\n" + err.Error())
	}
}

// TestDelete calls accounts.Delete, checking
// for a successful response.
func TestDelete(t *testing.T) {
	version = "1"
	resp, err := Delete(accountId, version)
	if resp == http.StatusNoContent{
		t.Log("Success!")
	} else {
		t.Error(err)
	}

}
