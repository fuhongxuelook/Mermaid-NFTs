package services

import (
	"github.com/issue9/watermark"
)

const ORIGIN = "static/img/origin/"
const RESOURCE = "static/img/resource/"

func GenerageWm(tokenId string, ext string) {
	b := watermark.IsAllowExt(".jpg")
	fmt.Printf("IsAllowExt %v\n", b)
	w, err := watermark.New("static/waterMarkLogo/wm.png", 2, watermark.Center)
	if err != nil{
	    log.Fatal("err is %#v\n", err)
	}
	fmt.Printf("\nw is %v, err is %v\n", w, err)
	file :=  RESOURCE + tokenId;

	Copy(file, file + ext)
	
	err = w.MarkFile("cc.jpg")

	Copy(file + ext, file)
	fmt.Printf("\n MarkFile err is %v\n",  err)
}