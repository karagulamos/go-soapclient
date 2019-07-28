package models

import "encoding/xml"

// ListOfContinentsByNameRequest ...
type ListOfContinentsByNameRequest struct {
	// for requests always use the form `xml:"<namespace> <soapAction>"` to avoid unintended side-effects
	XMLName xml.Name `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfContinentsByName"`
}

// Continent ...
type Continent struct {
	Code string `xml:"sCode"`
	Name string `xml:"sName"`
}

// ListOfContinentsByNameResponse ...
type ListOfContinentsByNameResponse struct {
	XMLName    xml.Name    `xml:"Envelope"`
	Continents []Continent `xml:"Body>ListOfContinentsByNameResponse>ListOfContinentsByNameResult>tContinent"`
}

// OR you could use the longer form
//
// type ListOfContinentsByNameResponseEnvelope struct {
// 	XMLName xml.Name `xml:"Envelope"`

// 	Body struct {
// 		Response struct {
// 			XMLName    xml.Name    `xml:"ListOfContinentsByNameResponse"`
// 			Continents []Continent `xml:"ListOfContinentsByNameResult>tContinent"`
// 		}
// 	}
// }
//
// Usage:
//
// envelope := new(models.ListOfContinentsByNameResponseEnvelope)
// ...
// continents := envelope.Body.Response.Continents
