package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// 请求返回成功
func SuccessJSON(c echo.Context, msg string, data interface{}) error {
	own := make(map[string]interface{})
	own["code"] = 200
	own["message"] = msg
	own["data"] = data
	return c.JSON(http.StatusOK, own)
}

// 请求返回错误
func ErrorJSON(c echo.Context, code int, msg string) error {
	own := make(map[string]interface{})
	own["code"] = code
	own["message"] = msg
	return c.JSON(http.StatusOK, own)
}

// 系统错误
func SystemErrorJSON(c echo.Context, code int, msg string) error {
	own := make(map[string]interface{})
	own["code"] = code
	own["message"] = msg
	return c.JSON(code, own)
}
