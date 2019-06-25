package fhir

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetDiagnosticReport will return a diagnostic report for a patient with id pid
func (c *Connection) FindDiagnosticReport(query string) (*DiagnosticReport, error) {
	fmt.Printf("\n\nFindDiagnosticReport : DiagnosticReport?%v\n", query)
	b, err := c.Query(fmt.Sprintf("DiagnosticReport?%s", query))
	//fmt.Printf("GetDiag error: %v\n", err)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("\n\n\n@@@ RAW DiagnosticReport: %s\n\n\n", pretty.Pretty(b))
	// spew.Dump(b)
	data := DiagnosticReport{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDiagnosticReport will return a diagnostic report for a patient with id pid
func (c *Connection) GetDiagnosticReport(id string) (*DiagnosticReport, error) {
	fmt.Printf("Get:  DiagnosticReport/%s\n", id)
	b, err := c.Query(fmt.Sprintf("DiagnosticReport/%s", id)) //?patient=%v", pid))
	//fmt.Printf("GetDiag error: %v\n", err)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("\n\n\n@@@ RAW DiagnosticReport: %s\n\n\n", pretty.Pretty(b))
	// spew.Dump(b)
	data := DiagnosticReport{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// GetPatientDiagnosticReport will return a diagnostic report for a patient with id pid
func (c *Connection) GetPatientDiagnosticReport(pid string) (*DiagnosticReport, error) {
	fmt.Printf("GetPat DiagnosticReport/%s\n", pid)
	b, err := c.Query(fmt.Sprintf("DiagnosticReport/?patient=%d", pid)) //, ", pid))
	if err != nil {
		return nil, err
	}
	// fmt.Printf("\n\n\n@@@ RAW DiagnosticReport: %s\n\n\n", pretty.Pretty(b))
	// spew.Dump(b)
	data := DiagnosticReport{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// DiagnosticReport is a FHIR report
type DiagnosticReport struct {
	SearchResult
	Entry []struct {
		EntryPartial
		Resource struct {
			ResourcePartial
			Encounter         Thing        `json:"encounter"`
			EffectiveDateTime time.Time    `json:"effective_date_time"`
			Issued            time.Time    `json:"issued"`
			Identifier        []Identifier `json:"identifier"`
			Meta              MetaData     `json:"meta"`
			Text              TextData     `json:"text"`
			Category          Category     `json:"category"`
			Code              Note         `json:"code"`
			PresentedForm     []Attachment `json:"presentedForm"`
			Request           []Thing      `json:"request"`
			Result            []Thing      `json:"result"`
		} `json:"resource"`
	} `json:"entry"`
}

// type ResourceEntry struct {
// 	EntryPartial EntryPartial
// 	Resource     DiagnosticReportResource `json:"resource"`
// }

// type DiagnosticReportResource struct {
// 	ResourcePartial   ResourcePartial
// 	EffectiveDateTime time.Time    `json:"effective_date_time"`
// 	Issued            time.Time    `json:"issued"`
// 	Identifier        []Identifier `json:"identifier"`
// 	Meta              MetaData     `json:"meta"`
// 	Text              TextData     `json:"text"`
// 	Category          CodeText     `json:"category"`
// 	Code              CodeText     `json:"code"`
// 	PresentedForm     []Attachment `json:"presentedForm"`
// 	Request           []Thing      `json:"request"`
// 	Encounter         Encounter    `json:"encounter"`
// 	Result            []Thing      `json:"result"`
// }

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
