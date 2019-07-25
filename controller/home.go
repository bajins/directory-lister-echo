package controller

import (
	"directory-lister-echo/result"
	"directory-lister-echo/utils"
	"fmt"
	"github.com/labstack/echo"
	"strconv"
	"strings"
)

/**
获取展示目录下的所有
*/
func GetDirectory(c echo.Context) error {
	// 页
	lows := c.QueryParam("low")
	low, err := strconv.Atoi(lows)
	if err != nil {
		return result.Error(c, 300, "请传入参数low")
	}
	// 数量
	highs := c.QueryParam("high")
	high, err := strconv.Atoi(highs)
	if err != nil {
		return result.Error(c, 300, "请传入参数high")
	}

	path := c.QueryParam("path")
	if path == "" {
		path = "/"
	}
	//config := services.ConfigService.GetByType("filePath")
	//var path string
	//if config != nil && utils.IsExistDir(config[0].Value) {
	//	path = config[0].Value
	//} else {
	//	path = utils.OsPath()
	//}
	// 获取目录下的文件和子目录信息
	list := utils.GetFileList("F:\\directory-lister-echo" + path)
	if list == nil {
		return result.Error(c, 402, "目录路径不正确")
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
		m["path"] = strings.Replace(path+"\\"+info.Name(), "\\", "/", -1)
		m["isDir"] = info.IsDir()
		// 放进切片中
		dir = append(dir, m)
	}

	total := len(dir)
	low = (low - 1) * high
	end := low + high
	if total < end {
		end = total
	}
	dir = dir[low:end]
	data := make(map[string]interface{})
	data["total"] = total
	data["file"] = dir
	links := utils.PathSplitter(path, "Bajins Soft")
	data["links"] = links
	if len(links) == 1 {
		data["parent"] = links[0]
	} else {
		data["parent"] = links[len(links)-2]
	}
	return result.Success(c, "获取文件列表成功", data)
}

/**
 * 下载文件
 *
 * @param null
 * @return
 * @Description
 * @author claer www.bajins.com
 * @date 2019/6/25 15:37
 */
func DownloadFile(c echo.Context) error {
	filePath := c.QueryParam("filePath")
	fmt.Println("============", filePath)
	if filePath == "" {
		//this.Ctx.WriteString("目录路径不正确")
		return result.Error(c, 402, "目录路径不正确")
	}
	filePath = utils.OsPath() + filePath
	fmt.Println(filePath)
	if !utils.IsFile(filePath) {
		return result.Error(c, 402, "目录路径不正确")
	}
	fileName := utils.GetFileName(filePath)
	//c.Response().Header().Set("Content-Type","application/x-msdownload")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+fileName)
	//第一个参数是文件的地址，第二个参数是下载显示的文件的名称
	return c.File(filePath)
}
