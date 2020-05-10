package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"time"
)

type BaseController struct {
	beego.Controller
}

type ReturnMsg struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Stime int64       `json:"time"`
	Data  interface{} `json:"data"`
}

type JsonData struct {
}

var bm, _ = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)

func (c *BaseController) SuccessJson(data interface{}) {
	serviceTime := time.Now().UnixNano() / 1e6
	res := ReturnMsg{
		200, "success", serviceTime, data,
	}
	jsons, _ := json.Marshal(res)

	c.Ctx.WriteString(string(jsons))
}

func (c *BaseController) ErrorJson(code int, msg string, data interface{}) {
	serviceTime := time.Now().UnixNano() / 1e6
	res := ReturnMsg{
		code, msg, serviceTime, data,
	}
	jsons, _ := json.Marshal(res)

	c.Ctx.WriteString(string(jsons))
}

