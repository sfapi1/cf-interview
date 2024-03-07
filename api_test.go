package main

import (
	"reflect"
	"testing"
)

func TestUnmarshalForecast(t *testing.T) {

	// declare JSON test data (simulated from API)
	testData := []byte(`{
		"properties": {
			"updated": "2024-02-25T16:02:15+00:00",
			"units": "us",
			"periods": [
				{
					"number": 1,
					"name": "Monday Night",
					"startTime": "2024-02-26T18:00:00-08:00",
					"endTime": "2024-02-27T06:00:00-08:00",
					"isDaytime": false,
					"temperature": 39,
					"temperatureUnit": "F",
					"probabilityOfPrecipitation": {
						"unitCode": "wmoUnit:percent",
						"value": 20
					},
					"dewpoint": {
						"unitCode": "wmoUnit:degC",
						"value": 9.4
					},
					"relativeHumidity": {
						"unitCode": "wmoUnit:percent",
						"value": 85
					},
					"windSpeed": "7 to 12 mph",
					"windDirection": "W",
					"shortForecast": "Slight Chance Rain Showers then Partly Cloudy",
					"detailedForecast": "A slight chance of rain showers before 10pm. Partly cloudy, with a low around 39. West wind 7 to 12 mph, with gusts as high as 18 mph. Chance of precipitation is 20%."
				}
			]
		}
	}`)

	// declare expected unmarshaled data into an APIData struct
	expectedAPIData := APIData{
		Properties: APIProperties{
			Updated: "2024-02-25T16:02:15+00:00",
			Units:   "us",
			Periods: []APIPeriod{
				{
					Number:           1,
					Name:             "Monday Night",
					IsDaytime:        false,
					Temp:             39,
					TempUnit:         "F",
					PrecipChance:     APIUnitInt{UnitCode: "wmoUnit:percent", Value: 20},
					Dewpoint:         APIUnitFloat{UnitCode: "wmoUnit:degC", Value: 9.4},
					Humidity:         APIUnitFloat{UnitCode: "wmoUnit:percent", Value: 85},
					WindSpeed:        "7 to 12 mph",
					WindDir:          "W",
					ShortForecast:    "Slight Chance Rain Showers then Partly Cloudy",
					DetailedForecast: "A slight chance of rain showers before 10pm. Partly cloudy, with a low around 39. West wind 7 to 12 mph, with gusts as high as 18 mph. Chance of precipitation is 20%.",
				},
			},
		},
	}

	// run test and store result
	result, err := unmarshalForecast(testData)

	// if error during operation
	if err != nil {
		t.Errorf("Error encountered while unmarshalling JSON: %v", err)
	}

	// if mismatched result
	if !reflect.DeepEqual(*result, expectedAPIData) {
		t.Errorf("Mismatch: got %+v\n\nexpected %+v", result, expectedAPIData)
	}

}
