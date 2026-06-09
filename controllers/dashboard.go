package controllers

import "travelsphere/services"


type DashboardController struct {
	BaseController
}

func (d *DashboardController) Prepare() {
	d.BaseController.Prepare()
	loggedIn, ok := d.Data["LoggedIn"].(bool)
	if !ok || !loggedIn {
		d.Redirect("/login", 302)
	}
}

func (d *DashboardController) Get() {
	summary := services.GetDashboardSummary()
	d.Data["Title"] = "Travel Dashboard"
	d.Data["Summary"] = summary
	d.Layout = "layout/main.tpl"
	d.TplName = "dashboard.tpl"
}
