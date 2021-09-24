package controllers

import (
	Model "MermaidNFT/models"
	beego "github.com/beego/beego/v2/server/web"
	_  "MermaidNFT/services"
)

type ListController struct {
	beego.Controller
}


type Res struct{
	List []Model.Nft `json:"list"`
	Total int64 `json:"total"`
	CurrentPage int `json:"current_page"`
}

const ONE_PAGE_NUM = 20

// 127.0.0.1:8080/static/img/resource/26.jpeg
func (c *ListController) Post() {
	page, _ := c.GetInt("page")
	query := c.GetString("query")

	start := (page - 1) * ONE_PAGE_NUM;

	res := Model.GetList(start, ONE_PAGE_NUM, query)

	num := Model.GetListNum(start, ONE_PAGE_NUM, query)

	r := Res{List: res, Total: num, CurrentPage: page}

	c.Data["json"] = &r
    c.ServeJSON()

}

func (c *ListController) Get() {
	page, _ := c.GetInt("page")
	query := c.GetString("query")

	start := (page - 1) * ONE_PAGE_NUM;

	num := Model.GetListNum(start, ONE_PAGE_NUM, query)


	c.Data["json"] = &num
    c.ServeJSON()

}
