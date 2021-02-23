package routers

import (
	"box/controllers"
	"github.com/astaxie/beego"
)

func init() {

	//获取当前区块高度
	beego.Router("/", &controllers.HomeController{})

	//获取banner
	beego.Router("/api/banner", &controllers.BannerController{})

	//获取token列表
	beego.Router("/api/tokens/list", &controllers.TokenListController{})

	//获取当前区块高度
	beego.Router("/api/ae/block_top", &controllers.BlockTopController{})

	//aeasy login
	beego.Router("/api/base/data", &controllers.ApiBaseDataController{})

	//所有域名基础信息
	beego.Router("/api/names/base", &controllers.NamesBaseController{})

	//转账
	beego.Router("/api/wallet/transfer", &controllers.WalletTransferController{})

	//广播
	beego.Router("/api/tx/broadcast", &controllers.TxBroadcastController{}) //广播
	beego.Router("/api/ae/th_hash", &controllers.ThHashController{})

	//aeasy register
	beego.Router("/api/wallet/transfer/record", &controllers.ApiWalletTransferRecordController{})

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
	beego.Router("/api/names/claim", &controllers.ApiNamesAddController{})

	//域名声明
	beego.Router("/api/names/preclaim", &controllers.PreclaimController{})

	//ApiNamesUpdateController
	beego.Router("/api/names/update", &controllers.ApiNamesUpdateController{})

	//ApiNamesPriceController
	beego.Router("/api/names/info", &controllers.ApiNamesInfoController{})

	//ApiTransferAddController
	beego.Router("/api/names/transfer", &controllers.ApiTransferAddController{})

	//ApiNamesPriceController
	beego.Router("/api/user/info", &controllers.ApiUserInfoController{})
	beego.Router("/api/version", &controllers.ApiVersionController{})
	//ApiNamesPriceController
	beego.Router("/api/contract/call", &controllers.ApiContractCallController{})

	//static call
	beego.Router("/api/contract/call/static", &controllers.ApiContractCallStaticController{})

	//ApiNamesPriceController
	beego.Router("/api/contract/decode", &controllers.ApiContractDecideController{})


	beego.Router("/api/contract/swap/records", &controllers.ApiContractSwapRecordController{})
	beego.Router("/api/contract/swap/records/my", &controllers.ApiContractSwapRecordMyController{})
	beego.Router("/api/contract/swap/records/my/buy", &controllers.ApiContractSwapRecordMyBuyController{})
	beego.Router("/api/contract/swap/records/my/sell", &controllers.ApiContractSwapRecordMySellController{})

	beego.Router("/api/contract/balance", &controllers.ApiContractBalanceController{})
	beego.Router("/api/contract/info", &controllers.ApiContractInfoController{})

	beego.Router("/api/contract/record", &controllers.ApiContractRecordController{})
	beego.Router("/api/contract/ranking", &controllers.ApiContractRankingController{})

	beego.Router("/api/contract/lock", &controllers.ApiContractLockController{})
	beego.Router("/api/contract/unlock", &controllers.ApiContractDecideController{})
	beego.Router("/api/contract/continue", &controllers.ApiContractDecideController{})

	beego.Router("/api/contract/transfer", &controllers.ApiContractTransferController{})
	beego.Router("/test", &controllers.TESTController{})

}
