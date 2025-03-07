package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"tinh/sun/pipelines"
	"tinh/sun/utils"
	"tinh/sun/weather"

	ics "github.com/arran4/golang-ical"
)

func getCurrentWeather() {
	apiKey := os.Getenv("WEATHERSTACK_API_KEY")
	if apiKey == "" {
		fmt.Println("WEATHERSTACK_API_KEY is not set")
		return
	}

	baseURL := "http://api.weatherstack.com/current"
	query := "New York"
	url := fmt.Sprintf("%s?access_key=%s&query=%s", baseURL, apiKey, query)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var prettyBody = utils.PrettyPrint(body)
	if prettyBody == nil {
		return
	}
	fmt.Println("Response:", string(prettyBody))
}

func printEventProperty(event *ics.VEvent, property ics.ComponentProperty) {
	prop := event.GetProperty(property)
	if prop != nil {
		fmt.Printf("%s: %s\n", property, prop.Value)
	} else {
		fmt.Printf("%s: None\n", property)
	}
}

func cal() {
	// Open the ICS file
	file, err := os.Open("/Users/tinhnguyen/workspace/sun/cta_event_cal.ics")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Parse the ICS file
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	cal, err := ics.ParseCalendar(strings.NewReader(string(content)))

	if err != nil {
		fmt.Println("Error parsing calendar:", err)
		return
	}

	// Process the events in the calendar
	for _, event := range cal.Events() {
		printEventProperty(event, ics.ComponentPropertySummary)
		printEventProperty(event, ics.ComponentPropertyDtStart)
		printEventProperty(event, ics.ComponentPropertyDtEnd)
		printEventProperty(event, ics.ComponentPropertyStatus)
		printEventProperty(event, ics.ComponentPropertyLocation)
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [command]")
		fmt.Println("Available commands: weather, currentWeather, cal, init-db, scrape")
		return
	}

	command := os.Args[1]

	switch command {
	case "weather":
		weather.WeatherForecast()
	case "currentWeather":
		getCurrentWeather()
	case "cal":
		cal()
	case "init-db":
		pipelines.Init()
	case "scrape":
		pipelines.Scrape()
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: weather, currentWeather, cal")
	}
}
