package routers

import (
	"box/controllers"
	"github.com/astaxie/beego"
)

func init() {

	//获取当前区块高度
	beego.Router("/", &controllers.HomeController{})

	//获取当前区块高度
	beego.Router("/server", &controllers.ServerController{})

	//获取banner
	beego.Router("/api/banner", &controllers.BannerController{})

	//获取token列表
	beego.Router("/api/tokens/list", &controllers.TokenListController{})

	//获取当前区块高度
	beego.Router("/api/ae/block_top", &controllers.BlockTopController{})

	//基础数据
	beego.Router("/api/base/data", &controllers.ApiBaseDataController{})

	//所有域名基础信息
	beego.Router("/api/names/base", &controllers.NamesBaseController{})

	//获取转账记录
	beego.Router("/api/wallet/transfer/record", &controllers.ApiWalletTransferRecordController{})

	//竞拍中域名-最新
	beego.Router("/api/names/auctions", &controllers.ApiNamesAuctionsController{})

	//竞拍中域名-价格
	beego.Router("/api/names/price", &controllers.ApiNamesPriceController{})

	//即将结束的域名
	beego.Router("/api/names/over", &controllers.ApiNamesOverController{})

	//竞拍中域名-最新-我的
	beego.Router("/api/names/my/register", &controllers.ApiNamesMyRegisterController{})

	//竞拍中域名-即将结束-我的
	beego.Router("/api/names/my/over", &controllers.ApiNamesMyOverController{})

	//域名详情
	beego.Router("/api/names/info", &controllers.ApiNamesInfoController{})

	//获取当前用户信息
	beego.Router("/api/user/info", &controllers.ApiUserInfoController{})

	//获取最新版本号
	beego.Router("/api/version", &controllers.ApiVersionController{})

	//获取具体aex9 abc数量
	beego.Router("/api/aex9/balance", &controllers.ApiContractBalanceController{})

	beego.Router("/api/defi/info", &controllers.ApiContractInfoController{})

	beego.Router("/api/aex9/ranking", &controllers.ApiContractRankingController{})

	beego.Router("/api/defi/status", &controllers.DefiStatusController{})

	beego.Router("/test", &controllers.TESTController{})

}
