package controllers

import (
	"box/models"
	"box/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
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
type ApiContractBalanceController struct {
	BaseController
}
type ApiContractInfoController struct {
	BaseController
}

type ApiContractRankingController struct {
	BaseController
}

type DefiStatusController struct {
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
		c.ErrorJson(-500, "ERROR", JsonData{})
		return
	}
	switch result.(type) {
	case map[string]interface{}:
		data := result.(map[string]interface{})
		account := data["account"].(string)
		count, _ := data["count"].(json.Number).Float64()
		countFormat := utils.FormatTokens(count, 2)
		if countFormat == "0" {
			countFormat = "0.00"
		}

		token, _ := data["token"].(json.Number).Float64()
		tokenFormat := utils.FormatTokens(token, 7)
		if tokenFormat == "0" {
			tokenFormat = "0.0000000"
		}

		allCount, _ := data["all_count"].(json.Number).Float64()
		allCountFormat := utils.FormatTokens(allCount, 0)

		if allCountFormat == "0" {
			allCountFormat = "0"
		}

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
		return
	}
	c.ErrorJson(-500, "ERROR", JsonData{})
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
	ctId = "ct_7UfopTwsRuLGFEcsScbYgQ6YnySXuyMxQWhw6fjycnzS5Nyzq"
	result, _, err := models.CallStaticContractFunction("ak_2MPzBmtTVXDyBBZALD2JfHrzwdpr8tXZGhu3FRtPJ9sEEPXV2T", ctId, "balance", []string{"ak_nZpU3hfmAfe4g6jiTPPcwa21hnQL68SEYvtizV3iEcfsSHCfD"})
	var output = 0.0

	switch result.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
	case map[string]interface{}:
		data := result.(map[string]interface{})
		balances := data["Some"].([]interface{})
		output, _ = balances[0].(json.Number).Float64()

	}

	resultOwner, _, err := models.CallStaticContractFunction("ak_2MPzBmtTVXDyBBZALD2JfHrzwdpr8tXZGhu3FRtPJ9sEEPXV2T", ctId, "balance", []string{"ak_2VuSVq5ESa5f7HXhqfxn742mexHApSHGd2Erxu2PGxgfdYYmyq"})
	var outputOwner = 0.0

	switch resultOwner.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
	case map[string]interface{}:
		data := resultOwner.(map[string]interface{})
		balances := data["Some"].([]interface{})
		outputOwner, _ = balances[0].(json.Number).Float64()
		output = output + outputOwner

	}
	output = 50000000*1000000000000000000 - output

	myResult, _, err := models.CallStaticContractFunction("ak_2uQYkMmupmAvBtSGtVLyua4EmcPAY62gKo4bSFEmfCNeNK9THX", ctId, "balances", []string{})
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

			if data[i].([]interface{})[0].(string) == "ak_2VuSVq5ESa5f7HXhqfxn742mexHApSHGd2Erxu2PGxgfdYYmyq" ||
				data[i].([]interface{})[0].(string) == "ak_GUpbJyXiKTZB1zRM8Z8r2xFq26sKcNNtz6i83fvPUpKgEAgjH" ||
				data[i].([]interface{})[0].(string) == "ak_2Xu6d6W4UJBWyvBVJQRHASbQHQ1vjBA7d1XUeY8SwwgzssZVHK" ||
				data[i].([]interface{})[0].(string) == "ak_nZpU3hfmAfe4g6jiTPPcwa21hnQL68SEYvtizV3iEcfsSHCfD" {
				continue
			}

			var item = Ranking{}
			item.Address = data[i].([]interface{})[0].(string)
			count, _ := data[i].([]interface{})[1].(json.Number).Float64()
			item.Count = utils.FormatTokens(count, 5)

			item.Proportion = strconv.FormatFloat(count/output*100, 'f', 2, 64)

			items = append(items, item)
		}

		sort.Sort(RankingSlice(items))

		if len(items) >= 100 {
			c.SuccessJson(map[string]interface{}{"out_count": utils.FormatTokens(output, 5), "ranking": items[:100]})
		} else {
			c.SuccessJson(map[string]interface{}{"out_count": utils.FormatTokens(output, 5), "ranking": items})
		}

	}
}

func (c *DefiStatusController) Get() {

	mappingAccounts, _, _ := models.CallStaticContractFunction("ak_2uQYkMmupmAvBtSGtVLyua4EmcPAY62gKo4bSFEmfCNeNK9THX", models.ABCLockContractV3, "get_mapping_accounts", []string{})
	var mappingAccountsList []string = []string{}
	switch mappingAccounts.(type) {
	case []interface{}:
		data := mappingAccounts.([]interface{})
		for i := 0; i < len(data); i++ {
			account := data[i].([]interface{})[1].(map[string]interface{})
			address := account["account"].(string)
			mappingCount, _ := account["mapping_count"].(json.Number).Float64()
			mappingCountFormat := utils.FormatTokens(mappingCount, 2)
			mappingAccountsList = append(mappingAccountsList, address+","+mappingCountFormat)
		}
	}

	getAccountsBlacklists, _, _ := models.CallStaticContractFunction("ak_2uQYkMmupmAvBtSGtVLyua4EmcPAY62gKo4bSFEmfCNeNK9THX", models.ABCLockContractV3, "get_accounts_blacklists", []string{})
	var accountBlacklistList []string = []string{}
	switch getAccountsBlacklists.(type) {
	case []interface{}:
		data := getAccountsBlacklists.([]interface{})
		for i := 0; i < len(data); i++ {
			address := data[i].([]interface{})[0].(string)
			accountBlacklistList = append(accountBlacklistList, address)
		}
	}

	statusResult, _, _ := models.CallStaticContractFunction("ak_2uQYkMmupmAvBtSGtVLyua4EmcPAY62gKo4bSFEmfCNeNK9THX", models.ABCLockContractV3, "get_status", []string{})

	blocksTop := models.ApiBlocksTop()
	c.SuccessJson(map[string]interface{}{
		"last_height":         models.LastHeight,
		"blocks_top":          blocksTop,
		"lock_account_size":   models.LockAccountSize,
		"call_time":           models.ConsumingTime,
		"accounts_blacklists": accountBlacklistList,
		"mapping_accounts":    mappingAccountsList,
		"contract_state":      statusResult,
	})

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
