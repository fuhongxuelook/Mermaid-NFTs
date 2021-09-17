package controllers

import (
	_ "fmt"
	beego "github.com/beego/beego/v2/server/web"
	Model "MermaidNFT/models"
)

type ConnectController struct {
	beego.Controller
}


func (c *ConnectController) Post() {
	
	address := c.GetString("address")
	tokenId := c.GetString("tokenId")
	
	Model.InsertUser(address, tokenId);
	
	

	c.Data["json"] = true
    c.ServeJSON()
}




