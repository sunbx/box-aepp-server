package controllers

import (
	. "box/models"
	"box/utils"
	"encoding/json"
	"fmt"
	"github.com/aeternity/aepp-sdk-go/config"
	"github.com/aeternity/aepp-sdk-go/naet"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
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
type ApiVersionController struct {
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
type ApiContractTransferController struct {
	BaseController
}
type TESTController struct {
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
type ApiContractRankingController struct {
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




func (c *TESTController) Get() {
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
	split := strings.Split(mnemonic, " ")
	for i := 0; i < len(split); i++ {
		if len(split[i])<=1{
			c.ErrorJson(-500, "mnemonic error", JsonData{})
			return
		}
	}
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
func (c *ApiVersionController) Post() {
	source, _ := ioutil.ReadFile("conf/version")
	c.Ctx.WriteString(string(source))
}

func (c *ApiContractDecideController) Post() {
	hash := c.GetString("hash")
	function := c.GetString("function")
	ctID := c.GetString("ct_id")
	response := utils.Get(NodeURL + "/v2/transactions/" + hash + "/info")
	var callInfoResult CallInfoResult
	err := json.Unmarshal([]byte(response), &callInfoResult)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	if callInfoResult.Reason == "Tx not mined" {
		c.ErrorJson(-1, "Tx not mined", JsonData{})
		return
	}
	if ctID == "" {
		ctID = ContractBoxAddress
	}

	compile := naet.NewCompiler("http://localhost:3080", false)

	var source []byte
	if ctID == ContractBoxAddress {
		source, _ = ioutil.ReadFile("contract/BoxContract.aes")
	} else if ctID == ContractBoxOldAddress {
		source, _ = ioutil.ReadFile("contract/BoxContractOld.aes")
	} else if ctID == ContractABCAddress {
		source, _ = ioutil.ReadFile("contract/AbcContract.aes")
	} else {
		source, _ = ioutil.ReadFile("contract/BoxContractOld.aes")
	}

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
	ctID := c.GetString("ct_id")
	params := c.GetString("params")
	amount, _ := c.GetFloat("amount", 0)

	if signingKey == "" {
		c.ErrorJson(-500, "signingKey error", JsonData{})
		return
	}
	if ctID == "" {
		ctID = ContractBoxAddress
	}

	if function != "lock" && function != "unlock" && function != "continue_lock" {
		c.ErrorJson(-500, "function error", JsonData{})
		return
	}

	if "lock" == function && amount == 0 {
		c.ErrorJson(-500, "amount error", JsonData{})
		return
	}

	account, err := SigningKeyHexStringAccount(signingKey)
	if err != nil {
		c.ErrorJson(-500, err.Error()+"123123", JsonData{})
		return
	}
	if amount > 0 {

		accountNet, err := ApiGetAccount(account.Address)
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

	call, functionEncode, err := CallContractFunction(account, ctID, function, []string{params}, amount)
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

	call, functionEncode, err := CallStaticContractFunction(address, ContractBoxAddress, function, []string{params})
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
		ctId = ContractBoxAddress
	}

	result, _, err := CallStaticContractFunction(address, ContractBoxAddress, "getTokenCallerBalance", []string{address})

	if err != nil {
		if "Error: Account not found" == err.Error() {
			c.SuccessJson(map[string]interface{}{"balance": "0.00000"})
			return
		}
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
		ctId = ContractBoxAddress
	}
	var contractBalance64Old = 0.0
	var myResultOldCount = 0.0

	contractResultOld, _, err := CallStaticContractFunction("ak_2uQYkMmupmAvBtSGtVLyua4EmcPAY62gKo4bSFEmfCNeNK9THX", ContractBoxOldAddress, "getContractBalance", []string{})
	contractBalance64Old, _ = contractResultOld.(json.Number).Float64()

	if strings.Contains("ak_2g2yq6RniwW1cjKRu4HdVVQXa5GQZkBaXiaVogQXnRxUKpmhS\",270824000000000000000],	[\"ak_3i4bwAbXBRHBqTDYFVLUSa8byQUeBAFzEgjfYk6rSyjWEXL3i\",259200000000000000000],	[\"ak_9XhfcrCtEyPFWPM3GVPC2BCFqetcYV3fDv3EjPpVdR9juAofA\",129600000000000000000],	[\"ak_ELsVMRbBe4LWEuqNU1pn2UCNpnNfdpHjRJjDFjT4R4yzRTeXt\",1390979520000000015854],	[\"ak_Evidt2ZUPzYYPWhestzpGsJ8uWzB1NgMpEvHHin7GCfgWLpjv\",499977516107119999999972654],	[\"ak_GUpbJyXiKTZB1zRM8Z8r2xFq26sKcNNtz6i83fvPUpKgEAgjH\",0],	[\"ak_QyFYYpgJ1vUGk1Lnk8d79WJEVcAtcfuNHqquuP2ADfxsL6yKx\",321088000000000000000],	[\"ak_V9SApNmgDGNLQcZWTzYb3PKtmFuwRn8ENdAg7WjZUdiwgkyUP\",84384000000000000000],	[\"ak_XtJGJrJuvxduT1HFMye4PuEkfUnU9L5rUE5CQ2F9MkqYQVr3f\",648000000000000000000],	[\"ak_fGPGYbqkEyWMV8R4tvQZznpzt28jb54EinF84TRSVCi997kiJ\",2448000000000000000],	[\"ak_o27hkgCTN2WZBkHd4vPcbfJPM2tzddv8xy1yaQnoyFEvqpZQK\",3596400000000000000],	[\"ak_tM5FE5HZSxUvDNAcBKMpSM9iXdsLviJ6tXffiH3BNpFrvgRoR\",383304960000000000000],	[\"ak_22HBW4s8HoCSa6ZKkd7CtFhs7vdBQ5Sgahi7FbRhp7xQ429WG2\",301216320000000007927],	[\"ak_25rsqRgVpcaD3fSZxCQVcyi4VNK3CTqf8CbzsnGtHCeu3ivrM1\",842670000000000000000],	[\"ak_281fyU5kV5yG6ZEgV9nnprLxRznSUKzxmgn2ZnxBhfD8ryWcuk\",128952000000000000000],	[\"ak_28LuZ8CG4LF6LvL47seA2GuCtaNEdXKiVMZP46ykYW8bEcuoVg\",13219200000000000000000],	[\"ak_294D9LQa95ckuJi5z7Who4TzKZWwEGimsyv1ZKM7osPE9c8Bx7\",521424000000000000000],	[\"ak_2JJNMYcnqPaABiSY5omockmv4cCoZefv4XzStAxKe9gM2xYz2r\",582912000000000000000],	[\"ak_2MHJv6JcdcfpNvu4wRDZXWzq8QSxGbhUfhMLR7vUPzRFYsDFw6\",977560560000000001188],	[\"ak_2UCUD59aWZyyhZzZbUdxoyP94r3mz9GvkH49HzJjsfC8MYqVPn\",81000000000000000000],	[\"ak_2Xu6d6W4UJBWyvBVJQRHASbQHQ1vjBA7d1XUeY8SwwgzssZVHK\",1955121120000000002377],	[\"ak_2gEL91xaQwvdN7psiCcGpSwcEMctTX1CVMT2g8f6NEp48tkvAr\",133164000000000000000],	[\"ak_2j2iyGwDnmiDZC9Dc2T8W371MYD9CQxDGSZ2Ne7WT2thY6q888\",213984000000000000000],	[\"ak_2mhBmzVv82SvtKATNBxfD1JhbLBrRNZZmah3QMqRkcK1SP3Bka\",33264000000000000000]", address) {
		myResultOld, _, _ := CallStaticContractFunction(address, ContractBoxOldAddress, "getAccountsHeight", []string{address})
		data := myResultOld.(map[string]interface{})
		myResultOldCount, _ = data["count"].(json.Number).Float64()
	}

	contractResult, _, err := CallStaticContractFunction("ak_2uQYkMmupmAvBtSGtVLyua4EmcPAY62gKo4bSFEmfCNeNK9THX", ContractBoxAddress, "getContractBalance", []string{})

	if err != nil {
		if "Error: Account not found" == err.Error() {
			c.SuccessJson(map[string]interface{}{"contract_balance": "0.00000", "my_balance": "0.00000"})
			return
		}
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	myResult, _, err2 := CallStaticContractFunction(address, ContractBoxAddress, "getAccountsHeight", []string{address})

	contractBalance64, _ := contractResult.(json.Number).Float64()
	contractBalance := utils.FormatTokens(contractBalance64+contractBalance64Old, 5)
	if contractBalance == "0" {
		contractBalance = "0.00000"
	}
	if err2 != nil {
		if "Error: Account not found" == err2.Error() {
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
		balance := utils.FormatTokens(balance64+myResultOldCount, 5)
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
		ctId = ContractBoxAddress
	}

	myResult, _, err := CallStaticContractFunction(address, ctId, "getAccountsHeight", []string{address})

	if err != nil {
		if "Error: Account not found" == err.Error() {
			c.SuccessJson([]JsonData{})
			return
		}
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	blockHeight := ApiBlocksTop()
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
			item["height"] = blockHeight
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

type RankingSlice []Ranking

func (s RankingSlice) Less(i, j int) bool {
	one, _ := strconv.ParseFloat(s[i].Count, 64)
	two, _ := strconv.ParseFloat(s[j].Count, 64)
	return one > two
}

func (s RankingSlice) Len() int { return len(s) }

func (s RankingSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (c *ApiContractRankingController) Post() {

	ctId := c.GetString("ct_id")

	if ctId == "" {
		ctId = ContractABCAddress
	}

	result, _, err := CallStaticContractFunction("ak_2MPzBmtTVXDyBBZALD2JfHrzwdpr8tXZGhu3FRtPJ9sEEPXV2T", ContractBoxAddress, "getTokenCallerBalance", []string{"ak_2MPzBmtTVXDyBBZALD2JfHrzwdpr8tXZGhu3FRtPJ9sEEPXV2T"})
	var output = 0.0
	switch result.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
	case map[string]interface{}:
		data := result.(map[string]interface{})
		balances := data["Some"].([]interface{})
		output, _ = balances[0].(json.Number).Float64()
		output = 500000000*1000000000000000000 - output
	}

	myResult, _, err := CallStaticContractFunction("ak_2uQYkMmupmAvBtSGtVLyua4EmcPAY62gKo4bSFEmfCNeNK9THX", ctId, "balances", []string{})
	if err != nil {
		if "Error: Account not found" == err.Error() {
			c.SuccessJson([]JsonData{})
			return
		}
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	//c.SuccessJson(myResult)
	//blockHeight := models.ApiBlocksTop()
	switch myResult.(type) {
	case []interface{}:

		data := myResult.([]interface{})
		var items []Ranking
		for i := 0; i < len(data); i++ {

			//if data[i].([]interface{})[0].(string) == "ak_2MPzBmtTVXDyBBZALD2JfHrzwdpr8tXZGhu3FRtPJ9sEEPXV2T" || data[i].([]interface{})[0].(string) == "ak_GUpbJyXiKTZB1zRM8Z8r2xFq26sKcNNtz6i83fvPUpKgEAgjH"  || data[i].([]interface{})[0].(string) == "ak_2Xu6d6W4UJBWyvBVJQRHASbQHQ1vjBA7d1XUeY8SwwgzssZVHK"  || data[i].([]interface{})[0].(string) == "ak_2MHJv6JcdcfpNvu4wRDZXWzq8QSxGbhUfhMLR7vUPzRFYsDFw6" {
			//	continue
			//}

			var item = Ranking{}
			item.Address = data[i].([]interface{})[0].(string)
			count, _ := data[i].([]interface{})[1].(json.Number).Float64()
			item.Count = utils.FormatTokens(count, 5)

			item.Proportion = strconv.FormatFloat(count/output*100, 'f', 2, 64)

			items = append(items, item)
		}


		sort.Sort(RankingSlice(items))

		if len(items)>=100{
			c.SuccessJson(map[string]interface{}{"out_count": utils.FormatTokens(output, 5), "ranking": items[:100]})
		}else{
			c.SuccessJson(map[string]interface{}{"out_count": utils.FormatTokens(output, 5), "ranking": items})
		}

	}
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

type Ranking struct {
	Address    string `json:"address"`
	Count      string `json:"count"`
	OutCount   string `json:"OutCount"`
	Proportion string `json:"proportion"`
}

func (c *ApiContractLockController) Post() {

	signingKey := c.GetString("signingKey")
	params := c.GetString("params")
	amount, _ := c.GetFloat("amount", 0)

	account, err := SigningKeyHexStringAccount(signingKey)
	if amount > 0 {

		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		accountNet, err := ApiGetAccount(account.Address)
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

	call, functionEncode, err := CallContractFunction(account, ContractBoxAddress, "lock", []string{params}, amount)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	c.SuccessJson(map[string]interface{}{"tx": call, "function": functionEncode})
}

func (c *ApiContractTransferController) Post() {

	signingKey := c.GetString("signingKey")
	address := c.GetString("address")
	amount, _ := c.GetFloat("amount", 0)

	account, err := SigningKeyHexStringAccount(signingKey)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	if amount > 0 {

		accountNet, err := ApiGetAccount(account.Address)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		tokens, err := strconv.ParseFloat(accountNet.Balance.String(), 64)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}

		if tokens/1000000000000000000 <= 1 {
			c.ErrorJson(-500, "Keep AE token greater than 1", JsonData{})
			return
		}

	}

	call, functionEncode, err := CallContractFunction(account, ContractABCAddress, "transfer", []string{address, utils.GetRealAebalanceBigInt(amount).String()}, 0)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}

	c.SuccessJson(map[string]interface{}{"tx": call, "function": functionEncode})
}
