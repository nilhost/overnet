package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/nilhost/overnet/backend/site/routers"
)

func main() {
	beego.Run()
}

