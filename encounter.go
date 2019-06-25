package fhir

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetEncounter will return an Encounter for a number (Encounter)
func (c *Connection) GetEncounter(eid string) (*Encounter, error) {
	fmt.Printf("%sEncounter/%s\n\n", c.BaseURL, eid)
	b, err := c.Query(fmt.Sprintf("Encounter/%s", eid))
	if err != nil {
		return nil, err
	}
	data := Encounter{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetEncounter will return Encounters for a patient with id pid
func (c *Connection) GetPatientEncounters(pid string) (*EncounterResults, error) {
	fmt.Printf("%sEncounter?patient=%s\n", c.BaseURL, pid)
	b, err := c.Query(fmt.Sprintf("Encounter?patient=%s", pid))

	if err != nil {
		fmt.Printf("Encounter Query, %s\n", err)
		return nil, err
	}
	//fmt.Printf("\n\n\n@@@ RAW Encounter: %s\n\n\n", pretty.Pretty(b))
	data := EncounterResults{}
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}
	// fmt.Printf("\n\n\nUnmarshalled:")
	//spew.Dump(data)
	return &data, nil
}

// GetEncounters will return Encounters for a patient with id pid
func (c *Connection) FindFhirEncounters(query string) (*EncounterResults, error) {
	fmt.Printf("%sEncounter?%s\n", c.BaseURL, query)
	b, err := c.Query(fmt.Sprintf("Encounter?%s", query))
	if err != nil {

		return nil, err
	}
	//fmt.Printf("\n\n\n@@@ RAW Encounter: %s\n\n\n", pretty.Pretty(b))
	data := EncounterResults{}
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}
	// fmt.Printf("\n\n\nUnmarshalled:")
	// spew.Dump(data)
	return &data, nil
}

type EncounterResults struct {
	SearchResult
	Entry []struct {
		EntryPartial
		Resource Encounter
		// Resource struct {
		// 	//ResourceType string `json:"resourceType"`
		// 	ResourcePartial
		// 	Code        Code     `json:"code"`
		// 	Category    Code     `json:"caategory"`
		// 	Description string   `json:"description"`
		// 	Text        TextData `json:"text"`
		// }
	}
}

type Encounter struct {
	ResourceType      string       `json:"resourceType"`
	EffectiveDateTime time.Time    `json:"effectiveDateTime"`
	RecordedDate      time.Time    `json:"recordedDate"`
	Status            string       `json:"status"`
	Type              []Concept    `json:"type"`
	Class             string       `json:"class"`
	ID                string       `json:"id"`
	Identifiers       []Identifier `json:"identifier"`
	Subject           Person       `json:"subject"`
	Patient           Person       `json:"patient"`
	Performer         Person       `json:"performer"`
	Recorder          Person       `json:"recorder"`
	Code              Code         `json:"code"`
	Category          Code         `json:"caategory"`
	Reasons           []*Reason    `json:"reason"`
	Description       string       `json:"description"`
	Text              TextData     `json:"text"`
	//Participant       []BackboneElement
	Period          Period
	Location        []Location
	ServiceProvider ServiceProvider
}
