package controllers

import (
	Model "MermaidNFT/models"
	beego "github.com/beego/beego/v2/server/web"
	service "MermaidNFT/services"
)

type ListController struct {
	beego.Controller
}

const ONE_PAGE_NUM = 10

// 127.0.0.1:8080/static/img/resource/26.jpeg
func (c *ListController) Post() {
	page, _ := c.GetInt("page")
	query := c.GetString("query")

	start := (page - 1) * ONE_PAGE_NUM;

	res := Model.GetList(start, ONE_PAGE_NUM, query)


	c.Data["json"] = &res
    c.ServeJSON()

}

func (c *ListController) Get() {
	tokenId := service.GetTokenId()
	address := c.GetString("address")
	name := c.GetString("name")
	image := c.GetString("image")

	Model.InsertNft(address, tokenId, name, image)


    c.ServeJSON()

}
