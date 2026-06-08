package routers

import (
    "travelsphere/controllers"

    beego "github.com/beego/beego/v2/server/web"
)

func init() {

    beego.Router("/", &controllers.HomeController{})

    beego.Router("/countries", &controllers.CountryController{}, "get:GetCountries")

    beego.Router("/countries/:slug", &controllers.CountryController{}, "get:GetCountryDetail")

    beego.Router("/wishlist", &controllers.WishlistController{})

    beego.Router("/dashboard", &controllers.DashboardController{})
}