package utils


import "github.com/beego/beego/v2/server/web"

type JSONResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}


func SendSuccess(c *web.Controller, data interface{}) {
	c.Data["json"] = JSONResponse{Success: true, Data: data}
	c.ServeJSON()
}

func SendError(c *web.Controller, code int,message string){
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Data["json"] = JSONResponse{Success: false, Message: message}
	c.ServeJSON()
} 