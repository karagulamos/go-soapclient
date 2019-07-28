package main

import (
	"log"

	"github.com/karagulamos/soapclient"
	"github.com/karagulamos/soapclient/examples/calculator/models"
)

func main() {
	client := soapclient.MakePOST(`http://www.dneonline.com/calculator.asmx`)

	request := &models.AddRequest{A: 5, B: 6}
	response := new(models.AddResponse)

	err := client.WithRequest(request).Fetch(response)

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("%d + %d = %d", request.A, request.B, response.Sum)
}
