package routers

import (
	"MermaidNFT/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/receiveImage", &controllers.DealImageController{})
    beego.Router("/", &controllers.MainController{})
    beego.Router("/", &controllers.RemoveWMController{})
}
