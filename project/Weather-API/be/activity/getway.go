package activity

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func formatStringApi(local, startDate, endDate string) string {
	if local != "" && startDate != "" && endDate != "" {
		return fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s?unitGroup=metric&key=X7P79VDYRAJCLYBAQZ85GTBAL&contentType=json", local, startDate, endDate)
	}
	return fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/next7days?unitGroup=metric&key=X7P79VDYRAJCLYBAQZ85GTBAL&contentType=json", local)
}

func fetchWeatherData(apiUrl string) (WeatherData, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiUrl)
	if err != nil {
		return WeatherData{}, err
	}
	defer resp.Body.Close()

	var weatherData WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return WeatherData{}, err
	}

	return weatherData, nil
}

func GetDataWeather(local, startDate, endDate string) WeatherData {
	apiUrl := formatStringApi(local, startDate, endDate)
	weatherData, err := fetchWeatherData(apiUrl)
	if err != nil {
		log.Println("Error fetching weather data:", err)
		return WeatherData{}
	}
	return weatherData
}
