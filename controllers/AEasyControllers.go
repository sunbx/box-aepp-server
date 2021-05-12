package controllers

import (
	"box/models"
	"box/utils"
	"github.com/aeternity/aepp-sdk-go/naet"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
)

type BlockTopController struct {
	BaseController
}
type HomeController struct {
	BaseController
}
type ServerController struct {
	BaseController
}
type ApiBaseDataController struct {
	BaseController
}
type NamesBaseController struct {
	BaseController
}

type ApiWalletTransferRecordController struct {
	BaseController
}

type ApiNamesAuctionsController struct {
	BaseController
}

type ApiNamesPriceController struct {
	BaseController
}

type ApiNamesOverController struct {
	BaseController
}
type ApiNamesMyRegisterController struct {
	BaseController
}
type ApiNamesMyOverController struct {
	BaseController
}

type ApiNamesInfoController struct {
	BaseController
}
type ApiUserInfoController struct {
	BaseController
}
type ApiVersionController struct {
	BaseController
}

type TESTController struct {
	BaseController
}




var HOST = "https://aeasy.io"

func (c *BlockTopController) Post() {
	resp, err := http.PostForm(HOST+"/api/ae/block_top",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *TESTController) Get() {
	//node := naet.NewNode(models.NodeUrl, false)
	compile := naet.NewCompiler(models.CompilerUrl, false)
	//var source []byte
	//source, _ = ioutil.ReadFile("contract/ABCLockContractV3.aes")

	//decodeResult, err := compile.DecodeCallResult("cb", tryRun.Results[0].CallObj.ReturnValue, function, string(source), config.Compiler.Backend)
	decodedData, err := compile.DecodeData("cb_KxH7D7nnG2+KqSrUWJVW1Jf/wIyjh+w=", "int")
	if err!=nil{
		c.Ctx.WriteString(string(err.Error()))
		return
	}
	c.SuccessJson(decodedData)

}

func (c *NamesBaseController) Post() {
	resp, err := http.PostForm(HOST+"/api/names/base",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *HomeController) Get() {
	if utils.IsMobile(c.Ctx.Input.Header("user-agent")) {
		c.TplName = "index_mobile.html"
	} else {

		c.TplName = "index.html"
	}
}

func (c *ServerController) Get() {
	c.TplName = "ae.html"

}

func (c *ApiBaseDataController) Post() {
	resp, err := http.PostForm(HOST+"/api/base/data",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiWalletTransferRecordController) Post() {
	address := c.GetString("address")
	page := c.GetString("page")
	resp, err := http.PostForm(HOST+"/api/wallet/transfer/record",
		url.Values{
			"app_id":  {beego.AppConfig.String("AEASY::appId")},
			"address": {address},
			"page":    {page},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesAuctionsController) Post() {
	page := c.GetString("page")
	resp, err := http.PostForm(HOST+"/api/names/auctions",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"page":   {page},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesPriceController) Post() {
	page := c.GetString("page")
	resp, err := http.PostForm(HOST+"/api/names/price",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"page":   {page},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesOverController) Post() {
	page := c.GetString("page")
	resp, err := http.PostForm(HOST+"/api/names/over",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"page":   {page},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesMyRegisterController) Post() {
	page := c.GetString("page")
	address := c.GetString("address")

	print("address:", address)
	print("page:", page)
	resp, err := http.PostForm(HOST+"/api/names/my/register",
		url.Values{
			"app_id":  {beego.AppConfig.String("AEASY::appId")},
			"page":    {page},
			"address": {address},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesMyOverController) Post() {
	page := c.GetString("page")
	address := c.GetString("address")
	resp, err := http.PostForm(HOST+"/api/names/my/over",
		url.Values{
			"app_id":  {beego.AppConfig.String("AEASY::appId")},
			"page":    {page},
			"address": {address},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiNamesInfoController) Post() {
	name := c.GetString("name")
	resp, err := http.PostForm(HOST+"/api/names/info",
		url.Values{
			"app_id": {beego.AppConfig.String("AEASY::appId")},
			"name":   {name},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiUserInfoController) Post() {
	address := c.GetString("address")
	resp, err := http.PostForm(HOST+"/api/user/info",
		url.Values{
			"app_id":  {beego.AppConfig.String("AEASY::appId")},
			"address": {address},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	_, err = models.ApiGetAccount(address)
	if err != nil {
		print(err.Error())
		if err.Error() == "Error: Account not found" {
			//account, _ := models.SigningKeyHexStringAccount(beego.AppConfig.String("AEASY::accountFoundation"))
			//tx, _ := models.ApiSpend(account, address, 0.00001, "Sponsored by China Foundation（中国基金会赞助）")
			//print(tx.Hash)

		}
	}

	c.Ctx.WriteString(string(body))
}
func (c *ApiVersionController) Post() {
	source, _ := ioutil.ReadFile("conf/version")
	c.Ctx.WriteString(string(source))
}






