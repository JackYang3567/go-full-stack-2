package main

import (
	_ "quickstart/routers"
	"github.com/astaxie/beego"

	//"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql" // import your used driver
)

func main() {
	beego.Run()
}

