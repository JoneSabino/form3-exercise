package model

import "time"

type CreateResponse struct {
	Data *CreateResponseData `json:"data,omitempty"`
}

type CreateResponseData struct {
	Attributes     *CreateResponseAttributes `json:"attributes,omitempty"`
	CreatedOn      time.Time                 `json:"created_on,omitempty"`
	ID             string                    `json:"id,omitempty"`
	ModifiedOn     time.Time                 `json:"modified_on,omitempty"`
	OrganisationID string                    `json:"organisation_id,omitempty"`
	Type           string                    `json:"type,omitempty"`
	Version        int                       `json:"version,omitempty"`
	Links          *CreateResponseLinks      `json:"links,omitempty"`
}

type CreateResponseLinks struct {
	Self string `json:"self,omitempty"`
}

type CreateResponseAttributes struct {
	AccountClassification string   `json:"account_classification,omitempty"`
	AlternativeNames      []string `json:"alternative_names,omitempty"`
	BankID                string   `json:"bank_id,omitempty"`
	BankIDCode            string   `json:"bank_id_code,omitempty"`
	Bic                   string   `json:"bic,omitempty"`
	Country               string   `json:"country,omitempty"`
	Name                  []string `json:"name,omitempty"`
}
