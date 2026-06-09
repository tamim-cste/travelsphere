package services

import (
	"encoding/json"
	"fmt"
	"travelsphere/models"
)

func GetWeather(city string) (*models.Weather, error) {
	apiKey := getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return nil, nil
	}

	url := fmt.Sprintf(
		"http://api.weatherapi.com/v1/current.json?key=%s&q=%s",
		apiKey, city,
	)

	resp, err := httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("weather API status: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	weather := &models.Weather{}

	// location
	if loc, ok := result["location"].(map[string]interface{}); ok {
		weather.City, _ = loc["name"].(string)
	}

	// current
	if current, ok := result["current"].(map[string]interface{}); ok {
		weather.TempC, _ = current["temp_c"].(float64)
		weather.Humidity = int(current["humidity"].(float64))
		weather.WindKph, _ = current["wind_kph"].(float64)

		if cond, ok := current["condition"].(map[string]interface{}); ok {
			weather.Condition, _ = cond["text"].(string)
			weather.Icon, _ = cond["icon"].(string)

			icon, _ := cond["icon"].(string)

			if len(icon) > 0 && icon[:2] == "//" {
				icon = "https:" + icon
			}
			weather.Icon = icon

		}
	}

	return weather, nil
}
