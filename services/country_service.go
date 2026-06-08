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


        // Coordinates
        if latlng, ok := item["latlng"].([]interface{}); ok && len(latlng) >= 2 {
            country.Lat, _ = latlng[0].(float64)
            country.Lon, _ = latlng[1].(float64)
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




// Fetch all countries

func GetAllCountries() ([]models.Country, error) {
    url := restCountriesBase + "/all?fields=name,capital,population,flags,currencies,languages,region"
    return fetchAndTransform(url)
}

// Search + Region filter
func SearchCountries(search, region string) ([]models.Country, error) {
    all, err := GetAllCountries()
    if err != nil {
        return nil, err
    }

    var result []models.Country
    search = strings.ToLower(search)

    for _, c := range all {
        // Region filter
        if region != "" && strings.ToLower(c.Region) != strings.ToLower(region) {
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

// Reusable helper — fetch + transform
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








// func GetFeaturedCountries() ([]models.Country, error) {
    
//     url := restCountriesBase + "/region/asia?fields=name,capital,population,flags,currencies,languages,region"
    
//     resp, err := http.Get(url)
//     if err != nil {
//         return nil, fmt.Errorf("API call failed: %w", err)
//     }
//     defer resp.Body.Close()

//     if resp.StatusCode != 200 {
//         return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
//     }

//     var raw []map[string]interface{}
//     if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
//         return nil, fmt.Errorf("JSON decode failed: %w", err)
//     }

//     return transformCountries(raw), nil
// }




func GetFeaturedCountries() ([]models.Country, error) {
    url := restCountriesBase + "/region/asia?..."
    return fetchAndTransform(url) 
}





func GetCountryBySlug(slug string) (*models.Country, error) {
    // extract country name from slug
    // "united-states" → "united states"
    name := strings.ReplaceAll(slug, "-", " ")
    
    url := restCountriesBase + "/name/" + name + 
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