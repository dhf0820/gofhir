package fhir

import (
	//log "github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"

	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

//const pid = "Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB"

//const ordercode = "8310-5"
//const baseurl = "https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/"

const pid = "4342009"
const baseurl = "https://fhir-open.sandboxcerner.com/dstu2/0b8a0111-e8e6-4c26-a91c-5069cbc6b1ca/"

// func TestDevice(t *testing.T) {
// 	c := New(baseurl)
// 	data, err := c.GetDevice(pid)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if data.Total == 0 {
// 		t.Error("Expected > 0 got 0")
// 	}
// }

func TestPatSearch(t *testing.T) {
	c := New(baseurl)
	startTime := time.Now()
	data, err := c.PatientSearch("family=Argonaut&given=Jason")
	//data, err := c.PatientSearch("family=smart&given=nancy")
	log.Infof("Fhir Patient Search took %s", time.Since(startTime))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Patient entry: %s\n", spew.Sdump(data.Entry))
	log.Infof("Number of Patients returned: %d, total: %d, results %d", len(data.Entry), data.Total, data.SearchResult.Total)
	log.Infof("Total PatientsL %d", data.SearchResult.Total)
	if data.SearchResult.Total == 0 {
		t.Error("Expected > 0 got 0")
	}

	//fmt.Printf("Patient: %s\n", spew.Sdump(data.SearchResult))
}

func TestPatient(t *testing.T) {
	c := New(baseurl)
	data, err := c.GetPatient(pid)
	if err != nil {
		t.Fatal(err)
	}
	if len(data.Name) == 0 {
		t.Error("Expected > 0 got 0")
	}
}

func TestPatientDiagnosticReport(t *testing.T) {
	c := New(baseurl)

	//https://fhir-open.sandboxcerner.com/dstu2/0b8a0111-e8e6-4c26-a91c-5069cbc6b1ca/DiagnosticReport?patient=1316020&_count=10
	data, err := c.GetPatientDiagnosticReport("1316020")
	//data, err := c.GetDocumentReference(pid)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Document: %s\n", spew.Sdump(data))
	// if data.Total == 0 {
	// 	t.Error("Expected > 0 got 0")
	// }
}
func TestDocument(t *testing.T) {
	c := New(baseurl)
	data, err := c.GetDocumentReference(pid)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Document: %s\n", data)
	// if data.Total == 0 {
	// 	t.Error("Expected > 0 got 0")
	// }
}

// func TestCondition(t *testing.T) {
// 	c := New(baseurl)
// 	data, err := c.GetCondition(pid)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if data.Total == 0 {
// 		t.Error("Expected > 0 got 0")
// 	}
// }

// func TestProcedure(t *testing.T) {
// 	c := New(baseurl)
// 	data, err := c.GetProcedure(pid)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if data.Total == 0 {
// 		t.Error("Expected > 0 got 0")
// 	}
// }

// func TestMedication(t *testing.T) {
// 	c := New(baseurl)
// 	data, err := c.GetMedication(pid)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if data.Total == 0 {
// 		t.Error("Expected > 0 got 0")
// 	}
// }

// func TestObservation(t *testing.T) {
// 	c := New(baseurl)
// 	data, err := c.GetObservation(pid, ordercode)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if data.Total == 0 {
// 		t.Error("Expected > 0 got 0")
// 	}
// }

// func TestImmunization(t *testing.T) {
// 	c := New(baseurl)
// 	data, err := c.GetImmunization(pid)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if data.Total == 0 {
// 		t.Error("Expected > 0 got 0")
// 	}
// }

// func TestAllergy(t *testing.T) {
// 	c := New(baseurl)
// 	data, err := c.GetAllergyIntolerence(pid)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if data.Total == 0 {
// 		t.Error("Expected > 0 got 0")
// 	}
// }

// func TestFamilyHx(t *testing.T) {
// 	c := New(baseurl)
// 	data, err := c.GetFamilyMemberHistory(pid)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if data.Total == 0 {
// 		t.Error("Expected > 0 got 0")
// 	}
// }
