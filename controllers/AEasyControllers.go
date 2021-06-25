package controllers

import (
	"box/models"
	"box/utils"
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
type TokenRecordController struct {
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
	//compile := naet.NewCompiler(models.CompilerUrl, false)
	////var source []byte
	////source, _ = ioutil.ReadFile("contract/ABCLockContractV3.aes")
	//
	////decodeResult, err := compile.DecodeCallResult("cb", tryRun.Results[0].CallObj.ReturnValue, function, string(source), config.Compiler.Backend)
	////decodedData, err := compile.DecodeData("cb_KxH7D7nnG2+KqSrUWJVW1Jf/wIyjh+w=", "int")
	//decodedData, err := compile.DecodeCalldataBytecode("cb_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BwEChP7tkXzfAjcDRwBHAAc3AAwBBAIDEZcFecwPAm+CJs8MAQQMAQACAxF2S4FBDwJvgibPKxoUhgAVGBQELRqGhgAs2jiGAgAUGDgELRqGhgIMAQAMAQIMAQRE/FMGBgQEBAAGBAMRZaXgD/7vy1DvAjcC5wCHAjcANwHnAOcADAECDAMrEW9MN6k/DAEABAMRJD/bEf7vzFjhADcCRwAHNwAMAQICAxGXBXnMDwJvgibPVQAMAQAnDAQPAgQMAgQCAxGLx/TGDwJvgibPLWqOjgQCVQAMAQAMAQJE/FMGBgQEBAIGBAMRZaXgD/7+rqT6ADcARwABAoK5AtIvIBEDFHvNsS5GdW5naWJsZVRva2VuRnVsbC5pbnRlcm5hbF9jaGFuZ2VfYWxsb3dhbmNlEREAT6YRc3dhcBEUN7Q4JW1ldGFfaW5mbxEd7Gb/HXN3YXBwZWQRId/6tkl0cmFuc2Zlcl9hbGxvd2FuY2URI8Ncaj1yZXNldF9hbGxvd2FuY2URJD/bETUuT3B0aW9uLm1hdGNoESbcN4UZLl4xMjYwES2dNN6RLkZ1bmdpYmxlVG9rZW5GdWxsLnJlcXVpcmVfYWxsb3dhbmNlETHAjXM9YWV4OV9leHRlbnNpb25zET2FWo5BY2hhbmdlX2FsbG93YW5jZRE+bHgzUWFsbG93YW5jZV9mb3JfY2FsbGVyEUTWRB8RaW5pdBFlpeAPLUNoYWluLmV2ZW50EWoAFl4lYWxsb3dhbmNlEW9MN6kZLl4xMjYxEXZLgUGJLkZ1bmdpYmxlVG9rZW5GdWxsLnJlcXVpcmVfYmFsYW5jZRGAJGtHIWJhbGFuY2VzEYShXaEhdHJhbnNmZXIRi8f0xsUuRnVuZ2libGVUb2tlbkZ1bGwucmVxdWlyZV9hbGxvd2FuY2Vfbm90X2V4aXN0ZW50EZcFecy1LkZ1bmdpYmxlVG9rZW5GdWxsLnJlcXVpcmVfbm9uX25lZ2F0aXZlX3ZhbHVlEbHvwXsRYnVybhG0jBaEHWJhbGFuY2URwnAhIoEuRnVuZ2libGVUb2tlbkZ1bGwucmVxdWlyZV9vd25lchHP3ZqiEW1pbnQR1jkNfiljaGVja19zd2FwEddMFd4pYWxsb3dhbmNlcxHbY3WoMXRvdGFsX3N1cHBseRHtkXzfkS5GdW5naWJsZVRva2VuRnVsbC5pbnRlcm5hbF90cmFuc2ZlchHvy1DvPS5PcHRpb24uZGVmYXVsdBHvzFjhQWNyZWF0ZV9hbGxvd2FuY2UR/q6k+hVvd25lcoIvAIU0LjMuMABInfYD", "cb_cUFDQ09VTlRfSU5TVUZGSUNJRU5UX0JBTEFOQ0Uj8B4f", config.Compiler.Backend)
	////decodedData, err := compile.DecodeCalldataBytecode("cb_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BwEChP7tkXzfAjcDRwBHAAc3AAwBBAIDEZcFecwPAm+CJs8MAQQMAQACAxF2S4FBDwJvgibPKxoUhgAVGBQELRqGhgAs2jiGAgAUGDgELRqGhgIMAQAMAQIMAQRE/FMGBgQEBAAGBAMRZaXgD/7vy1DvAjcC5wCHAjcANwHnAOcADAECDAMrEW9MN6k/DAEABAMRJD/bEf7vzFjhADcCRwAHNwAMAQICAxGXBXnMDwJvgibPVQAMAQAnDAQPAgQMAgQCAxGLx/TGDwJvgibPLWqOjgQCVQAMAQAMAQJE/FMGBgQEBAIGBAMRZaXgD/7+rqT6ADcARwABAoK5AtIvIBEDFHvNsS5GdW5naWJsZVRva2VuRnVsbC5pbnRlcm5hbF9jaGFuZ2VfYWxsb3dhbmNlEREAT6YRc3dhcBEUN7Q4JW1ldGFfaW5mbxEd7Gb/HXN3YXBwZWQRId/6tkl0cmFuc2Zlcl9hbGxvd2FuY2URI8Ncaj1yZXNldF9hbGxvd2FuY2URJD/bETUuT3B0aW9uLm1hdGNoESbcN4UZLl4xMjYwES2dNN6RLkZ1bmdpYmxlVG9rZW5GdWxsLnJlcXVpcmVfYWxsb3dhbmNlETHAjXM9YWV4OV9leHRlbnNpb25zET2FWo5BY2hhbmdlX2FsbG93YW5jZRE+bHgzUWFsbG93YW5jZV9mb3JfY2FsbGVyEUTWRB8RaW5pdBFlpeAPLUNoYWluLmV2ZW50EWoAFl4lYWxsb3dhbmNlEW9MN6kZLl4xMjYxEXZLgUGJLkZ1bmdpYmxlVG9rZW5GdWxsLnJlcXVpcmVfYmFsYW5jZRGAJGtHIWJhbGFuY2VzEYShXaEhdHJhbnNmZXIRi8f0xsUuRnVuZ2libGVUb2tlbkZ1bGwucmVxdWlyZV9hbGxvd2FuY2Vfbm90X2V4aXN0ZW50EZcFecy1LkZ1bmdpYmxlVG9rZW5GdWxsLnJlcXVpcmVfbm9uX25lZ2F0aXZlX3ZhbHVlEbHvwXsRYnVybhG0jBaEHWJhbGFuY2URwnAhIoEuRnVuZ2libGVUb2tlbkZ1bGwucmVxdWlyZV9vd25lchHP3ZqiEW1pbnQR1jkNfiljaGVja19zd2FwEddMFd4pYWxsb3dhbmNlcxHbY3WoMXRvdGFsX3N1cHBseRHtkXzfkS5GdW5naWJsZVRva2VuRnVsbC5pbnRlcm5hbF90cmFuc2ZlchHvy1DvPS5PcHRpb24uZGVmYXVsdBHvzFjhQWNyZWF0ZV9hbGxvd2FuY2UR/q6k+hVvd25lcoIvAIU0LjMuMABInfYD", "cb_KxGEoV2hK58AoA8jmX80DWnUWdKP6BrGYwVRia9m7DwD6P3lE6Kd0+C+b4hIKhxzAAf/wL1wHW0=", config.Compiler.Backend)
	//if err != nil {
	//	c.Ctx.WriteString(string(err.Error()))
	//	return
	//}
	//c.SuccessJson(decodedData)

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

func (c *TokenRecordController) Post() {
	address := c.GetString("address")
	ctId := c.GetString("ct_id")
	page := c.GetString("page")
	resp, err := http.PostForm(HOST+"/api/aex9/record",
		url.Values{
			"app_id":  {beego.AppConfig.String("AEASY::appId")},
			"address": {address},
			"ct_id":   {ctId},
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
		if err.Error() == "Account not found" {
			account, _ := models.SigningKeyHexStringAccount(beego.AppConfig.String("AEASY::accountFoundation"))
			tx, _ := models.ApiSpend(account, address, 0.00001, "Sponsored by China Foundation（中国基金会赞助）")
			print(tx.Hash)

		}
	}

	c.Ctx.WriteString(string(body))
}
func (c *ApiVersionController) Post() {
	source, _ := ioutil.ReadFile("conf/version")
	c.Ctx.WriteString(string(source))
}
