package services

import (
	"strings"
	"testing"
)

func TestTransformCountriesEmpty(t *testing.T) {
	result := transformCountries([]map[string]interface{}{})
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d", len(result))
	}
}

func TestTransformCountriesFullFields(t *testing.T) {
	raw := []map[string]interface{}{
		{
			"name":       map[string]interface{}{"common": "Bangladesh"},
			"capital":    []interface{}{"Dhaka"},
			"population": float64(170000000),
			"region":     "Asia",
			"subregion":  "Southern Asia",
			"flags":      map[string]interface{}{"svg": "https://flag.example.com/bd.svg"},
			"currencies": map[string]interface{}{
				"BDT": map[string]interface{}{"name": "Bangladeshi taka"},
			},
			"languages": map[string]interface{}{"ben": "Bengali"},
			"latlng":    []interface{}{float64(23.7), float64(90.4)},
		},
	}
	countries := transformCountries(raw)
	if len(countries) != 1 {
		t.Fatalf("expected 1 country, got %d", len(countries))
	}
	c := countries[0]
	if c.Name != "Bangladesh" {
		t.Errorf("expected Bangladesh, got %s", c.Name)
	}
	if c.Capital != "Dhaka" {
		t.Errorf("expected Dhaka, got %s", c.Capital)
	}
	if c.Population != 170000000 {
		t.Errorf("expected 170000000, got %d", c.Population)
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
	if c.Lat != 23.7 {
		t.Errorf("expected lat 23.7, got %f", c.Lat)
	}
	if c.Lon != 90.4 {
		t.Errorf("expected lon 90.4, got %f", c.Lon)
	}
	if !strings.Contains(c.Currency, "BDT") {
		t.Errorf("expected currency to contain BDT, got %s", c.Currency)
	}
	if c.Languages != "Bengali" {
		t.Errorf("expected Bengali, got %s", c.Languages)
	}
}

func TestTransformCountriesSlugSpaces(t *testing.T) {
	raw := []map[string]interface{}{
		{"name": map[string]interface{}{"common": "United States"}},
	}
	c := transformCountries(raw)
	if c[0].Slug != "united-states" {
		t.Errorf("expected united-states, got %s", c[0].Slug)
	}
}

func TestTransformCountriesNoCapital(t *testing.T) {
	raw := []map[string]interface{}{
		{"name": map[string]interface{}{"common": "Antarctica"}},
	}
	c := transformCountries(raw)
	if c[0].Capital != "" {
		t.Errorf("expected empty capital, got %s", c[0].Capital)
	}
}

func TestTransformCountriesMultipleLanguages(t *testing.T) {
	raw := []map[string]interface{}{
		{
			"name":      map[string]interface{}{"common": "Switzerland"},
			"languages": map[string]interface{}{"fra": "French", "deu": "German", "ita": "Italian"},
		},
	}
	c := transformCountries(raw)
	if !strings.Contains(c[0].Languages, "French") {
		t.Errorf("expected French in languages, got %s", c[0].Languages)
	}
}

func TestTransformCountriesNoFlags(t *testing.T) {
	raw := []map[string]interface{}{
		{"name": map[string]interface{}{"common": "Test"}},
	}
	c := transformCountries(raw)
	if c[0].FlagURL != "" {
		t.Errorf("expected empty flag URL, got %s", c[0].FlagURL)
	}
}

func TestSearchCountriesFilterLogic(t *testing.T) {
	// Test the filter logic directly using transformCountries output
	raw := []map[string]interface{}{
		{"name": map[string]interface{}{"common": "Bangladesh"}, "capital": []interface{}{"Dhaka"}, "region": "Asia"},
		{"name": map[string]interface{}{"common": "France"}, "capital": []interface{}{"Paris"}, "region": "Europe"},
		{"name": map[string]interface{}{"common": "Japan"}, "capital": []interface{}{"Tokyo"}, "region": "Asia"},
	}
	all := transformCountries(raw)

	// filter by region Asia
	var asiaOnly []string
	for _, c := range all {
		if strings.EqualFold(c.Region, "Asia") {
			asiaOnly = append(asiaOnly, c.Name)
		}
	}
	if len(asiaOnly) != 2 {
		t.Errorf("expected 2 Asia countries, got %d", len(asiaOnly))
	}

	// filter by name search
	var found []string
	for _, c := range all {
		if strings.Contains(strings.ToLower(c.Name), "ban") {
			found = append(found, c.Name)
		}
	}
	if len(found) != 1 || found[0] != "Bangladesh" {
		t.Errorf("expected Bangladesh in search results, got %v", found)
	}
}
