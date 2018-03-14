package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type ImgController struct {
	beego.Controller
}

func (this *ImgController) Get() {
	beego.Info("get Img")
	this.TplName = "file_test.html"
}

func (this *ImgController) Post() {

	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	url := this.Input().Get("url")
	token := this.Input().Get("token")
	imgWidth, _ := strconv.Atoi(this.Input().Get("imgWidth"))
	imgHeight, _ := strconv.Atoi(this.Input().Get("imgHeight"))

	beego.Info("url --- " + url)
	beego.Info(imgWidth)
	beego.Info(imgHeight)
	beego.Info("token --- " + token)

	mUrl, _ := MakeSmallThumb(url, imgWidth, imgHeight)
	beego.Info(mUrl)
	this.Data["json"] = map[string]interface{}{"status": 200, "file_url": mUrl, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}
