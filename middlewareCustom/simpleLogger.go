package middlewareCustom

import (
	"github.com/labstack/echo/v4"
	db "nc_student.com/v1/db"
	"nc_student.com/v1/model"
	"time"
)

// var Logger Customlogger

// type Customlogger struct {
// }

// func (p Customlogger) Write(data []byte) (n int, err error) {
// 	f, err := os.OpenFile("Http.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatalf("error opening file: %v", err)
// 	}
// 	defer f.Close()
// 	wrt := io.MultiWriter(os.Stdout, f)
// 	log.SetOutput(wrt)
// 	log.Println(string(data))
// 	return len(data), nil
// }

func SimpleLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			err := next(c)
			res := c.Response()
			var log model.Log
			log.Date = time.Now()
			log.Host = req.Host
			log.Path = req.URL.Path
			log.Method = req.Method
			log.Status = res.Status
			go db.AddLog(&log)
			return err
		}

	}
}
