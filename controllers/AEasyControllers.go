package controllers

import (
	"box/models"
	"box/utils"
	"encoding/json"
	"fmt"
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
type ApiContractBalanceController struct {
	BaseController
}
type ApiContractInfoController struct {
	BaseController
}
type ApiContractRecordController struct {
	BaseController
}
type ApiContractLockController struct {
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

func (c *ApiContractDecideController) Post() {
	hash := c.GetString("hash")
	function := c.GetString("function")
	response := utils.Get(models.NodeURL + "/v2/transactions/" + hash + "/info")
	var callInfoResult models.CallInfoResult
	err := json.Unmarshal([]byte(response), &callInfoResult)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	if callInfoResult.Reason == "Tx not mined" {
		c.ErrorJson(-1, "Tx not mined", JsonData{})
		return
	}

	println(response)
	println(hash)
	compile := naet.NewCompiler("http://localhost:3080", false)
	source, _ := ioutil.ReadFile("contract/AMBLockContract.aes")
	decodeResult, err := compile.DecodeCallResult(callInfoResult.CallInfo.ReturnType, callInfoResult.CallInfo.ReturnValue, function, string(source), config.Compiler.Backend)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	if function == "lock" {
		switch decodeResult.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
		case json.Number:
			number, _ := decodeResult.(json.Number).Float64()
			c.SuccessJson(map[string]interface{}{"token_number": utils.FormatTokens(number, 7)})
			return
		case map[string]interface{}:
			data := decodeResult.(map[string]interface{})
			about := data["abort"].([]interface{})
			c.ErrorJson(-500, about[0].(string), JsonData{})
			return
		}
		c.ErrorJson(-500, decodeResult, JsonData{})
		return
	}

	if function == "unlock" {
		switch decodeResult.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
		case json.Number:
			number, _ := decodeResult.(json.Number).Float64()
			c.SuccessJson(map[string]interface{}{"token_number": utils.FormatTokens(number, 7)})
			return
		case map[string]interface{}:
			data := decodeResult.(map[string]interface{})
			about := data["abort"].([]interface{})
			c.ErrorJson(-500, about[0].(string), JsonData{})
			return
		}
		c.ErrorJson(-500, decodeResult, JsonData{})
		return
	}

	if function == "continue_lock" {
		switch decodeResult.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
		case json.Number:
			number, _ := decodeResult.(json.Number).Float64()
			c.SuccessJson(map[string]interface{}{"token_number": utils.FormatTokens(number, 7)})
			return
		case map[string]interface{}:
			data := decodeResult.(map[string]interface{})
			about := data["abort"].([]interface{})
			c.ErrorJson(-500, about[0].(string), JsonData{})
			return
		}
		c.ErrorJson(-500, decodeResult, JsonData{})
		return
	}

	c.SuccessJson(decodeResult)
}

func (c *ApiContractCallController) Post() {

	signingKey := c.GetString("signingKey")
	function := c.GetString("function")
	params := c.GetString("params")
	amount, _ := c.GetFloat("amount", 0)

	if signingKey == "" {
		c.ErrorJson(-500, "signingKey error", JsonData{})
		return
	}

	if function != "lock" && function != "unlock" && function != "continue_lock" {
		c.ErrorJson(-500, "function error", JsonData{})
		return
	}

	if "lock" == function && amount == 0 {
		c.ErrorJson(-500, "amount error", JsonData{})
		return
	}
	println(signingKey)
	println(params)
	println(function)

	account, err := models.SigningKeyHexStringAccount(signingKey)
	if err != nil {
		c.ErrorJson(-500, err.Error()+"123123", JsonData{})
		return
	}
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

		if tokens/100000000000000000 <= amount {
			c.ErrorJson(-500, "token low", JsonData{})
			return
		}

	}
	call, functionEncode, err := models.CallContractFunction(account, "ct_hM2PJEB66Sqx2mkyCixbh3z9hLMaK8N1Sa8v5kaWRqXwPYgkQ", function, []string{params}, amount)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.SuccessJson(map[string]interface{}{"tx": call, "function": functionEncode})
}

func (c *ApiContractCallStaticController) Post() {

	function := c.GetString("function")
	params := c.GetString("params")
	address := c.GetString("address")

	//println(string(paramsArr[0]))
	//println(string(paramsArr[1]))

	call, functionEncode, err := models.CallStaticContractFunction(address, "ct_hM2PJEB66Sqx2mkyCixbh3z9hLMaK8N1Sa8v5kaWRqXwPYgkQ", function, []string{params})
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

func (c *ApiContractBalanceController) Post() {

	ctId := c.GetString("ct_id")
	address := c.GetString("address")

	if ctId == "" {
		ctId = "ct_hM2PJEB66Sqx2mkyCixbh3z9hLMaK8N1Sa8v5kaWRqXwPYgkQ"
	}
	//println(string(paramsArr[0]))
	//println(string(paramsArr[1]))

	result, _, err := models.CallStaticContractFunction(address, "ct_hM2PJEB66Sqx2mkyCixbh3z9hLMaK8N1Sa8v5kaWRqXwPYgkQ", "getTokenCallerBalance", []string{address})

	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	switch result.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
	case string:
		c.SuccessJson(map[string]interface{}{"balance": "0.00000"})
	case map[string]interface{}:
		data := result.(map[string]interface{})
		balances := data["Some"].([]interface{})
		balance64, _ := balances[0].(json.Number).Float64()
		balance := utils.FormatTokens(balance64, 5)
		if balance == "0" {
			c.SuccessJson(map[string]interface{}{"balance": "0.00000"})
			return
		}
		c.SuccessJson(map[string]interface{}{"balance": balance})
	}
}

func (c *ApiContractInfoController) Post() {

	ctId := c.GetString("ct_id")
	address := c.GetString("address")

	if ctId == "" {
		ctId = "ct_hM2PJEB66Sqx2mkyCixbh3z9hLMaK8N1Sa8v5kaWRqXwPYgkQ"
	}

	contractResult, _, err := models.CallStaticContractFunction("ak_qJZPXvWPC7G9kFVEqNjj9NAmwMsQcpRu6E3SSCvCQuwfqpMtN", "ct_hM2PJEB66Sqx2mkyCixbh3z9hLMaK8N1Sa8v5kaWRqXwPYgkQ", "getContractBalance", []string{})

	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	myResult, _, err2 := models.CallStaticContractFunction(address, "ct_hM2PJEB66Sqx2mkyCixbh3z9hLMaK8N1Sa8v5kaWRqXwPYgkQ", "getAccountsHeight", []string{address})


	contractBalance64, _ := contractResult.(json.Number).Float64()
	contractBalance := utils.FormatTokens(contractBalance64, 5)
	if contractBalance == "0" {
		contractBalance = "0.00000"
	}
	if err2 != nil {
		if "Error: Account not found" == err2.Error(){
			c.SuccessJson(map[string]interface{}{"contract_balance": contractBalance, "my_balance": "0.00000"})
			return
		}
		c.ErrorJson(-500, err2.Error(), JsonData{})
		return
	}

	var myBalance = "0.00000"
	switch v := myResult.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
	case int:
		fmt.Printf("%v is an int", v)
	case string:
		fmt.Printf("%v is string", v)
		//c.SuccessJson(map[string]interface{}{"balance": "0.00000"})

	case map[string]interface{}:
		data := myResult.(map[string]interface{})
		balance64, _ := data["count"].(json.Number).Float64()
		balance := utils.FormatTokens(balance64, 5)
		myBalance = balance
	}

	if myBalance == "0" {
		myBalance = "0.00000"
	}

	c.SuccessJson(map[string]interface{}{"contract_balance": contractBalance, "my_balance": myBalance})
}
func (c *ApiContractRecordController) Post() {

	ctId := c.GetString("ct_id")
	address := c.GetString("address")

	if ctId == "" {
		ctId = "ct_hM2PJEB66Sqx2mkyCixbh3z9hLMaK8N1Sa8v5kaWRqXwPYgkQ"
	}

	myResult, _, err := models.CallStaticContractFunction(address, "ct_hM2PJEB66Sqx2mkyCixbh3z9hLMaK8N1Sa8v5kaWRqXwPYgkQ", "getAccountsHeight", []string{address})

	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	blockHeight := models.ApiBlocksTop()
	switch myResult.(type) {
	case map[string]interface{}:
		data := myResult.(map[string]interface{})
		heights, _ := data["heights"].([]interface{})
		var items []interface{}
		for i := 0; i < len(heights); i++ {
			var item = map[string]interface{}{}
			height := heights[i].([]interface{})
			model := height[1].(map[string]interface{})

			item["account"] = model["account"]
			item["unlock_height"] = model["unlock_height"]
			item["continue_height"] = model["continue_height"]
			item["day"] = model["day"]
			item["height"] = blockHeight + 100000
			number, _ := model["number"].(json.Number).Float64()
			tokenNumber, _ := model["token_number"].(json.Number).Float64()
			item["number"] = utils.FormatTokens(number, 2)
			item["token_number"] = utils.FormatTokens(tokenNumber, 5)

			items = append(items, item)
		}
		if items == nil {
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(items)
	}
}

func (c *ApiContractLockController) Post() {

	signingKey := c.GetString("signingKey")
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

	call, functionEncode, err := models.CallContractFunction(account, "ct_hM2PJEB66Sqx2mkyCixbh3z9hLMaK8N1Sa8v5kaWRqXwPYgkQ", "lock", []string{params}, amount)
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
