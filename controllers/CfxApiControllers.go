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
type ApiCfxTokensListController struct {
	BaseController
}
type ApiCfxTokensByAddressController struct {
	BaseController
}
type ApiCfxTransactionController struct {
	BaseController
}
type ApiCfxTransactionHashController struct {
	BaseController
}
type ApiCfxCrc20TransactionHashController struct {
	BaseController
}

type ApiCfxNFTBalanceController struct {
	BaseController
}
type ApiCfxNFTTokenController struct {
	BaseController
}
type ApiCfxNFTPreviewController struct {
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
	resp, err := http.Get(CfxHost + "/token?fields=icon&transferType=ERC20&limit=100&accountAddress=" + address)
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
func (c *ApiCfxTokensListController) Post() {
	resp, err := http.Get("https://confluxscan.io/stat/tokens/list?fields=transferCount&fields=iconUrl&fields=price&fields=totalPrice&fields=quoteUrl&fields=transactionCount&fields=erc20TransferCount&limit=9999&orderBy=totalPrice&reverse=true&skip=0&transferType=ERC20")
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

func (c *ApiCfxTokensByAddressController) Post() {
	address := c.GetString("address")
	resp, err := http.Get("https://confluxscan.io/stat/tokens/by-address?address=" + address + "&fields=iconUrl&fields=transferCount&fields=price&fields=totalPrice&fields=quoteUrl")
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

	skip := strconv.Itoa(page*10 - 10)
	var resp *http.Response
	var err error
	resp, err = http.Get(CfxHost + "/transaction?limit=10&accountAddress=" + address + "&skip=" + skip)

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
func (c *ApiCfxCrc20TransactionHashController) Post() {

	address := c.GetString("address")
	contract := c.GetString("contract")
	page, _ := c.GetInt("page", 1)

	skip := strconv.Itoa(page*10 - 10)
	var resp *http.Response
	var err error
	//
	resp, err = http.Get("https://api.confluxscan.net/account/crc20/transfers?account=" + address + "&skip=" + skip + "&limit=10&contract=" + contract + "&from=" + address + "&to=" + address + "&sort=DESC")

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

func (c *ApiCfxNFTBalanceController) Post() {
	address := c.GetString("address")
	var resp *http.Response
	var err error
	resp, err = http.Get("https://confluxscan.io/stat/nft/checker/balance?ownerAddress=" + address)
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

func (c *ApiCfxNFTTokenController) Post() {
	address := c.GetString("address")
	contract := c.GetString("contract")
	var resp *http.Response
	var err error
	resp, err = http.Get("https://confluxscan.io/stat/nft/checker/token?contractAddress=" + contract + "&limit=12&ownerAddress=" + address + "&skip=0")
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

func (c *ApiCfxNFTPreviewController) Post() {
	tokenId := c.GetString("tokenId")
	contract := c.GetString("contract")
	var resp *http.Response
	var err error
	resp, err = http.Get("https://confluxscan.io/stat/nft/checker/preview?contractAddress=" + contract + "&tokenId=" + tokenId)
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
