package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

// Response of Google Maps request
type Response struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Bounds struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID string   `json:"place_id"`
		Types   []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

// GeoPoint : this is a geographic point
type GeoPoint struct {
	latitude  float64
	longitude float64
}

var geoPointType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GeoPoint",
		Fields: graphql.Fields{
			"latitude": &graphql.Field{
				Type: graphql.Float,
			},
			"longitude": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// Country : this is a country
type Country struct {
	Name string `json:"name"`
}

var countryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Country",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var myClient = &http.Client{Timeout: 10 * time.Second}

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"point": &graphql.Field{
				Type: countryType,
				Args: graphql.FieldConfigArgument{
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					}, "longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					point := GeoPoint{
						latitude:  p.Args["latitude"].(float64),
						longitude: p.Args["longitude"].(float64),
					}

					return getCountry(point)
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
