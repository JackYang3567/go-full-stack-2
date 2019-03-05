package controllers

import (
	"github.com/astaxie/beego"
	//f:"html/template"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get(){
	
	this.Ctx.WriteString("hello welcome")
}
