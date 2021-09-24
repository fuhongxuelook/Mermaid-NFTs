package controllers

import (
	_ "os"
	beego "github.com/beego/beego/v2/server/web"
	service "MermaidNFT/services"
)

type RemoveWMController struct {
	beego.Controller
}

const ORIGIN = "static/img/origin/"
const RESOURCE = "static/img/resource/"

func (c *RemoveWMController) Post() {
	//tokenId := c.GetString("tokenId")

	image := c.GetString("image")

	service.Copy(ORIGIN + image, RESOURCE + image)
	
	c.Data["json"] = image
    c.ServeJSON()

}
