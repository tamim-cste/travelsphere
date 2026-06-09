package api

import (
	"travelsphere/services"
	"travelsphere/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// CountriesController handles /api/countries JSON endpoints
type CountriesController struct {
	beego.Controller
}

// Get returns country list with optional search and region filters
func (c *CountriesController) Get() {
	search := c.GetString("search")
	region := c.GetString("region")

	countries, err := services.SearchCountries(search, region)
	if err != nil {
		utils.SendError(&c.Controller, 500, "Failed to fetch countries")
		return
	}
	utils.SendSuccess(&c.Controller, countries)
}

// GetOne returns a single country by slug
func (c *CountriesController) GetOne() {
	slug := c.Ctx.Input.Param(":slug")
	country, err := services.GetCountryBySlug(slug)
	if err != nil {
		utils.SendError(&c.Controller, 404, "Country not found")
		return
	}
	utils.SendSuccess(&c.Controller, country)
}
