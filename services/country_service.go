package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"travelsphere/models"
)

const restCountriesBase = "https://restcountries.com/v3.1"

// transformCountries converts raw REST Countries API response into []models.Country
func transformCountries(raw []map[string]interface{}) []models.Country {
	var countries []models.Country

	for _, item := range raw {
		country := models.Country{}

		if nameObj, ok := item["name"].(map[string]interface{}); ok {
			country.Name, _ = nameObj["common"].(string)
		}

		if capitals, ok := item["capital"].([]interface{}); ok && len(capitals) > 0 {
			country.Capital, _ = capitals[0].(string)
		}

		if pop, ok := item["population"].(float64); ok {
			country.Population = int64(pop)
		}

		if region, ok := item["region"].(string); ok {
			country.Region = region
		}

		if subregion, ok := item["subregion"].(string); ok {
			country.Subregion = subregion
		}

		if flags, ok := item["flags"].(map[string]interface{}); ok {
			country.FlagURL, _ = flags["svg"].(string)
		}

		if currencies, ok := item["currencies"].(map[string]interface{}); ok {
			for code, val := range currencies {
				if currObj, ok := val.(map[string]interface{}); ok {
					name, _ := currObj["name"].(string)
					country.Currency = code + " (" + name + ")"
				}
				break
			}
		}

		if langs, ok := item["languages"].(map[string]interface{}); ok {
			var langList []string
			for _, v := range langs {
				if s, ok := v.(string); ok {
					langList = append(langList, s)
				}
			}
			country.Languages = strings.Join(langList, ", ")
		}

		if latlng, ok := item["latlng"].([]interface{}); ok && len(latlng) >= 2 {
			country.Lat, _ = latlng[0].(float64)
			country.Lon, _ = latlng[1].(float64)
		}

		country.Slug = strings.ToLower(strings.ReplaceAll(country.Name, " ", "-"))
		countries = append(countries, country)
	}

	return countries
}


func fetchAndTransform(url string) ([]models.Country, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API status: %d", resp.StatusCode)
	}

	var raw []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}
	return transformCountries(raw), nil
}


func GetFeaturedCountries() ([]models.Country, error) {
	
	names := []string{"united states", "france", "japan", "australia", "brazil", "bangladesh"}
	var featured []models.Country
	for _, name := range names {
		url := restCountriesBase + "/name/" + strings.ReplaceAll(name, " ", "%20") +
			"?fields=name,capital,population,flags,currencies,languages,region,subregion,latlng"
		countries, err := fetchAndTransform(url)
		if err != nil || len(countries) == 0 {
			continue
		}
		featured = append(featured, countries[0])
	}
	return featured, nil
}


func GetAllCountries() ([]models.Country, error) {
	url := restCountriesBase + "/all?fields=name,capital,population,flags,currencies,languages,region,subregion,latlng"
	return fetchAndTransform(url)
}


func SearchCountries(search, region string) ([]models.Country, error) {
	all, err := GetAllCountries()
	if err != nil {
		return nil, err
	}

	search = strings.ToLower(search)
	var result []models.Country

	for _, c := range all {
		if region != "" && !strings.EqualFold(c.Region, region) {
			continue
		}
		if search != "" {
			if !strings.Contains(strings.ToLower(c.Name), search) &&
				!strings.Contains(strings.ToLower(c.Capital), search) {
				continue
			}
		}
		result = append(result, c)
	}
	return result, nil
}


func GetCountryBySlug(slug string) (*models.Country, error) {
	name := strings.ReplaceAll(slug, "-", " ")
	url := restCountriesBase + "/name/" + strings.ReplaceAll(name, " ", "%20") +
		"?fields=name,capital,population,flags,currencies,languages,region,subregion,latlng"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("country not found")
	}

	var raw []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}
	if len(raw) == 0 {
		return nil, fmt.Errorf("country not found")
	}

	countries := transformCountries(raw)
	return &countries[0], nil
}
