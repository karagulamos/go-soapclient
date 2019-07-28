package main

import (
	"log"

	"github.com/karagulamos/soapclient"
	"github.com/karagulamos/soapclient/examples/weather/models"
)

func main() {
	client := soapclient.MakePOST(`https://graphical.weather.gov:443/xml/SOAP_server/ndfdXMLserver.php`)

	request := &models.CityListRequest{DisplayLevel: "1"}
	response := new(models.CityListResponse)

	err := client.WithRequest(request).Fetch(response)

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println(response.XMLResult)
}
