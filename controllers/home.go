package controllers

import (
	"travelsphere/services"
	"fmt"
	"travelsphere/models"
)
type HomeController struct {
    BaseController  
}

func (h *HomeController) Get() {
    countries, err := services.GetFeaturedCountries()
    if err != nil {
        fmt.Println("ERROR fetching countries:", err)
        h.Data["FeaturedCountries"] = []models.Country{}
    } else {
        fmt.Println("Fetched countries count:", len(countries))
        h.Data["FeaturedCountries"] = countries
    }

    h.Data["Title"] = "Home"
    h.Layout = "layout/main.tpl"
    h.TplName = "home.tpl"
}