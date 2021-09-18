package controllers

import (
	_ "fmt"
	beego "github.com/beego/beego/v2/server/web"
	service "MermaidNFT/services"
	_ "github.com/beego/beego/v2/core/logs"
)

type EtheController struct {
	beego.Controller
}


// func (c *EtheController) Get() {
// 	logs.SetLogger("console")

// 	logs.SetLevel(logs.LevelDebug)
// 	res := service.GetTokenId()
// 	logs.Debug("\n ----res--- is %v ", res)

	

// 	c.Data["json"] = res
//     c.ServeJSON()

// }

// func (c *EtheController) Post() {

// 	addr := c.GetString("address")

// 	res := service.MintNFT(addr)


// 	c.Data["json"] = &res
//     c.ServeJSON()

// }

func (c *EtheController) Get() {
	hash := c.GetString("hash")
	res := service.LoopSearchTx(hash)


	c.Data["json"] = &res
    c.ServeJSON()
}








