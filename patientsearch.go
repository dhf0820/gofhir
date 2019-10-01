package fhir

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// PatientSearch will search for a patient based on the query string
// identifier, family, given, birthdate, gender, address, telecom
// i.e. family=Argonaut&given=Jason
func (c *Connection) PatientSearch(query string) (*PatientResult, error) {
	fmt.Printf("%sPatient?%v\n\n", c.BaseURL, query)
	startTime := time.Now()
	b, err := c.Query(fmt.Sprintf("Patient?%v", query))
	log.Infof("Query time: %s", time.Since(startTime))
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\n\n\n@@@ RAW Patient: %s\n\n\n", pretty.Pretty(b))

	startTime = time.Now()
	data := PatientResult{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	log.Infof("Unmarshal time: %s", time.Since(startTime))
	return &data, nil
}

// PatientResult is a patient search result
type PatientResult struct {
	SearchResult
	Entry []struct {
		EntryPartial
		Patient Patient `json:"resource"`
	} `json:"entry"`
}
