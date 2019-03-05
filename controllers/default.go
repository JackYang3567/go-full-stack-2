package controllers

import (
	"github.com/astaxie/beego"
	//f:"html/template"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["Name"] = "Jack"

	//this.XSRFExpire = 7200
	this.Data["xsrfdata"]= beego.AppConfig.String("xsrfkey")
	this.Layout = "frame_layout/master.html"
	this.TplName = "index.1.tpl"
	this.LayoutSections = make(map[string]string)
    this.LayoutSections["HtmlHead"] = "frame_layout/header.html"    
	this.LayoutSections["Sidebar"] = "frame_layout/sidebar.html"
	this.LayoutSections["HtmlFoot"] = "frame_layout/footer.html"
	this.LayoutSections["Scripts"] = "frame_layout/scripts.tpl"
	//this.TplName = "index.html"
	//this.Ctx.WriteString("hello")
}

func (this *MainController) Welcome() {
	//this.Ctx.WriteString("==hello welcome ===")
	this.Layout = "layout/master.html"
	this.TplName = "admin/welcome.tpl"
	this.LayoutSections = make(map[string]string)
    this.LayoutSections["HtmlHead"] = "layout/header.html"    
	this.LayoutSections["Sidebar"] = "layout/sidebar.html"
	this.LayoutSections["HtmlFoot"] = "layout/footer.html"
	this.LayoutSections["Scripts"] = "layout/scripts.tpl"
}

