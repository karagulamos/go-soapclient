package models

import "encoding/xml"

// AddRequest ...
type AddRequest struct {
	// for requests always use the form `xml:"<namespace> <soapAction>"` to avoid unintended side-effects
	XMLName xml.Name `xml:"http://tempuri.org/ Add"`
	A       int      `xml:"intA"`
	B       int      `xml:"intB"`
}

// AddResponse ...
type AddResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Sum     int      `xml:"Body>AddResponse>AddResult"`
}
