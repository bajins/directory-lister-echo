package main

import (
	"directory-lister-echo/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"path"
	"path/filepath"
)

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "index")
}

func Admin(c echo.Context) error {
	fmt.Println(c.Param("path"))
	return c.JSON(http.StatusOK, "ok")
}

// 获取展示目录下的所有
func GetDir(c echo.Context) error {

	dir := c.QueryParam("dir")
	//root, _ := os.Getwd()
	root := "D:\\cyw\\upload.1"
	if c.Request().URL.Path == "/" || dir != "" {
		if dir == "" {
			dir = "/"
		}

		fmt.Println(c.Path(), c.Request().URL.Path, dir)
		return GetDirList(c, root, dir)
	}

	return DownloadFile(c, root)

}

func GetDirList(c echo.Context, root, dir string) error {
	p := path.Join(root, dir)
	if utils.IsExistDir(p) {
		return Error(c, 300, "不是目录")
	}
	// 获取目录下的文件和子目录信息
	list := utils.GetFileList(p)
	if list == nil {
		return nil
	}
	// 创建切片
	var dirs []map[string]interface{}
	for _, info := range list {
		// 创建map
		m := make(map[string]interface{})

		m["name"] = info.Name()
		link := path.Join(dir, info.Name())
		if !info.IsDir() {
			// Size()单位为Byte,所以要转换
			m["size"] = utils.ByteSize(uint64(info.Size()))
		} else {
			m["size"] = "-"
			//link = "/?dir=" + path.Join(dir, info.Name())
		}
		// 时间
		m["modTime"] = utils.TimeToString(info.ModTime())
		// 权限
		//m["mode"] = info.Mode().String()
		m["link"] = utils.PathSeparatorSlash(link)
		m["isDir"] = info.IsDir()
		// 放进切片中
		dirs = append(dirs, m)
	}

	data := make(map[string]interface{})
	//data["total"] = total
	data["file"] = dirs
	links := utils.PathSplitter(dir, "Bajins Soft")
	data["links"] = links
	return Success(c, "获取文件列表成功", data)
}

// 下载文件
func DownloadFile(c echo.Context, root string) error {
	filePath := path.Join(root, c.Request().URL.Path)
	if !utils.IsFileExist(filePath) {
		return Error(c, 402, "不是文件")
	}
	filename := filepath.Base(filePath)
	//ft, err := utils.GetContentType(filepath.Ext(filename))
	//if err != nil {
	//	return Error(c, http.StatusInternalServerError, err.Error())
	//}
	//c.Response().Header().Set(echo.HeaderContentType, ft)
	//fi,err:=os.Stat(filename)
	//if err != nil {
	//	return Error(c, http.StatusInternalServerError, err.Error())
	//}
	//c.Response().Header().Set(echo.HeaderContentLength, utils.ToString(fi.Size()))
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+filename)
	//第一个参数是文件的地址，第二个参数是下载显示的文件的名称
	return c.File(filePath)
}
