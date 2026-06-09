package services

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestTransformCountriesMapsFieldsAndSlug(t *testing.T) {
	raw := []map[string]interface{}{
		{
			"name":       map[string]interface{}{"common": "Bangladesh"},
			"capital":    []interface{}{"Dhaka"},
			"population": float64(170000000),
			"region":     "Asia",
			"subregion":  "Southern Asia",
			"flags":      map[string]interface{}{"svg": "https://flag.example/bd.svg"},
			"currencies": map[string]interface{}{
				"BDT": map[string]interface{}{"name": "Taka"},
			},
			"languages": map[string]interface{}{"bn": "Bengali", "en": "English"},
			"latlng":    []interface{}{23.81, 90.41},
		},
	}

	countries := transformCountries(raw)
	if len(countries) != 1 {
		t.Fatalf("expected 1 country, got %d", len(countries))
	}

	country := countries[0]
	if country.Name != "Bangladesh" {
		t.Fatalf("expected country name Bangladesh, got %q", country.Name)
	}
	if country.Capital != "Dhaka" {
		t.Fatalf("expected capital Dhaka, got %q", country.Capital)
	}
	if country.Population != 170000000 {
		t.Fatalf("expected population 170000000, got %d", country.Population)
	}
	if country.Currency != "BDT (Taka)" {
		t.Fatalf("expected currency BDT (Taka), got %q", country.Currency)
	}
	if country.Languages != "Bengali, English" {
		t.Fatalf("expected languages to be joined, got %q", country.Languages)
	}
	if country.Slug != "bangladesh" {
		t.Fatalf("expected slug bangladesh, got %q", country.Slug)
	}
}

func TestGetWeatherWithoutApiKeyReturnsNil(t *testing.T) {
	t.Setenv("WEATHER_API_KEY", "")

	weather, err := GetWeather("Dhaka")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if weather != nil {
		t.Fatalf("expected nil weather when API key is missing, got %+v", weather)
	}
}

func TestGetAttractionsWithoutApiKeyReturnsError(t *testing.T) {
	t.Setenv("OPENTRIPMAP_KEY", "")

	attractions, err := GetAttractions(23.81, 90.41, 5)
	if err == nil {
		t.Fatal("expected error when OpenTripMap key is missing")
	}
	if attractions != nil {
		t.Fatalf("expected nil attractions on error, got %+v", attractions)
	}
}

func TestGetAllCountriesUsesInjectedHTTP(t *testing.T) {
	oldHTTP := httpGet
	defer func() { httpGet = oldHTTP }()

	httpGet = func(url string) (*http.Response, error) {
		if !strings.Contains(url, "/all") {
			t.Fatalf("expected all countries URL, got %s", url)
		}

		body := []map[string]interface{}{{
			"name":       map[string]interface{}{"common": "France"},
			"capital":    []interface{}{"Paris"},
			"population": float64(67000000),
			"region":     "Europe",
			"subregion":  "Western Europe",
			"flags":      map[string]interface{}{"svg": "https://flag.example/fr.svg"},
			"currencies": map[string]interface{}{"EUR": map[string]interface{}{"name": "Euro"}},
			"languages":  map[string]interface{}{"fr": "French"},
			"latlng":     []interface{}{48.86, 2.35},
		}}
		data, _ := json.Marshal(body)
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(string(data))),
		}, nil
	}

	countries, err := GetAllCountries()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(countries) != 1 || countries[0].Name != "France" {
		t.Fatalf("expected France country, got %+v", countries)
	}
}

func TestGetWeatherUsesInjectedKeyAndHTTP(t *testing.T) {
	oldHTTP := httpGet
	oldEnv := getenv
	defer func() {
		httpGet = oldHTTP
		getenv = oldEnv
	}()

	httpGet = func(url string) (*http.Response, error) {
		if !strings.Contains(url, "api.weatherapi.com") {
			t.Fatalf("expected weather URL, got %s", url)
		}
		body := `{"location":{"name":"Dhaka"},"current":{"temp_c":32.5,"humidity":70,"wind_kph":12.5,"condition":{"text":"Sunny","icon":"//cdn.weatherapi.com/weather/64x64/day/113.png"}}}`
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(body)),
		}, nil
	}
	getenv = func(key string) string {
		if key == "WEATHER_API_KEY" {
			return "secret-key"
		}
		return ""
	}

	weather, err := GetWeather("Dhaka")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if weather == nil {
		t.Fatal("expected weather result")
	}
	if weather.City != "Dhaka" || weather.Icon != "https://cdn.weatherapi.com/weather/64x64/day/113.png" {
		t.Fatalf("expected weather and icon to be parsed, got %+v", weather)
	}
}

func TestGetAttractionsUsesInjectedKeyAndHTTP(t *testing.T) {
	oldHTTP := httpGet
	oldEnv := getenv
	defer func() {
		httpGet = oldHTTP
		getenv = oldEnv
	}()

	httpGet = func(url string) (*http.Response, error) {
		if !strings.Contains(url, "api.opentripmap.com") {
			t.Fatalf("expected OpenTripMap URL, got %s", url)
		}
		body := `{"features":[{"properties":{"name":"Lalbagh Fort","kinds":"forts","xid":"1"}}]}`
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(body)),
		}, nil
	}
	getenv = func(key string) string {
		if key == "OPENTRIPMAP_KEY" {
			return "secret-key"
		}
		return ""
	}

	attractions, err := GetAttractions(23.81, 90.41, 5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(attractions) != 1 || attractions[0].Name != "Lalbagh Fort" {
		t.Fatalf("expected one attraction, got %+v", attractions)
	}
}
