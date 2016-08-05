package middleware

import (
	"net/http"
	"strings"

	"gopkg.in/labstack/echo.v1"
)

func CORS(origins ...string) func(*echo.Context) error {
	return func(ctx *echo.Context) error {
		origin := ctx.Request().Header.Get("Origin")

		if origin == "" && ctx.Request().Method != echo.OPTIONS {
			return nil
		}

		found := false
		for _, o := range origins {
			if strings.Index(origin, o) != -1 {
				found = true
				break
			}
		}

		if !found {
			return nil
		}

		rsp := ctx.Response()
		rsp.Header().Set("Access-Control-Allow-Origin", origin)
		rsp.Header().Set("Access-Control-Allow-Credentials", "true")
		rsp.Header().Set("Access-Control-Allow-Headers", "Content-Type,Ajax")

		if ctx.Request().Method == echo.OPTIONS {
			return echo.NewHTTPError(http.StatusOK, "ok")
		}

		return nil
	}
}
