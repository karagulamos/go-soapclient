package main

import (
	"log"

	"github.com/karagulamos/soapclient"
	"github.com/karagulamos/soapclient/examples/dataflex/models"
)

func main() {
	client := soapclient.MakePOST(`http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso`)

	request := new(models.ListOfContinentsByNameRequest)
	response := new(models.ListOfContinentsByNameResponse)

	err := client.WithRequest(request).Fetch(response)

	if err != nil {
		log.Fatalf(err.Error())
	}

	continents := response.Continents

	for _, continent := range continents {
		log.Printf("%s (%s)", continent.Name, continent.Code)
	}
}
