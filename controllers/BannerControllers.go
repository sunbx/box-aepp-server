package controllers

import (
	"io/ioutil"
)

type BannerController struct {
	BaseController
}


func (c *BannerController) Post() {
	source, _ := ioutil.ReadFile("conf/banner.json")
	c.Ctx.WriteString(string(source))
}
