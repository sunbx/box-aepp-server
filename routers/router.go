package routers

import (
	"box/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/api/ae/block_top", &controllers.BlockTopController{})


	//抓取数据

	beego.Router("/article/data", &controllers.ArticleDataController{})

	//api调用
	beego.Router("/article/list", &controllers.ArticleListController{})

	//aeasy login
	beego.Router("/api/login", &controllers.ApiLoginController{})

	//aeasy register
	beego.Router("/api/register", &controllers.ApiRegisterController{})

	//ApiNamesAuctionsController
	beego.Router("/api/names/auctions", &controllers.ApiNamesAuctionsController{})

	//ApiNamesPriceController
	beego.Router("/api/names/price", &controllers.ApiNamesPriceController{})

	//ApiNamesPriceController
	beego.Router("/api/names/over", &controllers.ApiNamesOverController{})

	//ApiNamesPriceController
	beego.Router("/api/names/my/register", &controllers.ApiNamesMyRegisterController{})

	//ApiNamesPriceController
	beego.Router("/api/names/my/over", &controllers.ApiNamesMyOverController{})

	//ApiNamesPriceController
	beego.Router("/api/names/add", &controllers.ApiNamesAddController{})

	//ApiNamesUpdateController
	beego.Router("/api/names/update", &controllers.ApiNamesUpdateController{})

	//ApiNamesPriceController
	beego.Router("/api/names/info", &controllers.ApiNamesInfoController{})

	//ApiTransferAddController
	beego.Router("/api/names/transfer", &controllers.ApiTransferAddController{})

	//ApiNamesPriceController
	beego.Router("/api/user/info", &controllers.ApiUserInfoController{})

}
