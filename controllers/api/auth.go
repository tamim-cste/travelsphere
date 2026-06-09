package controllers

// AuthController handles login and logout
type AuthController struct {
	BaseController
}

// Get renders the login page
func (a *AuthController) Get() {
	a.Data["Title"] = "Login"
	a.Layout = "layout/main.tpl"
	a.TplName = "login.tpl"
}

// Post processes login form submission
func (a *AuthController) Post() {
	username := a.GetString("username")
	password := a.GetString("password")

	// Simple hardcoded check — no database needed for this assessment
	if username == "beta" && password == "1234" {
		a.SetSession("username", username)
		a.Redirect("/dashboard", 302)
		return
	}

	a.Data["Title"] = "Login"
	a.Data["Error"] = "Invalid username or password"
	a.Layout = "layout/main.tpl"
	a.TplName = "login.tpl"
}

// Logout clears the session
type LogoutController struct {
	BaseController
}

func (l *LogoutController) Get() {
	l.DestroySession()
	l.Redirect("/", 302)
}
