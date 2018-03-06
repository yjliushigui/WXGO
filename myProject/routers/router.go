package routers

import (
	"myProject/controllers"
	"github.com/astaxie/beego"
	"myProject/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/admin/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/admin", &admin.IndexController{})
	beego.Router("/admin/index", &admin.IndexController{})

	beego.Router("/test", &admin.IndexController{},"get:Test")

	beego.Router("/admin/blog", &admin.BlogController{},"get:Get")
	beego.Router("/admin/blog_operate", &admin.BlogController{}, "get:Operate;post:Post")

	beego.Router("/admin/table", &admin.TableController{})
	beego.Router("/admin/basic_table", &admin.TableController{})
	beego.Router("/admin/responsive_table", &admin.TableController{}, "get:ResponsiveTable")
	beego.Router("/admin/form_component", &admin.FormsController{})
	beego.Router("/admin/calendar", &admin.ComponentsController{}, "get:Calendar")
	beego.Router("/admin/gallery", &admin.ComponentsController{}, "get:Gallery")
	beego.Router("/admin/todo_list", &admin.ComponentsController{}, "get:TodoList")
	beego.Router("/admin/general", &admin.UIController{}, "get:General")
	beego.Router("/admin/buttons", &admin.UIController{}, "get:Buttons")
	beego.Router("/admin/panels", &admin.UIController{}, "get:Panels")
	beego.Router("/admin/morris", &admin.ChartsController{}, "get:Morris")
	beego.Router("/admin/chartjs", &admin.ChartsController{}, "get:Chartjs")

	/*自定义错误界面*/
	beego.Router("/error", &controllers.ErrorController{}, "*:Error")
	beego.Router("/error/:status/:message/:url/", &controllers.ErrorController{}, "*:Error")
	beego.Router("/error/:status/:message/", &controllers.ErrorController{}, "*:Error")
}
