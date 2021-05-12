package controllers

import (
	"box/models"
	"box/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type TokenListController struct {
	BaseController
}

type TokenListStruct []struct {
	Name      string `json:"name"`
	Image     string `json:"image"`
	CtAddress string `json:"ct_address"`
	Type      string `json:"type"`
}

type ApiContractBalanceController struct {
	BaseController
}

func (c *TokenListController) Post() {

	address := c.GetString("address")

	source, _ := ioutil.ReadFile("conf/tokens.json")
	var tokensResult []map[string]string
	err := json.Unmarshal(source, &tokensResult)
	if err != nil {
		fmt.Println("tokens json err: ", err)
	}

	for i := 0; i < len(tokensResult); i++ {
		result, _, _ := models.TokenBalanceFunction(address, tokensResult[i]["ct_address"], tokensResult[i]["type"], "balance", []string{address})
		switch result.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
		case map[string]interface{}:
			data := result.(map[string]interface{})
			balances := data["Some"].([]interface{})
			balanceFloat, _ := balances[0].(json.Number).Float64()
			balanceFloatFormat := strconv.FormatFloat(balanceFloat/1000000000000000000, 'f', 5, 64)
			tokensResult[i]["count"] = balanceFloatFormat
		}

	}
	c.SuccessJson(tokensResult)
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