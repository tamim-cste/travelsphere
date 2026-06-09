package controllers

import (
	"travelsphere/models"
	"travelsphere/services"
)


type HomeController struct {
	BaseController
}


func (h *HomeController) Get() {
	countries, err := services.GetFeaturedCountries()
	if err != nil {
		h.Data["FeaturedCountries"] = []models.Country{}
	} else {
		h.Data["FeaturedCountries"] = countries
	}

	h.Data["Title"] = "Home"
	h.Layout = "layout/main.tpl"
	h.TplName = "home.tpl"
}
