package controllers

import (
	"os"
	beego "github.com/beego/beego/v2/server/web"
)

type RemoveWMController struct {
	beego.Controller
}

const ORIGIN = "static/img/origin/"
const RESOURCE = "static/img/resource/"

func (c *RemoveWMController) Post() {
	tokenId := c.GetString("tokenId")

	// 安全的文件替换
	if _, err := os.Stat(ORIGIN + tokenId); os.IsNotExist(err) {
		os.Rename(ORIGIN + tokenId, RESOURCE + tokenId)
	}
	

}
