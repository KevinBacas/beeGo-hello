package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "kevinbacas.fr"
	c.Data["Email"] = "contact@kevinbacas.fr"
	c.TplNames = "index.tpl"
}
