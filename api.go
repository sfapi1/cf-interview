package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type APIData struct {
	Properties APIProperties `json:"properties"`
}

type APIProperties struct {
	Updated string      `json:"updated"`
	Units   string      `json:"units"`
	Periods []APIPeriod `json:"periods"`
}

type APIPeriod struct {
	Number           uint8        `json:"number"`
	Name             string       `json:"name"`
	IsDaytime        bool         `json:"isDaytime"`
	Temp             int16        `json:"temperature"`
	TempUnit         string       `json:"temperatureUnit"`
	PrecipChance     APIUnitInt   `json:"probabilityOfPrecipitation"`
	Dewpoint         APIUnitFloat `json:"dewpoint"`
	Humidity         APIUnitFloat `json:"relativeHumidity"`
	WindSpeed        string       `json:"windSpeed"`
	WindDir          string       `json:"windDirection"`
	ShortForecast    string       `json:"shortForecast"`
	DetailedForecast string       `json:"detailedForecast"`
}

type APIUnitInt struct {
	UnitCode string `json:"unitCode"`
	Value    int8   `json:"value"`
}

type APIUnitFloat struct {
	UnitCode string  `json:"unitCode"`
	Value    float32 `json:"value"`
}

// requests forecast data from weather API endpoint for specific region
// input: string region in format STO/48,75
// returns the response body as []byte
func getForecast(region string) ([]byte, error) {
	// URL for weather API
	url := "https://api.weather.gov/gridpoints/" + region + "/forecast"

	// get url and handle errors
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Response status code: " + strconv.Itoa(resp.StatusCode))
	}

	// read body of response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Unmarshal JSON []byte response body from weather API into APIData struct
// keeps same hierarchy as JSON returned from API
func unmarshalForecast(forecastBytes []byte) (*APIData, error) {
	var apidata APIData
	err := json.Unmarshal(forecastBytes, &apidata)
	if err != nil {
		return nil, err
	}

	return &apidata, nil
}

// prints the given APIPeriod structs to stdout in a semi-nice format
func displayForecastPeriods(periods []APIPeriod) {
	for _, period := range periods {
		fmt.Printf("%v: %v | %v°%v | %v%% Humidity\n", period.Name, period.ShortForecast, period.Temp, period.TempUnit, period.Humidity.Value)
	}
}
