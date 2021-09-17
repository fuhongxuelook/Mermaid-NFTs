package controllers

import (
	"math/big"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	gasBiger := big.NewInt(10)
	gas      := big.NewInt(2)

	gas.Mul(gas, gasBiger)

	c.Data["json"] = gas
    c.ServeJSON()
}
