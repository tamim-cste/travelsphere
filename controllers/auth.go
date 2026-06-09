package controllers

import "travelsphere/services"

// AuthController handles /login and /logout
type AuthController struct {
	BaseController
}

// Get renders the login page
func (a *AuthController) Get() {
	a.Data["Title"] = "Login"
	a.Layout = "layout/main.tpl"
	a.TplName = "login.tpl"
}

// Post processes login — username only, no password required
func (a *AuthController) Post() {
	username := a.GetString("username")

	user, err := services.GetOrCreateUser(username)
	if err != nil {
		a.Data["Title"] = "Login"
		a.Data["Error"] = err.Error()
		a.Layout = "layout/main.tpl"
		a.TplName = "login.tpl"
		return
	}

	a.SetSession("username", user.Username)
	a.Redirect("/dashboard", 302)
}

// LogoutController handles GET /logout
type LogoutController struct {
	BaseController
}

func (l *LogoutController) Get() {
	l.DestroySession()
	l.Redirect("/", 302)
}
