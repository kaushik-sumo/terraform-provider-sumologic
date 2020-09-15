// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated. Do not submit
//     changes to this file.
//
// ----------------------------------------------------------------------------
package sumologic

import (
	"encoding/json"
	"fmt"
)

func (s *Client) CreateIngestBudgetV2(ingestBudgetV2 IngestBudgetV2) (string, error) {
	urlWithoutParams := "v2/ingestBudgets"

	data, err := s.Post(urlWithoutParams, ingestBudgetV2)
	if err != nil {
		return "", err
	}

	var createdIngestBudgetV2 IngestBudgetV2

	err = json.Unmarshal(data, &createdIngestBudgetV2)
	if err != nil {
		return "", err
	}

	return createdIngestBudgetV2.ID, nil

}

func (s *Client) GetIngestBudgetV2(id string) (*IngestBudgetV2, error) {
	urlWithoutParams := "v2/ingestBudgets/%s"
	paramString := ""
	sprintfArgs := []interface{}{}
	sprintfArgs = append(sprintfArgs, id)

	urlWithParams := fmt.Sprintf(urlWithoutParams+paramString, sprintfArgs...)

	data, _, err := s.Get(urlWithParams)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	var ingestBudgetV2 IngestBudgetV2

	err = json.Unmarshal(data, &ingestBudgetV2)
	if err != nil {
		return nil, err
	}

	return &ingestBudgetV2, nil

}

func (s *Client) DeleteIngestBudgetV2(id string) error {
	urlWithoutParams := "v2/ingestBudgets/%s"
	paramString := ""
	sprintfArgs := []interface{}{}
	sprintfArgs = append(sprintfArgs, id)

	urlWithParams := fmt.Sprintf(urlWithoutParams+paramString, sprintfArgs...)

	_, err := s.Delete(urlWithParams)

	return err
}

func (s *Client) UpdateIngestBudgetV2(ingestBudgetV2 IngestBudgetV2) error {
	urlWithoutParams := "v2/ingestBudgets/%s"
	paramString := ""
	sprintfArgs := []interface{}{}
	sprintfArgs = append(sprintfArgs, ingestBudgetV2.ID)

	urlWithParams := fmt.Sprintf(urlWithoutParams+paramString, sprintfArgs...)

	ingestBudgetV2.ID = ""

	_, err := s.Put(urlWithParams, ingestBudgetV2)

	return err

}

type IngestBudgetV2 struct {
	AuditThreshold int    `json:"auditThreshold"`
	Action         string `json:"action"`
	ResetTime      string `json:"resetTime"`
	Name           string `json:"name"`
	ID             string `json:"id,omitempty"`
	Scope          string `json:"scope"`
	Timezone       string `json:"timezone"`
	Description    string `json:"description"`
	CapacityBytes  int    `json:"capacityBytes"`
}