package services

import (
    "encoding/json"
	"strings"
    "fmt"
    "net/http"
    "travelsphere/models"
)

const restCountriesBase = "https://restcountries.com/v3.1"

func transformCountries(raw []map[string]interface{}) []models.Country {
    var countries []models.Country

    for _, item := range raw {
        country := models.Country{}

        // Name
        if nameObj, ok := item["name"].(map[string]interface{}); ok {
            country.Name = nameObj["common"].(string)
        }

        // Capitals
        if capitals, ok := item["capital"].([]interface{}); ok && len(capitals) > 0 {
            country.Capital = capitals[0].(string)
        }

        // Population
        if pop, ok := item["population"].(float64); ok {
            country.Population = int64(pop)
        }

        // Region
        if region, ok := item["region"].(string); ok {
            country.Region = region
        }

        // Flag
        if flags, ok := item["flags"].(map[string]interface{}); ok {
            country.FlagURL = flags["svg"].(string)
        }

        
        if currencies, ok := item["currencies"].(map[string]interface{}); ok {
            for code, val := range currencies {
                if currObj, ok := val.(map[string]interface{}); ok {
                    country.Currency = code + " (" + currObj["name"].(string) + ")"
                }
                break 
            }
        }

      
        if langs, ok := item["languages"].(map[string]interface{}); ok {
            var langList []string
            for _, v := range langs {
                langList = append(langList, v.(string))
            }
            country.Languages = strings.Join(langList, ", ")
        }

       
        country.Slug = strings.ToLower(
            strings.ReplaceAll(country.Name, " ", "-"),
        )

        countries = append(countries, country)
    }

    return countries
}



func GetFeaturedCountries() ([]models.Country, error) {
    
    url := restCountriesBase + "/region/asia?fields=name,capital,population,flags,currencies,languages,region"
    
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("API call failed: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
    }

    var raw []map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
        return nil, fmt.Errorf("JSON decode failed: %w", err)
    }

    return transformCountries(raw), nil
}