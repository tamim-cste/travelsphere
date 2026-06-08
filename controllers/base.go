package controllers

import "github.com/beego/beego/v2/server/web"

type BaseController struct {
    web.Controller
}

func (b *BaseController) Prepare() {
    // This section is used for highlighting the active page in Navigation
    b.Data["AppName"] = "travelsphere"
    b.Data["CurrentPath"] = b.Ctx.Request.URL.Path
}