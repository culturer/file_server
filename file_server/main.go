package main

import (
	_ "file_server/routers"
	"github.com/astaxie/beego"
	"os"
)

func main() {
	//创建附件目录
	os.Mkdir("files", os.ModePerm)
	beego.SetStaticPath("files", "files")
	beego.Info("created by song , qq --- 78901214")

	beego.Run()
}
