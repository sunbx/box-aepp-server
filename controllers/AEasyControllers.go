package controllers

import (
	"box/models"
	"box/utils"
	"encoding/json"
	"github.com/aeternity/aepp-sdk-go/config"
	"github.com/aeternity/aepp-sdk-go/naet"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type BlockTopController struct {
	BaseController
}
type HomeController struct {
	BaseController
}
type ApiBaseDataController struct {
	BaseController
}
type NamesBaseController struct {
	BaseController
}
type ApiLoginController struct {
	BaseController
}
type ApiRegisterController struct {
	BaseController
}

type ApiWalletTransferRecordController struct {
	BaseController
}
type WalletTransferController struct {
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

type ApiNamesUpdateController struct {
	BaseController
}

type ApiNamesInfoController struct {
	BaseController
}
type ApiNamesAddController struct {
	BaseController
}
type ApiTransferAddController struct {
	BaseController
}
type ApiUserInfoController struct {
	BaseController
}
type ApiContractCallController struct {
	BaseController
}
type ApiContractCallStaticController struct {
	BaseController
}
type ApiContractDecideController struct {
	BaseController
}

//var HOST =   "https://aeasy.io"
var HOST = "http://localhost:8088"

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

	c.TplName = "index.html"
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

func (c *ApiLoginController) Post() {
	mnemonic := c.GetString("mnemonic")
	resp, err := http.PostForm(HOST+"/api/user/login",
		url.Values{
			"app_id":   {beego.AppConfig.String("AEASY::appId")},
			"mnemonic": {mnemonic},
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

func (c *WalletTransferController) Post() {
	amount := c.GetString("amount")
	address := c.GetString("address")
	signingKey := c.GetString("signingKey")
	data := c.GetString("data")
	resp, err := http.PostForm(HOST+"/api/wallet/transfer",
		url.Values{
			"app_id":     {beego.AppConfig.String("AEASY::appId")},
			"address":    {address},
			"amount":     {amount},
			"signingKey": {signingKey},
			"data":       {data},
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

func (c *ApiRegisterController) Post() {
	resp, err := http.PostForm(HOST+"/api/user/register",
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

func (c *ApiNamesUpdateController) Post() {
	name := c.GetString("name")
	signingKey := c.GetString("signingKey")
	resp, err := http.PostForm(HOST+"/api/names/update",
		url.Values{
			"app_id":     {beego.AppConfig.String("AEASY::appId")},
			"signingKey": {signingKey},
			"name":       {name},
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

func (c *ApiNamesAddController) Post() {
	name := c.GetString("name")
	signingKey := c.GetString("signingKey")
	resp, err := http.PostForm(HOST+"/api/names/add",
		url.Values{
			"app_id":     {beego.AppConfig.String("AEASY::appId")},
			"name":       {name},
			"signingKey": {signingKey},
		})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		c.ErrorJson(-500, e.Error(), JsonData{})
		return
	}
	c.Ctx.WriteString(string(body))
}

func (c *ApiTransferAddController) Post() {
	name := c.GetString("name")
	signingKey := c.GetString("signingKey")
	recipientAddress := c.GetString("recipientAddress")
	resp, err := http.PostForm(HOST+"/api/names/transfer",
		url.Values{
			"app_id":           {beego.AppConfig.String("AEASY::appId")},
			"name":             {name},
			"signingKey":       {signingKey},
			"recipientAddress": {recipientAddress},
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
	print("address->", address)
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
	c.Ctx.WriteString(string(body))
}

func (c *ApiContractDecideController) Get() {
	hash := c.GetString("hash")
	function := c.GetString("function")
	response := utils.Get(models.NodeURL + "/v2/transactions/" + hash + "/info")
	var callInfoResult models.CallInfoResult
	err := json.Unmarshal([]byte(response), &callInfoResult)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	compile := naet.NewCompiler("http://localhost:3080", false)
	source, _ := ioutil.ReadFile("contract/AMBLockContract.aes")
	decodeResult, err := compile.DecodeCallResult(callInfoResult.CallInfo.ReturnType, callInfoResult.CallInfo.ReturnValue, function, string(source), config.Compiler.Backend)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.SuccessJson(decodeResult)
}

func (c *ApiContractCallController) Get() {

	signingKey := c.GetString("signingKey")
	function := c.GetString("function")
	params := c.GetString("params")
	amount, _ := c.GetFloat("amount", 0)

	account, err := models.SigningKeyHexStringAccount(signingKey)
	if amount > 0 {

		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		accountNet, err := models.ApiGetAccount(account.Address)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		tokens, err := strconv.ParseFloat(accountNet.Balance.String(), 64)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}

		if tokens/1000000000000000000 <= amount {
			c.ErrorJson(-500, "token low", JsonData{})
			return
		}

	}

	//println(string(paramsArr[0]))
	//println(string(paramsArr[1]))

	call, functionEncode, err := models.CallContractFunction(account, "ct_2bGpeejyCkhgv452BGchWQYpi7qAudCWKJk33VMWnZ6o7V8eze", function, []string{params}, amount)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	//tx, err := json.Marshal(&call)
	//if err != nil {
	//	c.ErrorJson(-500, err.Error(), JsonData{})
	//	return
	//}
	//c.SuccessJson(call)
	c.SuccessJson(map[string]interface{}{"tx": call, "function": functionEncode})
}



func (c *ApiContractCallStaticController) Get() {

	function := c.GetString("function")
	params := c.GetString("params")
	address := c.GetString("address")



	//println(string(paramsArr[0]))
	//println(string(paramsArr[1]))

	call, functionEncode, err := models.CallStaticContractFunction(address, "ct_2bGpeejyCkhgv452BGchWQYpi7qAudCWKJk33VMWnZ6o7V8eze", function, []string{params})
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	//tx, err := json.Marshal(&call)
	//if err != nil {
	//	c.ErrorJson(-500, err.Error(), JsonData{})
	//	return
	//}
	//c.SuccessJson(call)
	c.SuccessJson(map[string]interface{}{"tx": call, "function": functionEncode})
}

