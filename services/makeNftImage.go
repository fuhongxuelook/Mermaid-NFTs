package services

import (
	"log"
	"github.com/issue9/watermark"
)

const ORIGIN = "static/img/origin/"
const RESOURCE = "static/img/resource/"

func GenerageWm(tokenId string, ext string) {
	b := watermark.IsAllowExt(ext)
	if(!b) {
		log.Fatal("image not support")
	}
	w, err := watermark.New("static/img/waterMarkLogo/wm.png", 2, watermark.Center)
	if err != nil{
	    log.Fatal("err is %#v\n", err)
	}
	w.MarkFile(RESOURCE + tokenId + ext)
}