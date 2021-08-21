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
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
	ProcessingService       string   `json:"processing_service,omitempty"`
	UserDefinedInformation  string   `json:"user_defined_information,omitempty"`
	ValidationType          string   `json:"validation_type,omitempty"`
	ReferenceMask           string   `json:"reference_mask,omitempty"`
	AcceptanceQualifier     string   `json:"acceptance_qualifier,omitempty"`
}
