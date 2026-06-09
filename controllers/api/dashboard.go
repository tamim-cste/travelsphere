package api

import (
	"travelsphere/services"
	"travelsphere/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// DashboardAPIController handles /api/dashboard/summary
type DashboardAPIController struct {
	beego.Controller
}

func (d *DashboardAPIController) Get() {
	summary := services.GetDashboardSummary()
	utils.SendSuccess(&d.Controller, summary)
}
