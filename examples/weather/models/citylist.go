package models

import (
	"encoding/xml"
)

// CityListRequest ...
type CityListRequest struct {
	// for requests always use the form `xml:"<namespace> <soapAction>"` to avoid unintended side-effects
	XMLName      xml.Name `xml:"LatLonListCityNames"`
	DisplayLevel string   `xml:"displayLevel"`
}

// CityListResponse ...
type CityListResponse struct {
	XMLName   xml.Name `xml:"Envelope"`
	XMLResult string   `xml:"Body>LatLonListCityNamesResponse>listLatLonOut"`
}

// OR the long form
//
// type CityListResponseEnvelope struct {
// 	XMLName xml.Name `xml:"Envelope"`
// 	Body    struct {
// 		XMLName  xml.Name `xml:"Body"`
// 		Response struct {
// 			XMLName   xml.Name `xml:"LatLonListCityNamesResponse"`
// 			XMLResult string   `xml:"listLatLonOut"`
// 		}
// 	}
// }
//
// Usage:
//
// envelope := new(models.CityListResponseEnvelope)
// ...
// log.Println(envelope.Body.Response.XMLResult)
