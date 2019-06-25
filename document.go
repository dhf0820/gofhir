package fhir

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tidwall/pretty"
)

// GetDocumentReference will return one document
func (c *Connection) GetDocumentReference(pid string) (*DocumentReference, error) {
	fmt.Printf("\n%sDocumentReference/%v\n", c.BaseURL, pid)
	b, err := c.Query(fmt.Sprintf("DocumentReference/%v", pid))
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n\n\n@@@ RAW DocumentReference: %s\n\n\n", pretty.Pretty(b))
	data := DocumentReference{}
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Printf("GetDocumentReference err: %v\n, err")
		return nil, err
	}
	return &data, nil
}

// GetDocumentReference will return one document
func (c *Connection) GetDocumentReferences(pid string) (*DocumentReferences, error) {
	fmt.Printf("%sDocumentReference?%v\n", c.BaseURL, pid)
	b, err := c.Query(fmt.Sprintf("DocumentReference?%v", pid))
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\n\n\n@@@ RAW DocumentReference: %s\n\n\n", pretty.Pretty(b))
	data := DocumentReferences{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// DocumentReference is a FHIR document
type DocumentReferences struct {
	Bundle
	Entry []struct {
		FullURL  string `json:"fullUrl"`
		Resource struct {
			ResourceType  string    `json:"resourceType"`
			ID            string    `json:"id"`
			Meta          MetaData  `json:"meta"`
			Text          TextData  `json:"text"`
			Subject       Person    `json:"subject"`
			Type          Concept   `json:"type"`
			Authenticator Person    `json:"authenticator"`
			Created       time.Time `json:"created"`
			Indexed       time.Time `json:"indexed"`
			DocStatus     Concept   `json:"docStatus"`
			Description   string    `json:"description"`

			Content []struct {
				Attachment struct {
					ContentType string `json:"contentType"`
					URL         string `json:"url"`
				} `json:"attachment"`
			} `json:"content"`
			//} `json:"content"`
			//Content       []Attachment `json:"content"`
			Context struct {
				EncounterNum struct {
					Reference string `json:"reference"`
				} `json:"encounter"`
			} `json:"context"`
			//Context EncounterReference `json:"context"`

			// ResourcePartial
			// //		Encounter        Thing        `json:"encounter"`
			// Created          time.Time    `json:"created"`
			// Indexed          time.Time    `json:"indexed"`
			// Class            Concept      `json:"class"`
			// Type             Concept      `json:"type"`
			// Content          []Attachment `json:"content"`
			// MasterIdentifier Identifier   `json:"masterIdentifier"`
		} `json:"resource"`
	} `json:"entry"`
}

// DocumentReference is a single FHIR DocumentReference.
// Use DocumentReferences for a bundle.
type DocumentReference struct {
	ResourceType      string    `json:"resourceType"`
	ID                string    `json:"id"`
	EffectiveDateTime time.Time `json:"effectiveDateTime"`
	Meta              MetaData  `json:"meta"`
	Text              TextData  `json:"text"`
	Subject           Person    `json:"subject"`
	Type              Concept   `json:"type"`
	Authenticator     Person    `json:"authenticator"`
	Created           time.Time `json:"created"`
	Indexed           time.Time `json:"indexed"`
	DocStatus         Concept   `json:"docStatus"`
	Description       string    `json:"description"`

	Content []struct {
		Attachment struct {
			ContentType string `json:"contentType"`
			URL         string `json:"url"`
		} `json:"attachment"`
	} `json:"content"`
	//} `json:"content"`
	//Content       []Attachment `json:"content"`
	Context struct {
		EncounterNum struct {
			Reference string `json:"reference"`
		} `json:"encounter"`
	} `json:"context"`
}
