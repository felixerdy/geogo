# Geo Go

This is my first Go app. It serves a web map and returns the country name if you click on one.

## Technology

After a click on the map, the frontend requests the Go backend for the country name through GraphQL. The backend then uses the lat / lon coordinates and makes a request on Google Maps API. **Please insert your Google Maps API Key in the `countryNameFinder.go` file in line 19!**

## Installation

You can either compile and run the .go files by yourself or use the provided Docker files.
