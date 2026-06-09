package controllers

import (
    "fmt"
    "travelsphere/models"
    "travelsphere/services"
)

type CountryController struct {
    BaseController
}

func (c *CountryController) Get() {
    countries, err := services.GetAllCountries()
    if err != nil {
        c.Data["Countries"] = []interface{}{}
        c.Data["Error"] = "Could not load countries"
    } else {
        c.Data["Countries"] = countries
    }

    c.Data["Title"] = "Country Explorer"
    c.Layout = "layout/main.tpl"
    c.TplName = "countries.tpl"
}

func (c *CountryController) GetOne() {
    slug := c.Ctx.Input.Param(":slug")

    country, err := services.GetCountryBySlug(slug)
    if err != nil {
        c.Data["Title"] = "Not Found"
        c.Data["Error"] = "Country not found: " + slug
        c.Layout = "layout/main.tpl"
        c.TplName = "404.tpl"
        c.Ctx.ResponseWriter.WriteHeader(404)
        return
    }

    attractions, err := services.GetAttractions(country.Lat, country.Lon, 10)
    
    weather, _ := services.GetWeather(country.Capital)
    
    //These lines are used for debugging
    fmt.Println("Weather result:", weather)  
    fmt.Println("Capital:", country.Capital) 

    c.Data["Weather"] = weather  
    if err != nil {
        fmt.Println("Attractions error:", err)
        attractions = []models.Attraction{}
    }

    c.Data["Country"] = country
    c.Data["Attractions"] = attractions
    c.Data["Title"] = country.Name
    c.Layout = "layout/main.tpl"
    c.TplName = "destination.tpl"
}