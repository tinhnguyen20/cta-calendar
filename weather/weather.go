package weather

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"tinh/sun/utils"
)

func WeatherForecast() {
	apiKey := os.Getenv("WEATHERSTACK_API_KEY")
	if apiKey == "" {
		fmt.Println("WEATHERSTACK_API_KEY is not set")
		return
	}

	baseURL := "http://api.weatherstack.com/forecast"
	query := "New York"
	forecastDays := "7"
	hourly := "1"
	interval := "3"
	units := "m"
	language := "en"
	callback := "MY_CALLBACK"

	url := fmt.Sprintf("%s?access_key=%s&query=%s&forecast_days=%s&hourly=%s&interval=%s&units=%s&language=%s&callback=%s",
		baseURL, apiKey, query, forecastDays, hourly, interval, units, language, callback)

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
