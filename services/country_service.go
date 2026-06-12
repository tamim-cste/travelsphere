package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"travelsphere/models"
)

var (
	httpGet = http.Get
	getenv  = os.Getenv
)

const restCountriesBase = "https://api.restcountries.com/countries/v5"

func restCountriesKey() string {
	return getenv("RESTCOUNTRIES_KEY")
}

// doGet performs an authenticated GET to REST Countries v5
func doGet(url string) (*http.Response, error) {
	key := restCountriesKey()
	if key == "" {
		return nil, fmt.Errorf("RESTCOUNTRIES_KEY not set in .env")
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+key)
	return http.DefaultClient.Do(req)
}

// v5Response mirrors the top-level wrapper: {"data":{"objects":[...]}}
type v5Response struct {
	Data struct {
		Objects []map[string]interface{} `json:"objects"`
	} `json:"data"`
}

// transformCountries converts v5 objects into []models.Country
func transformCountries(raw []map[string]interface{}) []models.Country {
	var countries []models.Country

	for _, item := range raw {
		country := models.Country{}

		// names.common
		if names, ok := item["names"].(map[string]interface{}); ok {
			country.Name, _ = names["common"].(string)
		}

		// capitals[0].name
		if caps, ok := item["capitals"].([]interface{}); ok && len(caps) > 0 {
			if cap0, ok := caps[0].(map[string]interface{}); ok {
				country.Capital, _ = cap0["name"].(string)
			}
		}

		// population
		if pop, ok := item["population"].(float64); ok {
			country.Population = int64(pop)
		}

		// region / subregion
		if r, ok := item["region"].(string); ok {
			country.Region = r
		}
		if sr, ok := item["subregion"].(string); ok {
			country.Subregion = sr
		}

		// flag.url_svg
		if flag, ok := item["flag"].(map[string]interface{}); ok {
			country.FlagURL, _ = flag["url_svg"].(string)
		}

		// currencies[0] → "BDT (Bangladeshi taka)"
		if currs, ok := item["currencies"].([]interface{}); ok && len(currs) > 0 {
			if c0, ok := currs[0].(map[string]interface{}); ok {
				code, _ := c0["code"].(string)
				name, _ := c0["name"].(string)
				country.Currency = code + " (" + name + ")"
			}
		}

		// languages[] → "Bengali, ..."
		if langs, ok := item["languages"].([]interface{}); ok {
			var langList []string
			for _, l := range langs {
				if lmap, ok := l.(map[string]interface{}); ok {
					if name, ok := lmap["name"].(string); ok && name != "" {
						langList = append(langList, name)
					}
				}
			}
			country.Languages = strings.Join(langList, ", ")
		}

		// coordinates.lat / coordinates.lng
		if coords, ok := item["coordinates"].(map[string]interface{}); ok {
			country.Lat, _ = coords["lat"].(float64)
			country.Lon, _ = coords["lng"].(float64)
		}

		country.Slug = strings.ToLower(strings.ReplaceAll(country.Name, " ", "-"))
		if country.Name != "" {
			countries = append(countries, country)
		}
	}
	return countries
}

// fetchAndTransform does GET → unwrap data.objects → transform
func fetchAndTransform(url string) ([]models.Country, error) {
	resp, err := doGet(url)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API status: %d", resp.StatusCode)
	}

	var wrapper v5Response
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}
	return transformCountries(wrapper.Data.Objects), nil
}

// GetFeaturedCountries returns a curated list for the home page
func GetFeaturedCountries() ([]models.Country, error) {
	names := []string{"united+states", "france", "japan", "australia", "brazil", "bangladesh"}
	var featured []models.Country
	for _, name := range names {
		countries, err := fetchAndTransform(restCountriesBase + "?q=" + name)
		if err != nil || len(countries) == 0 {
			continue
		}
		featured = append(featured, countries[0])
	}
	return featured, nil
}

// GetAllCountries returns all countries
func GetAllCountries() ([]models.Country, error) {
    
    return fetchAndTransform(restCountriesBase)
}

// SearchCountries filters by name/capital and optional region
func SearchCountries(search, region string) ([]models.Country, error) {
	var all []models.Country
	var err error

	if search != "" {
		q := strings.ReplaceAll(strings.TrimSpace(search), " ", "+")
		all, err = fetchAndTransform(restCountriesBase + "?q=" + q)
		if err != nil {
			return nil, err
		}
	} else {
		all, err = GetAllCountries()
		if err != nil {
			return nil, err
		}
	}

	if region == "" {
		return all, nil
	}

	var result []models.Country
	for _, c := range all {
		if strings.EqualFold(c.Region, region) {
			result = append(result, c)
		}
	}
	return result, nil
}

// GetCountryBySlug finds a single country by URL slug e.g. "united-states"
func GetCountryBySlug(slug string) (*models.Country, error) {
	q := strings.ReplaceAll(slug, "-", "+")
	resp, err := doGet(restCountriesBase + "?q=" + q)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("country not found")
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API status: %d", resp.StatusCode)
	}

	var wrapper v5Response
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}
	if len(wrapper.Data.Objects) == 0 {
		return nil, fmt.Errorf("country not found")
	}

	countries := transformCountries(wrapper.Data.Objects)
	if len(countries) == 0 {
		return nil, fmt.Errorf("country not found")
	}
	return &countries[0], nil
}
