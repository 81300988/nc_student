package middlewareCustom

import (
	"io"
	"log"
	"os"
)

var Logger Customlogger

type Customlogger struct {
}

func (p Customlogger) Write(data []byte) (n int, err error) {
	f, err := os.OpenFile("D:/GoNordic/go2/nc_student/logger/Http.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println(string(data))
	return len(data), nil
}

// func SimpleLogger() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			req := c.Request()
// 			err := next(c)
// 			res := c.Response()
// 			log.Println(req.URL.Host, req.URL.Path, req.Method, res.Status)
// 			return err
// 		}

// 	}
// }
