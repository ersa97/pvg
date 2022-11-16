package middlewares

import "github.com/labstack/echo/v4"

func RouteHandler(err error, c echo.Context) {
	he, _ := err.(*echo.HTTPError)
	if he.Code != 200 {
		c.JSON(he.Code, map[string]interface{}{
			"status":  "failed",
			"code":    he.Code,
			"message": he.Message,
		})
		return
	}
}
