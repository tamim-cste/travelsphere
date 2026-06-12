package services

import (
	"strings"
	"testing"
)

func TestTransformCountriesEmpty(t *testing.T) {
	result := transformCountries([]map[string]interface{}{})
	if len(result) != 0 {
		t.Errorf("expected empty, got %d", len(result))
	}
}

func TestTransformCountriesV5Fields(t *testing.T) {
	raw := []map[string]interface{}{
		{
			"names":      map[string]interface{}{"common": "Bangladesh"},
			"capitals":   []interface{}{map[string]interface{}{"name": "Dhaka"}},
			"population": float64(175423000),
			"region":     "Asia",
			"subregion":  "Southern Asia",
			"flag":       map[string]interface{}{"url_svg": "https://flags.example.com/bd.svg"},
			"currencies": []interface{}{map[string]interface{}{"code": "BDT", "name": "Bangladeshi taka"}},
			"languages":  []interface{}{map[string]interface{}{"name": "Bengali"}},
			"coordinates": map[string]interface{}{"lat": float64(24.0), "lng": float64(90.0)},
		},
	}
	countries := transformCountries(raw)
	if len(countries) != 1 {
		t.Fatalf("expected 1, got %d", len(countries))
	}
	c := countries[0]
	if c.Name != "Bangladesh" {
		t.Errorf("expected Bangladesh, got %s", c.Name)
	}
	if c.Capital != "Dhaka" {
		t.Errorf("expected Dhaka, got %s", c.Capital)
	}
	if c.Population != 175423000 {
		t.Errorf("expected 175423000, got %d", c.Population)
	}
	if c.Region != "Asia" {
		t.Errorf("expected Asia, got %s", c.Region)
	}
	if c.Subregion != "Southern Asia" {
		t.Errorf("expected Southern Asia, got %s", c.Subregion)
	}
	if c.Slug != "bangladesh" {
		t.Errorf("expected slug bangladesh, got %s", c.Slug)
	}
	if !strings.Contains(c.Currency, "BDT") {
		t.Errorf("expected BDT in currency, got %s", c.Currency)
	}
	if c.Languages != "Bengali" {
		t.Errorf("expected Bengali, got %s", c.Languages)
	}
	if c.Lat != 24.0 || c.Lon != 90.0 {
		t.Errorf("expected lat=24, lon=90, got %f, %f", c.Lat, c.Lon)
	}
	if c.FlagURL != "https://flags.example.com/bd.svg" {
		t.Errorf("expected flag URL, got %s", c.FlagURL)
	}
}

func TestTransformCountriesSlug(t *testing.T) {
	raw := []map[string]interface{}{
		{"names": map[string]interface{}{"common": "United States"}},
	}
	c := transformCountries(raw)
	if c[0].Slug != "united-states" {
		t.Errorf("expected united-states, got %s", c[0].Slug)
	}
}

func TestTransformCountriesNoCapital(t *testing.T) {
	raw := []map[string]interface{}{
		{"names": map[string]interface{}{"common": "Antarctica"}},
	}
	c := transformCountries(raw)
	if c[0].Capital != "" {
		t.Errorf("expected empty capital, got %s", c[0].Capital)
	}
}

func TestTransformCountriesMultipleLanguages(t *testing.T) {
	raw := []map[string]interface{}{
		{
			"names": map[string]interface{}{"common": "Switzerland"},
			"languages": []interface{}{
				map[string]interface{}{"name": "French"},
				map[string]interface{}{"name": "German"},
				map[string]interface{}{"name": "Italian"},
			},
		},
	}
	c := transformCountries(raw)
	if !strings.Contains(c[0].Languages, "French") {
		t.Errorf("expected French in languages, got %s", c[0].Languages)
	}
}

func TestTransformCountriesSkipsEmptyName(t *testing.T) {
	raw := []map[string]interface{}{
		{"population": float64(100)}, // no names field
	}
	result := transformCountries(raw)
	if len(result) != 0 {
		t.Errorf("expected empty result for item without name, got %d", len(result))
	}
}

func TestTransformCountriesMultiple(t *testing.T) {
	raw := []map[string]interface{}{
		{"names": map[string]interface{}{"common": "France"}},
		{"names": map[string]interface{}{"common": "Japan"}},
		{"names": map[string]interface{}{"common": "Brazil"}},
	}
	result := transformCountries(raw)
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func TestDoGetNoKey(t *testing.T) {
	origGetenv := getenv
	getenv = func(key string) string { return "" }
	defer func() { getenv = origGetenv }()

	_, err := doGet("https://example.com")
	if err == nil {
		t.Error("expected error when key missing")
	}
	if !strings.Contains(err.Error(), "RESTCOUNTRIES_KEY") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestRestCountriesKeyReadsFromEnv(t *testing.T) {
	origGetenv := getenv
	getenv = func(key string) string {
		if key == "RESTCOUNTRIES_KEY" {
			return "test-key"
		}
		return ""
	}
	defer func() { getenv = origGetenv }()

	if restCountriesKey() != "test-key" {
		t.Error("expected test-key")
	}
}

func TestRestCountriesKeyEmptyWhenNotSet(t *testing.T) {
	origGetenv := getenv
	getenv = func(key string) string { return "" }
	defer func() { getenv = origGetenv }()

	if restCountriesKey() != "" {
		t.Error("expected empty key")
	}
}

func TestRegionFilterLogic(t *testing.T) {
	countries := transformCountries([]map[string]interface{}{
		{"names": map[string]interface{}{"common": "Bangladesh"}, "region": "Asia"},
		{"names": map[string]interface{}{"common": "France"}, "region": "Europe"},
		{"names": map[string]interface{}{"common": "Japan"}, "region": "Asia"},
	})

	var asia []string
	for _, c := range countries {
		if strings.EqualFold(c.Region, "Asia") {
			asia = append(asia, c.Name)
		}
	}
	if len(asia) != 2 {
		t.Errorf("expected 2 Asia countries, got %d", len(asia))
	}
}
