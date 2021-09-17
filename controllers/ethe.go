package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	service "MermaidNFT/services"
)

type EtheController struct {
	beego.Controller
}


func (c *EtheController) Get() {


	res := service.GetTokenId()


	c.Data["json"] = &res
    c.ServeJSON()

}

func (c *EtheController) Post() {

	addr := c.GetString("address")

	res := service.MintNFT(addr)


	c.Data["json"] = &res
    c.ServeJSON()

}

