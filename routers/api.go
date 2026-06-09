package routers

import (
	"travelsphere/controllers/api"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	
	web.Router("/api/countries", &api.CountriesController{})
	web.Router("/api/countries/:slug", &api.CountriesController{}, "get:GetOne")
	web.Router("/api/wishlist", &api.WishlistAPIController{})
	web.Router("/api/wishlist/:id", &api.WishlistAPIController{})
	web.Router("/api/dashboard/summary", &api.DashboardAPIController{})
}
