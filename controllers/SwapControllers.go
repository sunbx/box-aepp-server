package controllers

import (
	"box/models"
	"box/utils"
	"encoding/json"
	"io/ioutil"
	"sort"
	"strconv"
)

type SwapCoinListController struct {
	BaseController
}

type SwapCoinAccountController struct {
	BaseController
}

type SwapCoinAccountMyController struct {
	BaseController
}
type SwapCoinOrderMyController struct {
	BaseController
}

type SwapAccount struct {
	Account    string  `json:"account"`
	CoinName   string  `json:"coin_name"`
	AeCount    string  `json:"ae_count"`
	TokenCount string  `json:"token_count"`
	Token      string  `json:"token"`
	Rate       string `json:"rate"`
}

type SwapAccountSlice []SwapAccount

func (s SwapAccountSlice) Less(i, j int) bool {
	oneAeCount, _ := strconv.ParseFloat(s[i].AeCount, 64)
	oneTokenCount, _ := strconv.ParseFloat(s[i].TokenCount, 64)
	oneRate := oneAeCount / oneTokenCount

	twoAeCount, _ := strconv.ParseFloat(s[j].AeCount, 64)
	twoTokenCount, _ := strconv.ParseFloat(s[j].TokenCount, 64)

	twoRate := twoTokenCount / twoAeCount
	return oneRate > twoRate
}

func (s SwapAccountSlice) Len() int { return len(s) }

func (s SwapAccountSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (c *SwapCoinListController) Post() {
	source, _ := ioutil.ReadFile("conf/swap.json")
	c.Ctx.WriteString(string(source))
}

func (c *SwapCoinAccountController) Post() {
	ctId := c.GetString("ct_id")
	contractResult, _, e := models.CallStaticContractFunction("ak_2uQYkMmupmAvBtSGtVLyua4EmcPAY62gKo4bSFEmfCNeNK9THX", models.BoxSwapContractV2, "get_swap_accounts", []string{ctId})
	if e != nil {
		c.SuccessJson(e.Error())
		return
	}
	switch contractResult.(type) {
	case map[string]interface{}:
		contractResultMap := contractResult.(map[string]interface{})

		//transaction := contractResultMap["transaction"]
		coinAccount := contractResultMap["coin_account"]

		coinAccountArray := coinAccount.([]interface{})

		var coinAccountMap []SwapAccount = []SwapAccount{}
		for i := 0; i < len(coinAccountArray); i++ {
			accountArr := coinAccountArray[i].([]interface{})
			account := accountArr[1].(map[string]interface{})

			var item = SwapAccount{}
			item.Account = account["account"].(string)
			item.CoinName = account["coin_name"].(string)

			aeCount, _ := account["ae_count"].(json.Number).Float64()
			item.AeCount = utils.FormatTokens(aeCount, 2)

			tokenCount, _ := account["token_count"].(json.Number).Float64()
			item.TokenCount = utils.FormatTokens(tokenCount, 2)
			if aeCount/tokenCount>1 {
				item.Rate = strconv.FormatFloat(aeCount/tokenCount, 'f', 2, 64)
			}else{
				item.Rate = strconv.FormatFloat(aeCount/tokenCount, 'f', 4, 64)
			}

			item.Token = account["token"].(string)
			coinAccountMap = append(coinAccountMap, item)
		}
		sort.Sort(SwapAccountSlice(coinAccountMap))
		c.SuccessJson(coinAccountMap)
		return
	}

	c.SuccessJson(contractResult)
}

func (c *SwapCoinAccountMyController) Post() {
	address := c.GetString("address")
	contractResult, _, e := models.CallStaticContractFunction("ak_2uQYkMmupmAvBtSGtVLyua4EmcPAY62gKo4bSFEmfCNeNK9THX", models.BoxSwapContractV2, "get_coin_account_map", []string{address})
	if e != nil {
		c.SuccessJson(e.Error())
		return
	}
	switch contractResult.(type) {
	case []interface{}:
		coinAccountArray := contractResult.([]interface{})

		var coinAccountMap []SwapAccount = []SwapAccount{}
		for i := 0; i < len(coinAccountArray); i++ {
			accountArr := coinAccountArray[i].([]interface{})
			account := accountArr[1].(map[string]interface{})

			var item = SwapAccount{}
			item.Account = account["account"].(string)
			item.CoinName = account["coin_name"].(string)

			aeCount, _ := account["ae_count"].(json.Number).Float64()
			item.AeCount = utils.FormatTokens(aeCount, 2)

			tokenCount, _ := account["token_count"].(json.Number).Float64()

			item.TokenCount = utils.FormatTokens(tokenCount, 2)

			if aeCount/tokenCount>1 {
				item.Rate = strconv.FormatFloat(aeCount/tokenCount, 'f', 2, 64)
			}else{
				item.Rate = strconv.FormatFloat(aeCount/tokenCount, 'f', 4, 64)
			}
			item.Token = account["token"].(string)
			coinAccountMap = append(coinAccountMap, item)
		}
		sort.Sort(SwapAccountSlice(coinAccountMap))
		c.SuccessJson(coinAccountMap)
		return
	}

	c.SuccessJson(contractResult)
}
func (c *SwapCoinOrderMyController) Post() {
	address := c.GetString("address")
	contractResult, _, e := models.CallStaticContractFunction("ak_2uQYkMmupmAvBtSGtVLyua4EmcPAY62gKo4bSFEmfCNeNK9THX", models.BoxSwapContractV2, "get_accounts_records", []string{address})
	if e != nil {
		c.SuccessJson(e.Error())
		return
	}
	height := models.ApiBlocksTop()
	switch contractResult.(type) {
	case []interface{}:
		orderArray := contractResult.([]interface{})

		for i := 0; i < len(orderArray); i++ {
			account := orderArray[i].(map[string]interface{})

			aeCount, _ := account["ae_count"].(json.Number).Float64()
			account["ae_count"] = utils.FormatTokens(aeCount, 2)

			tokenCount, _ := account["token_count"].(json.Number).Float64()
			account["token_count"] = utils.FormatTokens(tokenCount, 2)
			account["current_height"] = height

		}
		c.SuccessJson(orderArray)
		return
	}

	c.SuccessJson(contractResult)
}
