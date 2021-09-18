package controllers

import (
	_ "fmt"
	beego "github.com/beego/beego/v2/server/web"
	Model "MermaidNFT/models"
	service "MermaidNFT/services"
)

type ConnectController struct {
	beego.Controller
}


func (c *ConnectController) Post() {


	address := c.GetString("address")

	num := Model.AddressExist(address)
	
	if num == 0 {
		tokenId := service.GetTokenId()
	
		Model.InsertUser(address, tokenId);
	}
	
	

	c.Data["json"] = true
    c.ServeJSON()
}




