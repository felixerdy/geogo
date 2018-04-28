package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func getCountry(searchPoint GeoPoint) (Country, error) {
	response := Response{}

	lat := strconv.FormatFloat(searchPoint.latitude, 'f', 5, 64)
	lon := strconv.FormatFloat(searchPoint.longitude, 'f', 5, 64)

	googleMapsAPIKey := ""

	res, getErr := http.Get("https://maps.googleapis.com/maps/api/geocode/json?latlng=" + lat + "," + lon + "&result_type=country&key=" + googleMapsAPIKey)
	if getErr != nil {
		fmt.Println(getErr)
	}

	body, parseErr := ioutil.ReadAll(res.Body)
	if parseErr != nil {
		fmt.Println(parseErr)
	}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	myCountry := Country{
		Name: "No results 😕",
	}

	if response.Status != "OK" {
		return myCountry, errors.New("Status is not OK")
	}
	if len(response.Results) == 0 {
		return myCountry, errors.New("No results found")
	}

	myCountry = Country{
		Name: response.Results[0].AddressComponents[0].LongName,
	}
	return myCountry, nil
}
