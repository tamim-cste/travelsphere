package routers

import (
	"log"
	"time"
	"travelsphere/controllers"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	// Logging filter
	web.InsertFilter("/*", web.BeforeRouter, func(ctx *context.Context) {
		start := time.Now()
		defer func() {
			log.Printf("[%s] %s — %v",
				ctx.Request.Method,
				ctx.Request.URL.Path,
				time.Since(start),
			)
		}()
	})

	// SSR page routes
	web.Router("/", &controllers.HomeController{})
	web.Router("/countries", &controllers.CountryController{}, "get:Get")
	web.Router("/countries/:slug", &controllers.CountryController{}, "get:GetOne")
	// login must handle both GET (show form) and POST (submit form)
	web.Router("/login", &controllers.AuthController{}, "get:Get;post:Post")
	web.Router("/logout", &controllers.LogoutController{}, "get:Get")
	web.Router("/wishlist", &controllers.WishlistController{}, "get:Get")
	web.Router("/dashboard", &controllers.DashboardController{}, "get:Get")
}
