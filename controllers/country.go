package controllers

import "travelsphere/services"

type CountryController struct {
    BaseController
}

func (c *CountryController) Get() {
    // Initial SSR load
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