package result

import (
	"github.com/labstack/echo"
	"net/http"
)

/**
 * 请求返回成功
 *
 * @author claer www.bajins.com
 * @date 2019/6/28 12:48
 */
func Success(c echo.Context, msg string, data interface{}) error {
	own := make(map[string]interface{})
	own["code"] = 200
	own["message"] = msg
	own["data"] = data
	return c.JSON(http.StatusOK, own)
}

/**
 * 请求返回错误
 *
 * @author claer www.bajins.com
 * @date 2019/6/28 12:48
 */
func Error(c echo.Context, code int, msg string) error {
	own := make(map[string]interface{})
	own["code"] = code
	own["message"] = msg
	return c.JSON(http.StatusOK, own)
}
