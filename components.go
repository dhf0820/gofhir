package fhir

import "time"

// SearchResult is a search result
type SearchResult struct {
	ResourceType string  `json:"resourceType"`
	Type         string  `json:"type"`
	Total        int     `json:"total"`
	Link         []Link  `json:"link"`
	ID           string  `json:"id"`
	Issues       []Issue `json:"issue"`
}

// Issue is a FHIR issue
type Issue struct {
	Severity string   `json:"severity"`
	Location []string `json:"location"`
	Code     string   `json:"code"`
	Details  Concept  `json:"details"`
}

// EntryPartial are the common entry fields
type EntryPartial struct {
	FullURL string     `json:"fullUrl"`
	Link    []Link     `json:"link"`
	Search  SearchMode `json:"search"`
}

// ResourcePartial are the common resource fields
type ResourcePartial struct {
	ResourceType      string             `json:"resourceType"`
	EffectiveDateTime time.Time          `json:"effectiveDateTime"`
	RecordedDate      time.Time          `json:"recordedDate"`
	Status            string             `json:"status"`
	ID                string             `json:"id"`
	Subject           Person             `json:"subject"`
	Patient           Person             `json:"patient"`
	Performer         Person             `json:"performer"`
	Recorder          Person             `json:"recorder"`
	Encounter         EncounterReference `json:"Encounter"`
}

// Bundle is the header for any returned Bundle
type Bundle struct {
	ResourceType string `json:"resourceType"`
	ID           string `json:"id"`
	Type         string `json:"type"`
	Link         []Link `json:"link"`
}

// SearchMode is a FHIR search mode
type SearchMode struct {
	Mode string `json:"mode"`
}

// Note is a note
type Note struct {
	Text string `json:"text"`
}

// CodeText is a healthcare condition
type CodeText struct {
	Code Note `json:"code"`
}

// Attachment is a url attachment
type Attachment struct {
	ContentType string `json:"contentType"`
	URL         string `json:"url"`
}

// Link is a resource link
type Link struct {
	Relation string `json:"relation"`
	URL      string `json:"url"`
}

// Period is a period of time
type Period struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// Coding is a code and system
type Coding struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display"`
}

// Concept is a general concept such as language
type Concept struct {
	Text   string   `json:"text"`
	Coding []Coding `json:"coding"`
}

// Thing is a FHIR thing
type Thing struct {
	Display   string `json:"display"`
	Reference string `json:"reference"`
}

// Person is a human
type Person Thing

// Name is a persons name
type Name struct {
	Use    string   `json:"use"`
	Family []string `json:"family"`
	Given  []string `json:"given"`
	Suffix []string `json:"suffix"`
	Prefix []string `json:"prefix"`
}

// Identifier can identify things
type Identifier struct {
	Use    string  `json:"use"`
	System string  `json:"system"`
	Value  string  `json:"value"`
	Type   Concept `json:"type"`
}

// Address is a physical address
type Address struct {
	Use        string   `json:"use"`
	Line       []string `json:"line"`
	City       string   `json:"city"`
	State      string   `json:"state"`
	PostalCode string   `json:"postalCode"`
	Country    string   `json:"country"`
	Period     Period   `json:"period,omitempty"`
}

// Telecom is a phone number
type Telecom struct {
	System string `json:"system"`
	Value  string `json:"value"`
	Use    string `json:"use,omitempty"`
	Period Period `json:"period,omitempty"`
}

// Communication is the language people speak
type Communication struct {
	Preferred bool    `json:"preferred"`
	Language  Concept `json:"language"`
}

// Context encounter only initially
type Context struct {
	EncounterRef EncounterReference `json:"encounter"`
}

// Extension is a codified FHIR extension
type Extension struct {
	URL                  string  `json:"url"`
	ValueCodeableConcept Concept `json:"valueCodeableConcept"`
}

// Reaction is a human reaction
type Reaction struct {
	Certainty     string    `json:"certainty"`
	Onset         time.Time `json:"onset"`
	Manifestation []Note    `json:"manifestation"`
	Note          Note      `json:"note"`
}

// Quantity is a quantity of something
type Quantity struct {
	Value  float32 `json:"value"`
	Unit   string  `json:"unit"`
	Code   string  `json:"code"`
	System string  `json:"system"`
}

// Repeat is a time based repeat of something
type Repeat struct {
	Frequency    float32 `json:"frequency"`
	Period       float32 `json:"period"`
	PeriodUnits  string  `json:"periodUnits"`
	BoundsPeriod Period  `json:"boundsPeriod"`
}

// Timing is the timing of something
type Timing struct {
	Repeat Repeat `json:"timing"`
}

// DosageInstruction are the medication instructions for dosage
type DosageInstruction struct {
	Text            string   `json:"text"`
	AsNeededBoolean bool     `json:"asNeededBoolean"`
	Route           Concept  `json:"route"`
	Method          Concept  `json:"method"`
	Timing          Timing   `json:"timing"`
	DoseQuantity    Quantity `json:"doseQuantity"`
}

// DispenseRequest is a dispensing request
type DispenseRequest struct {
	ValidityPeriod Period `json:"validityPeriod"`
}

//Category the DiagnosticReport Category
type Category struct {
	Text string `json:"text"`
}

//EncounterReference of the report
type EncounterReference struct {
	Reference string `json:"reference"`
}

//MetaData meta field in DocumentReference/DiagnosticReport
type MetaData struct {
	VersionID   string    `json:"versionId"`
	LastUpdated time.Time `json:"lastUpdated"`
}

//TextData is the html text
type TextData struct {
	Status string `json:"status"`
	Div    string `json:"div"`
}

type Code string
type Status string

type BackboneElement struct {
	Type       []CodeableConcept
	Period     Period
	Individual Person
}

type CodeableConcept struct {
	Coding []Coding
	Text   string `json:"text"`
}

type Location struct {
	Location struct {
		Reference string `json:"reference"`
		Display   string `json:"display"`
	}
	Status       Status          `json:"status"`
	PhysicalType CodeableConcept `json:"physicalType"`
}

type ServiceProvider struct {
	Reference string `json:"reference"`
}

type Reason struct {
	Text string `json:"text"`
}

// type Reference struct {
// 	Reference  string `json:"reference"`
// 	Identifier Identifier
// 	Display    string `json:"display"`
// }
