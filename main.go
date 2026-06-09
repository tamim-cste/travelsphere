package main

import (
	"github.com/joho/godotenv"
	_ "travelsphere/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// Load .env file for API keys
	godotenv.Load()

	// Ensure session is enabled at runtime
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.WebConfig.Session.SessionName = "travelsphere_session"

	beego.Run()
}
