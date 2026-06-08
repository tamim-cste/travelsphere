package api

import (
    "travelsphere/services"
    "travelsphere/utils"
    beego "github.com/beego/beego/v2/server/web"
)

type CountriesController struct {
    beego.Controller
}

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