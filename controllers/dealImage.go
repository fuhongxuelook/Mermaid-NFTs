package controllers

import (
	_ "fmt"
	"log"
	"path"
	beego "github.com/beego/beego/v2/server/web"
	service "MermaidNFT/services"
	Model "MermaidNFT/models"

)

type DealImageController struct {
	beego.Controller
}


func (c *DealImageController) Post() {
	name := c.GetString("name")
	address := c.GetString("address")
	tokenId := service.GetTokenId()
	//beego.MaxMemory = 1<<22
	f, h, err := c.GetFile("img")
    if err != nil {
        log.Fatal("getfile err ", err)
    }
    defer f.Close()

    ext := path.Ext(h.Filename);

    c.SaveToFile("img", ORIGIN + tokenId + ext) 

    // 安全的文件移动文件
	service.Copy(ORIGIN + tokenId + ext, RESOURCE + tokenId + ext)
	
	service.GenerageWm(tokenId, ext)

	Model.InsertNft(address, tokenId, name, tokenId + ext)

	c.Ctx.Output.Body([]byte(name));
}

func (c *DealImageController) Get() {
    c.TplName = "index.tpl"
}