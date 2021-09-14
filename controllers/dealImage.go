package controllers

import (
	_ "fmt"
	"log"
	"path/filepath"
	beego "github.com/beego/beego/v2/server/web"
	service "MermaidNFT/services"
)

type DealImageController struct {
	beego.Controller
}


func (c *DealImageController) Post() {
	name := c.GetString("name")
	tokenId := c.GetString("tokenId")
	//beego.MaxMemory = 1<<22
	f, _, err := c.GetFile("img")
    if err != nil {
        log.Fatal("getfile err ", err)
    }
    defer f.Close()

    ext := filepath.Ext(f)
    c.SaveToFile("img", ORIGIN + tokenId + ".jpeg") 

    // 安全的文件移动文件
	service.Copy(ORIGIN + tokenId, RESOURCE + tokenId)
	
	service.GenerageWm(tokenId, "")

	c.Ctx.Output.Body([]byte(name));
}

func (c *DealImageController) Get() {
    c.TplName = "index.tpl"
}