package controllers

import (
	"box/models"
	"box/utils"
	"encoding/json"
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
type ApiContractBalanceController struct {
	BaseController
}
type ApiContractInfoController struct {
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
			account, _ := models.SigningKeyHexStringAccount(beego.AppConfig.String("AEASY::accountFoundation"))
			tx, _ := models.ApiSpend(account, address, 0.001, "Sponsored by China Foundation（中国基金会赞助）")
			print(tx.Hash)

		}
	}

	c.Ctx.WriteString(string(body))
}
func (c *ApiVersionController) Post() {
	source, _ := ioutil.ReadFile("conf/version")
	c.Ctx.WriteString(string(source))
}

func (c *ApiContractBalanceController) Post() {
	ctId := c.GetString("ct_id")
	address := c.GetString("address")
	result, _, err := models.CallStaticContractFunction(address, ctId, "balance", []string{address})
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
	address := c.GetString("address")
	result, _, err := models.CallStaticContractFunction(address, models.ABCLockContractV3, "get_data_info", []string{address})
	if err != nil {
		if "Error: Account not found" == err.Error() {
			c.SuccessJson(map[string]interface{}{
				"account":      address,
				"after_height": -1,
				"all_count":    -1,
				"count":        -1,
				"height":       -1,
				"min_height":   -1,
				"token":        -1,
			})
			return
		}
	}
	switch result.(type) {
	case map[string]interface{}:
		data := result.(map[string]interface{})
		account := data["account"].(string)
		count, _ := data["count"].(json.Number).Float64()
		countFormat := utils.FormatTokens(count, 5)

		token, _ := data["token"].(json.Number).Float64()
		tokenFormat := utils.FormatTokens(token, 7)

		allCount, _ := data["all_count"].(json.Number).Float64()
		allCountFormat := utils.FormatTokens(allCount, 5)

		height := data["height"].(json.Number)
		afterHeight := data["after_height"].(json.Number)
		minHeight := data["min_height"].(json.Number)
		c.SuccessJson(map[string]interface{}{
			"account":      account,
			"count":        countFormat,
			"token":        tokenFormat,
			"all_count":    allCountFormat,
			"height":       height,
			"after_height": afterHeight,
			"min_height":   minHeight,
		})
	}
	c.ErrorJson(-500, "ERROR", JsonData{})
}
