package fhir

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetDiagnosticReport will return a diagnostic report for a patient with id pid
func (c *Connection) GetDiagnosticReport(pid string) (*DiagnosticReport, error) {
	fmt.Printf("DiagnosticReport?patient=%v", pid)
	b, err := c.Query(fmt.Sprintf("DiagnosticReport?patient=%v", pid))
	if err != nil {
		return nil, err
	}
	data := DiagnosticReport{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// DiagnosticReport is a FHIR report
type DiagnosticReport struct {
	SearchResult
	Entry []ResourceEntry `json:"entry"`

	// Entry []struct {
	// 	EntryPartial
	// 	Resource struct {
	// 		ResourcePartial
	// 		EffectiveDateTime time.Time    `json:"effective_date_time"`
	// 		Issued            time.Time    `json:"issued"`
	// 		Identifier        []Identifier `json:"identifier"`
	// 		Meta              MetaData     `json:"meta"`
	// 		Text              TextData     `json:"text"`
	// 		Category          CodeText     `json:"category"`
	// 		Code              CodeText     `json:"code"`
	// 		PresentedForm     []Attachment `json:"presentedForm"`
	// 		Request           []Thing      `json:"request"`
	// 		Encounter         Encounter    `json:"encounter"`
	// 		Result            []Thing      `json:"result"`
	// 	} `json:"resource"`
	// } `json:"entry"`
}

type ResourceEntry struct {
	EntryPartial EntryPartial
	Resource     DiagnosticReportResource `json:"resource"`
}

type DiagnosticReportResource struct {
	ResourcePartial   ResourcePartial
	EffectiveDateTime time.Time    `json:"effective_date_time"`
	Issued            time.Time    `json:"issued"`
	Identifier        []Identifier `json:"identifier"`
	Meta              MetaData     `json:"meta"`
	Text              TextData     `json:"text"`
	Category          CodeText     `json:"category"`
	Code              CodeText     `json:"code"`
	PresentedForm     []Attachment `json:"presentedForm"`
	Request           []Thing      `json:"request"`
	Encounter         Encounter    `json:"encounter"`
	Result            []Thing      `json:"result"`
}

//Category the DiagnosticReport Category
type Category struct {
	Text string `json:"text"`
}

//Encounter of the report
type Encounter struct {
	Reference string `json:"reference"`
}

//MetaData meta field in DiagnosticReport
type MetaData struct {
	VersionID   string    `json:"versionId"`
	LastUpdated time.Time `json:"lastUpdated"`
}

//TextData is the text of a diagonistic report
type TextData struct {
	Status string `json:"status"`
	Div    string `json:"div"`
}

// // Return the actual decoded text attachment
// func (a *Attachment) DecodeImage() (string, error) {
// 	switch a.ContentType {
// 	case "text/html":
// 		data, err := decodeURL(a.URL)
// 		// ...
// 	case "application/pdf":
// 		data, err := decodeURL(a.URL)
// 		// ...
// 	}
// 	return data, err
// }

// func decodeURL(url string, filePath string) (string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	out, err := os.Create(filePath)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer out.Close()

// }
