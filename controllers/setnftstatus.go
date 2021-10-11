package controllers

import (
	Model "MermaidNFT/models"
	beego "github.com/beego/beego/v2/server/web"
)

type SetStatusController struct {
	beego.Controller
}

func (c *SetStatusController) Post() {
	status := c.GetString("status")
	tokenId := c.GetString("tokenId")

	Model.ChangeNftTokenIdStatus(tokenId, status)

	c.Data["json"] = true
    c.ServeJSON()
}


func (c *SetStatusController) Get() {
	tokenId := Model.GetNFTId()

	c.Data["json"] = tokenId
    c.ServeJSON()
}