package main

import (
    "github.com/joho/godotenv"
    _ "travelsphere/routers"
    beego "github.com/beego/beego/v2/server/web"
)

func main() {
    
    godotenv.Load()
    beego.Run()
}