package middlewares

import (
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

func zeroLogMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()
	p := req.URL.Path

	bytesIn := req.Header.Get(echo.HeaderContentLength)
	log.Info().
		Str("remote_ip", c.RealIP()).
		Str("host", req.Host).
		Str("uri", req.RequestURI).
		Str("method", req.Method).
		Str("path", p).
		Str("referer", req.Referer()).
		Str("user_agent", req.UserAgent()).
		Int("status", res.Status).
		Str("bytes_in", bytesIn).
		Str("bytes_out", strconv.FormatInt(res.Size, 10)).
		Str("latency", strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10)).
		Str("latency_human", stop.Sub(start).String()).
		Str("time_rfc3339", stop.Sub(start).String()).
		Msg("Handled request")

	return nil
}

func logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return zeroLogMiddlewareHandler(c, next)
	}
}

func NewLogger() echo.MiddlewareFunc {
	return logger
}
