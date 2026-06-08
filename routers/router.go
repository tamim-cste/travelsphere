package routers

import (
    "log"
    "time"
    "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/server/web/context"
)

func init() {
    // Logging filter 
    web.InsertFilter("/*", web.BeforeRouter, func(ctx *context.Context) {
        start := time.Now()
        //It will show logs after end of request
        defer func() {
            log.Printf("[%s] %s — %v",
                ctx.Request.Method,
                ctx.Request.URL.Path,
                time.Since(start),
            )
        }()
    })

    // SSR Routes 
    web.Router("/", &controllers.HomeController{})
}