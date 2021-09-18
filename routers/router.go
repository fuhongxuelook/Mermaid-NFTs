package routers

import (
	"MermaidNFT/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/receiveImage", &controllers.DealImageController{})
	beego.Router("/remove", &controllers.RemoveWMController{})
	beego.Router("/list", &controllers.ListController{})
	beego.Router("/connect", &controllers.ConnectController{})
    beego.Router("/ethe", &controllers.EtheController{})
    //beego.Router("/ethe", &controllers.EtheController{})
    beego.Router("/", &controllers.MainController{})

    

}
//remixd -s ./ -u "https://remix.ethereum.org/#optimize=false&runs=200&evmVersion=null&version=soljson-v0.8.7+commit.e28d00a7.js"