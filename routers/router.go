package routers

import (
	"box/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//抓取数据
	beego.Router("/article/data", &controllers.ArticleDataController{})

	//api调用
	beego.Router("/article/list", &controllers.ArticleListController{})

	//aeasy login
	beego.Router("/api/login", &controllers.ApiLoginController{})

	//aeasy register
	beego.Router("/api/register", &controllers.ApiRegisterController{})

	//ApiNamesAuctionsController
	beego.Router("/api/name/auctions", &controllers.ApiNamesAuctionsController{})

	//ApiNamesPriceController
	beego.Router("/api/name/price", &controllers.ApiNamesPriceController{})

	//ApiNamesPriceController
	beego.Router("/api/name/over", &controllers.ApiNamesOverController{})

	//ApiNamesPriceController
	beego.Router("/api/name/my/register", &controllers.ApiNamesMyRegisterController{})

	//ApiNamesPriceController
	beego.Router("/api/name/my/over", &controllers.ApiNamesMyOverController{})

	//ApiNamesPriceController
	beego.Router("/api/name/add", &controllers.ApiNamesAddController{})

	//ApiNamesUpdateController
	beego.Router("/api/name/update", &controllers.ApiNamesUpdateController{})

	//ApiNamesPriceController
	beego.Router("/api/name/info", &controllers.ApiNamesInfoController{})

	//ApiTransferAddController
	beego.Router("/api/name/transfer", &controllers.ApiTransferAddController{})

	//ApiNamesPriceController
	beego.Router("/api/user/info", &controllers.ApiUserInfoController{})

}
