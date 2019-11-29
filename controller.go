package main

import (
	"directory-lister-echo/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func Test(c echo.Context) error {
	//这边有个地方值得注意，template.New()函数中参数名字要和ParseFiles（）
	//函数的文件名要相同，要不然就会报错："" is an incomplete template
	//tmpl := template.New("test.html")
	//tmpl = tmpl.Funcs(template.FuncMap{"EqJudge": EqJudge})
	//tmpl, _ = tmpl.ParseFiles("test.html")
	path := c.QueryParam("path")
	fmt.Println(len(utils.PathSplitter(path, "Bajins Soft")))
	return c.Render(http.StatusOK, "test.html", map[string]interface{}{
		"name":        "Dolly!",
		"web_title":   "Bajins",
		"breadcrumbs": utils.PathSplitter(path, "Bajins Soft"),
	})
}

func GetDirList(root, path string) []map[string]interface{} {
	if utils.IsExistDir(root) {
		return nil
	}
	// 获取目录下的文件和子目录信息
	list := utils.GetFileList(root + path)
	if list == nil {
		return nil
	}
	// 创建切片
	var dir []map[string]interface{}
	for _, info := range list {
		// 创建map
		m := make(map[string]interface{})

		m["name"] = info.Name()
		if !info.IsDir() {
			// Size()单位为Byte,所以要转换
			m["size"] = utils.ByteSize(uint64(info.Size()))
		} else {
			m["size"] = "-"
		}
		// 时间
		m["modTime"] = utils.TimeToString(info.ModTime())
		// 权限
		//m["mode"] = info.Mode().String()
		m["path"] = strings.ReplaceAll(path+"\\"+info.Name(), "\\", "/")
		m["isDir"] = info.IsDir()
		// 放进切片中
		dir = append(dir, m)
	}
	return dir
}

// 获取展示目录下的所有
func GetDir(c echo.Context) error {
	// 页
	/*lows := c.QueryParam("low")
	var low int
	var err error
	if lows != "" {
		low, err = strconv.Atoi(lows)
		if err != nil {
			return Error(c, 300, "请传入参数low")
		}
	}
	// 数量
	highs := c.QueryParam("high")
	var high int
	if highs != "" {
		high, err = strconv.Atoi(highs)
		if err != nil {
			return Error(c, 300, "请传入参数high")
		}
	}*/

	toPath := c.QueryParam("path")
	//config := services.ConfigService.GetByType("filePath")
	//var path string
	//if config != nil && utils.IsExistDir(config[0].Value) {
	//	path = config[0].Value
	//} else {
	//	path = utils.OsPath()
	//}
	d, _ := os.Getwd()
	dir := GetDirList(d, toPath)
	//total := len(dir)
	//if lows != "" && highs != "" {
	//	low = (low - 1) * high
	//	end := low + high
	//	if total < end {
	//		end = total
	//	}
	//	dir = dir[low:end]
	//}
	data := make(map[string]interface{})
	//data["total"] = total
	data["file"] = dir
	links := utils.PathSplitter(toPath, "Bajins Soft")
	data["links"] = links
	return Success(c, "获取文件列表成功", data)
}

// 下载文件
func DownloadFile(c echo.Context) error {
	fmt.Println("============", c.QueryParams())
	filePath := c.QueryParam("filePath")
	if filePath == "" {
		//this.Ctx.WriteString("目录路径不正确")
		return Error(c, 402, "目录路径不正确")
	}
	filePath = path.Join(utils.OsPath(), filePath)
	fmt.Println(filePath)
	if !utils.IsFile(filePath) {
		return Error(c, 402, "目录路径不正确")
	}
	fileName := filepath.Base(filePath)
	//c.Response().Header().Set("Content-Type","application/x-msdownload")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+fileName)
	//第一个参数是文件的地址，第二个参数是下载显示的文件的名称
	return c.File(filePath)
}

func EqJudge(obj []interface{}, index int) bool {
	if len(obj)-1 != index {
		return true
	}
	return false
}

func MapGetValue(m map[string]interface{}, key string) string {
	fmt.Println(m[key])
	return m[key].(string)
}
