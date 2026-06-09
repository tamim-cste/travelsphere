package controllers

import "github.com/beego/beego/v2/server/web"


type BaseController struct {
	web.Controller
}

func (b *BaseController) Prepare() {
	b.Data["AppName"] = "TravelSphere"
	b.Data["CurrentPath"] = b.Ctx.Request.URL.Path
	b.Data["LoggedIn"] = false
	b.Data["Username"] = ""

	
	username := b.readSession("username")
	if username != "" {
		b.Data["LoggedIn"] = true
		b.Data["Username"] = username
	}
}


func (b *BaseController) readSession(key string) string {
	defer func() { recover() }()
	val := b.GetSession(key)
	if val == nil {
		return ""
	}
	if s, ok := val.(string); ok {
		return s
	}
	return ""
}
