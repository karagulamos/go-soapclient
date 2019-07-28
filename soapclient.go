package soapclient

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/karagulamos/SoapClient/httpclient"
	"golang.org/x/net/html/charset"
)

type soapRequestEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    struct {
		Request interface{}
	}
}

func (s *soapRequestEnvelope) Serialize() ([]byte, error) {
	payload, err := xml.Marshal(s)

	if err != nil {
		return nil, fmt.Errorf("Error serializing envelope: %s", err.Error())
	}

	return payload, nil
}

// SoapClient used to handle SOAP requests
type SoapClient struct {
	httpMethod string
	url        string
	soapAction string
	envelope   soapRequestEnvelope

	headers map[string]string

	httpclient httpclient.HTTPClient
}

// NewSoapClient creates a new instance of the SOAP client
func NewSoapClient(httpclient httpclient.HTTPClient, httpMethod, url, soapAction string) *SoapClient {
	client := new(SoapClient)

	client.httpMethod = httpMethod
	client.url = url
	client.soapAction = soapAction
	client.envelope = soapRequestEnvelope{}
	client.headers = map[string]string{}
	client.httpclient = httpclient

	return client
}

// MakePOST creates a SOAP client instance for POST requests.
//          It takes a url and soapAction as arguments.
func MakePOST(url string, client ...httpclient.HTTPClient) *SoapClient {
	if len(client) > 0 {
		return NewSoapClient(client[0], "POST", url, "")
	}

	return NewSoapClient(httpclient.New(false), "POST", url, "")
}

// MakeGET creates a SOAP client instance for GET requests.
//          It takes a url and soapAction as arguments.
func MakeGET(url string, client ...httpclient.HTTPClient) *SoapClient {
	if len(client) > 0 {
		return NewSoapClient(client[0], "GET", url, "")
	}

	return NewSoapClient(httpclient.New(false), "GET", url, "")
}

// WithRequest accepts the request payload
func (client *SoapClient) WithRequest(request interface{}) *SoapClient {
	client.envelope.Body.Request = request
	return client
}

// WithAction accepts the soap action for the request
func (client *SoapClient) WithAction(soapAction string) *SoapClient {
	client.soapAction = soapAction
	return client
}

// SetHeader sets or adds a header to the request
func (client *SoapClient) SetHeader(key, value string) {
	client.headers[key] = value
}

// Fetch executes the SOAP requests and fetches the result
func (client *SoapClient) Fetch(result interface{}) error {
	payload, err := client.envelope.Serialize()

	if err != nil {
		return err
	}

	request, err := http.NewRequest(client.httpMethod, client.url, bytes.NewReader(payload))

	if err != nil {
		return fmt.Errorf("Parse error: %s", err.Error())
	}

	request.Header.Set("Content-type", "text/xml; charset=\"UTF-8\"")

	if client.soapAction != "" {
		request.Header.Set("SOAPAction", client.soapAction)
	}

	for key, value := range client.headers {
		request.Header.Set(key, value)
	}

	response, err := client.httpclient.Do(request)

	if err != nil {
		return fmt.Errorf("Error sending request: %s", err.Error())
	}

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(response.Body)
	// log.Printf(buf.String())

	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&result)

	if err != nil {
		return fmt.Errorf("Error deserializing XML: %s", err.Error())
	}

	return nil
}
