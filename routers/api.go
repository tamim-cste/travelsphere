package routers

import (
    "travelsphere/controllers/api"
    "github.com/beego/beego/v2/server/web"
)

func init() {
    web.Router("/api/countries", &api.CountriesController{})
}