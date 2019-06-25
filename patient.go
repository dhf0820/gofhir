package fhir

import (
	"encoding/json"
	"fmt"
)

// GetPatient will return patient information for a patient with id pid
func (c *Connection) GetPatient(pid string) (*Patient, error) {
	fmt.Printf("FHIR GetPatient url: %s/Patient/%v", c.BaseURL, pid)
	b, err := c.Query(fmt.Sprintf("Patient/%v", pid))
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\n\n\n@@@ Patient 15 RAW Patient: %s\n\n\n", pretty.Pretty(b))
	data := Patient{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Connection) FindFhirPatient(qry string) (*PatientResult, error) {
	fmt.Printf("QRY: %s\n", qry)
	fmt.Printf("With v: Patient?%v\n", qry)
	fmt.Printf("Patient?%s\n", qry)
	fmt.Printf("FHIR FindPatient url: %sPatient?%s\n", c.BaseURL, qry)
	b, err := c.Query(fmt.Sprintf("Patient?%v", qry))
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\n\n\n@@@ Patient 15 RAW Patient: %s\n\n\n", pretty.Pretty(b))
	data := PatientResult{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Connection) FindFhirPatients(qry string) (*PatientResult, error) {
	fmt.Printf("QRY: %s\n", qry)
	fmt.Printf("With v: Patient?%v\n", qry)
	fmt.Printf("Patient?%s\n", qry)
	fmt.Printf("FHIR FindPatient url: %sPatient?%s\n", c.BaseURL, qry)
	b, err := c.Query(fmt.Sprintf("Patient?%v", qry))
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\n\n\n@@@ Patient 15 RAW Patient: %s\n\n\n", pretty.Pretty(b))
	data := PatientResult{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Patient is a FHIR patient
type Patient struct {
	ResourceType    string          `json:"resourceType"`
	BirthDate       string          `json:"birthDate"`
	Active          bool            `json:"active"`
	Gender          string          `json:"gender"`
	DeceasedBoolean bool            `json:"deceasedBoolean"`
	ID              string          `json:"id"`
	Text            TextData        `json:"text"`
	CareProvider    []Person        `json:"careProvider"`
	Name            []Name          `json:"name"`
	Identifier      []Identifier    `json:"identifier"`
	Address         []Address       `json:"address"`
	Telecom         []Telecom       `json:"telecom"`
	MaritalStatus   Concept         `json:"maritalStatus"`
	Communication   []Communication `json:"communication"`
	Extension       []Extension     `json:"extension"`
}

type PatientBundle struct {
	SearchResult
	Entry []struct {
		EntryPartial
		Resource struct {
			Patient
		}
	}
}
