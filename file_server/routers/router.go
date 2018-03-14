package routers

import (
	"file_server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/upload", &controllers.MainController{})
	beego.Router("/picture", &controllers.ImgController{})
}
