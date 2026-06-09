package services

import (
	"os"
	"testing"
)

func TestGetWeatherNoAPIKey(t *testing.T) {
	
	os.Unsetenv("WEATHER_API_KEY")
	weather, err := GetWeather("Dhaka")
	if err != nil {
		t.Errorf("expected nil error when key missing, got: %v", err)
	}
	if weather != nil {
		t.Error("expected nil weather when key missing")
	}
}

func TestGetWeatherEmptyCityNoKey(t *testing.T) {
	os.Unsetenv("WEATHER_API_KEY")
	weather, err := GetWeather("")
	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}
	if weather != nil {
		t.Error("expected nil weather for empty city with no key")
	}
}
