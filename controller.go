package main

import (
	"directory-lister-echo/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "index")
}

func Admin(c echo.Context) error {
	log.Println(c.Param("path"))
	return c.JSON(http.StatusOK, "ok")
}

// 获取展示目录下的所有
func GetDir(c echo.Context) error {
	dir := c.QueryParam("dir")
	if c.Request().URL.Path == "/" || dir != "" {
		if dir == "" {
			dir = "/"
		}

		log.Println(c.Path(), c.Request().URL.Path, dir)
		return GetDirList(c, dir)
	}

	return DownloadFile(c)

}

// 对路径进行重组为目录名+路径
// path string 路径
// rootName string 路径头，根目录的名称，就是/的名称
func PathSplitter(toPath string, rootName string) []map[string]string {
	// 替换路径中的分割符
	toPath = utils.PathSeparatorSlash(toPath)
	// 判断第一个字符是否为分割符
	indexSplitter := strings.Index(toPath, "/")
	if indexSplitter != 0 {
		toPath = path.Join("/", toPath)
	}
	var links []map[string]string
	rootLink := make(map[string]string)
	rootLink["name"] = rootName
	rootLink["path"] = "/"
	links = append(links, rootLink)
	// 如果是根目录，那么就返回
	if utils.IsStringEmpty(toPath) || toPath == "/" {
		return links
	}
	// 避免分割路径时多分割一次，去掉第一个分割符，并对路径分割
	split := strings.Split(toPath[1:], "/")
	for _, v := range split {
		link := make(map[string]string)
		link["name"] = v
		link["path"] = path.Join(toPath[0:strings.Index(toPath, v)], v)
		links = append(links, link)
	}
	return links
}

func GetDirList(c echo.Context, dir string) error {
	//root, _ := os.Getwd()
	root := "D:\\v2rayN"
	p := filepath.Join(root, dir)
	if utils.IsExistDir(p) {
		return ErrorJSON(c, 300, "不是目录")
	}
	// 获取目录下的文件和子目录信息
	list, err := utils.GetFileList(p)
	if list == nil || err != nil {
		return ErrorJSON(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	// 创建切片
	var dirs []map[string]interface{}
	for _, info := range list {
		// 创建map
		m := make(map[string]interface{})

		m["name"] = info.Name()
		link := filepath.Join(dir, info.Name())
		if !info.IsDir() {
			// Size()单位为Byte,所以要转换
			m["size"] = utils.ByteSize(uint64(info.Size()))
		} else {
			m["size"] = "-"
			//link = "/?dir=" + link
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
	links := PathSplitter(dir, "Bajins Soft")
	data["links"] = links
	return SuccessJSON(c, "获取文件列表成功", data)
}

// 下载文件
func DownloadFile(c echo.Context) error {
	filePath := strings.Replace(c.Request().URL.Path, "/download", "", 1)
	log.Println(filePath, c.QueryParams())
	if !utils.IsFileExist(filePath) {
		//root, _ := os.Getwd()
		root := "D:\\v2rayN"
		filePath = filepath.Join(root, filePath)
	}
	if !utils.IsFileExist(filePath) {
		return SystemErrorJSON(c, http.StatusNotFound, "文件不存在")
	}
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return SystemErrorJSON(c, http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name()))

	//fi,err:=os.Stat(filename)
	//if err != nil {
	//	return ErrorJSON(c, http.StatusInternalServerError, err.ErrorJSON())
	//}
	//c.Response().Header().Set(echo.HeaderContentLength, utils.ToString(fi.Size()))

	// 获取后缀并判断
	if filepath.Ext(filePath) == "" {
		//ft, err := utils.GetContentType(filePath)
		//if err != nil {
		//	return SystemErrorJSON(c, http.StatusInternalServerError, err.ErrorJSON())
		//}
		c.Response().Header().Set(echo.HeaderContentType, "")

		//f, err := os.Open(filePath)
		//if err != nil {
		//	return echo.NotFoundHandler(c)
		//}
		//return c.Stream(http.StatusOK, "", f)
	}
	return c.File(filePath)
}
