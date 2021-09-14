package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ListController struct {
	beego.Controller
}

const ONE_PAGE_NUM = 10

func (c *ListController) Get() {
	page, _ := c.GetInt("page")

	start := (page - 1) * ONE_PAGE_NUM;

	//list := getDataList(start, ONE_PAGE_NUM);
	if start < 0 {
		return;
	}


}
