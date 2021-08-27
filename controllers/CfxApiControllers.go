package controllers

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

type ApiCfxBalanceController struct {
	BaseController
}
type ApiCfxTokensController struct {
	BaseController
}
type ApiCfxTransactionController struct {
	BaseController
}
type ApiCfxTransactionHashController struct {
	BaseController
}

var CfxHost = "https://confluxscan.io/v1"

func (c *ApiCfxBalanceController) Post() {
	address := c.GetString("address")
	resp, err := http.Get(CfxHost + "/account/" + address)
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

func (c *ApiCfxTokensController) Post() {

	address := c.GetString("address")
	resp, err := http.Get(CfxHost + "/token?fields=icon&accountAddress=" + address)
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

func (c *ApiCfxTransactionHashController) Post() {

	//https://confluxscan.io/v1/transaction/0x214a1c853cabd8d556c016f7095e2ff994ddceae6ed7ecfdc9fd74bcc5e81c44

	hash := c.GetString("hash")
	resp, err := http.Get(CfxHost + "/transaction/" + hash)
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
func (c *ApiCfxTransactionController) Post() {
	address := c.GetString("address")
	page, _ := c.GetInt("page", 1)
	ctAddress := c.GetString("ct_address")

	skip := strconv.Itoa(page*10 - 10)
	var resp *http.Response
	var err error
	if ctAddress == "" {
		resp, err = http.Get(CfxHost + "/transaction?limit=10&accountAddress=" + address + "&skip=" + skip)
	} else {
		resp, err = http.Get(CfxHost + "/transfer?transferType=ERC20&limit=10&accountAddress=" + address + "&skip=" + skip + "&address=" + ctAddress)
	}

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
