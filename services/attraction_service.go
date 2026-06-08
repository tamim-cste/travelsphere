package services

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "travelsphere/models"
)

func GetAttractions(lat, lon float64, limit int) ([]models.Attraction, error) {
    apiKey := os.Getenv("OPENTRIPMAP_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("OpenTripMap key not set")
    }

    url := fmt.Sprintf(
        "https://api.opentripmap.com/0.1/en/places/radius?radius=10000&lon=%f&lat=%f&limit=%d&apikey=%s",
        lon, lat, limit, apiKey,
    )

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Features []struct {
            Properties struct {
                Name  string `json:"name"`
                Kinds string `json:"kinds"`
                XID   string `json:"xid"`
            } `json:"properties"`
        } `json:"features"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    var attractions []models.Attraction
    for _, f := range result.Features {
        if f.Properties.Name == "" {
            continue
        }
        attractions = append(attractions, models.Attraction{
            Name:  f.Properties.Name,
            Kinds: f.Properties.Kinds,
            XID:   f.Properties.XID,
        })
    }
    return attractions, nil
}