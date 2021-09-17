package controllers

import (
	Model "MermaidNFT/models"
	beego "github.com/beego/beego/v2/server/web"
)

type ListController struct {
	beego.Controller
}

const ONE_PAGE_NUM = 10

func (c *ListController) Post() {
	page, _ := c.GetInt("page")

	start := (page - 1) * ONE_PAGE_NUM;

	res := Model.GetList(start, ONE_PAGE_NUM)


	c.Data["json"] = &res
    c.ServeJSON()

}
