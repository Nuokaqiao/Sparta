package routers

import (
	"github.com/astaxie/beego"
	"sparta/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
