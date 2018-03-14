package controllers

import (
	"file_server/lib/graphics-go/graphics"
	"fmt"
	"github.com/astaxie/beego"
	"image"
	_ "image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	beego.Info("test")
	this.TplName = "file_test.html"
}

func (this *MainController) Post() {

	filePath := this.Input().Get("path")
	filePath_small := filePath + "_small"
	token := this.Input().Get("token")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	beego.Info("path --- " + filePath)
	beego.Info("filePath_small --- " + filePath_small)
	beego.Info("token --- " + token)

	// 获取附件
	f, h, ee := this.GetFile("file")
	err := os.MkdirAll("files/"+filePath, os.ModePerm)

	if ee != nil {
		beego.Error(ee)
		this.Data["json"] = map[string]interface{}{"status": 400, "message": ee.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	var attachment string
	if h != nil {
		//保存附件
		attachment = h.Filename
		beego.Info("filenewm --- " + attachment)
		files := path.Join("files"+filePath+"/", attachment)
		beego.Info("url --- " + files)
		// err = this.SaveToFile("file", files)

		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		err, l, s := this.SaveFile(f, h, "file", files)
		beego.Info(l)
		beego.Info(s)
		if err != nil {
			s = l
			beego.Error(err)
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "file_url": files + "/" + l, "small_file_url": files + "/" + s, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

}

//上传图片,返回值为原图及缩略图路径
func (this *MainController) SaveFile(f multipart.File, h *multipart.FileHeader, par_name string, filePath string) (err error, l string, s string) {

	defer f.Close()
	upload_path := filePath
	dir := upload_path
	dirf := ""
	// dirf := strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day()) + "/"
	err = os.MkdirAll(dir, 0755)

	if err != nil {
		beego.Info(err)
		return
	}
	filename := h.Filename
	newname := this.FormatFileName(filename)
	imgpath := dir + "/" + newname
	err = this.SaveToFile(par_name, imgpath)

	if err != nil {
		beego.Info(err)
		return
	} else {
		width_s := 40
		height_s := 40
		// height_s, _ := strconv.Atoi(image["height_s"])
		_, err = MakeSmallThumb(imgpath, width_s, height_s)
		if err != nil {
			beego.Info(err)

		}

		l = dirf + newname
		s = dirf + strings.Replace(newname, "_l.", "_s.", 1)
		return err, l, s

	}

}

//small image
func MakeSmallThumb(orgpath string, width int, height int) (string, error) {

	src, err := LoadImage(orgpath)

	if err != nil {
		return "", err
	}
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	err = graphics.Scale(dst, src) //缩小图片

	if err != nil {
		return "", err
	}
	filepath_s := strings.Replace(orgpath, "_l.", "_s.", 1)
	err = SaveImage(filepath_s, dst)
	return filepath_s, err
}

//读取文件
func LoadImage(path string) (img image.Image, err error) {

	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()
	img, _, err = image.Decode(file) //解码图片
	fmt.Println(err)
	return

}

//保存文件
func SaveImage(path string, img image.Image) (err error) {

	imgfile, err := os.Create(path)

	defer imgfile.Close()
	err = png.Encode(imgfile, img) //编码图片
	return
}

//重新生成文件名
func (this *MainController) FormatFileName(filename string) string {

	filepath := strings.Split(filename, ".")
	beego.Info(filepath)
	newname := filepath[0] + "_l." + filepath[1]
	return newname
}
