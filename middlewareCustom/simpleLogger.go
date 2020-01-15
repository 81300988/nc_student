package middlewareCustom

import (
	"log"

	"github.com/labstack/echo/v4"
)

func SimpleLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			err := next(c)
			res := c.Response()
			log.Println(req.URL.Host, req.URL.Path, req.Method, res.Status)
			return err
		}

	}
}
