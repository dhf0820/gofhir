package fhir

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	timeout = 3
)

// RetData is the mapped json of the request
type RetData map[string]interface{}

// Connection is a FHIR connection
type Connection struct {
	BaseURL string
	client  *http.Client
}

// New creates a new connection
func New(baseurl string) *Connection {
	return &Connection{
		BaseURL: baseurl,
		client: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   time.Duration(timeout*3) * time.Second,
					KeepAlive: time.Duration(timeout*3) * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   time.Duration(timeout) * time.Second,
				ResponseHeaderTimeout: time.Duration(timeout) * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
	}
}

// Query sends a query to the base url
func (c *Connection) Query(q string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v%v", c.BaseURL, q), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Printf(" !!!fhir query returned err: %s\n", err)
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Query Error: %v\n", err)
			return nil, err
		}
		es := err.Error()
		if strings.Contains(es, "timeout") {
			err = fmt.Errorf("408|%v", err)
		}

		//fmt.Printf("!!!ERROR Response Status Code: %d,  Status: %s\n", resp.StatusCode, string(b))
		// err = &FhirError{
		// 	HttpStatusCode: resp.StatusCode,
		// 	HttpStatus:     resp.Status,
		// 	Message:        string(b),
		// }
		err = fmt.Errorf("%d|%s", resp.StatusCode, string(b))
		return nil, err
	}
	//fmt.Printf("Response Status Code: %s,  Status: %s\n", resp.StatusCode, resp.Status)
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
