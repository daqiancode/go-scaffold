package irisx

import (
	"strings"

	"github.com/kataras/iris/v12"
)

const PageSize = 20

type PageData struct {
	Items     any   `json:"items"`
	Total     int64 `json:"total"`
	PageIndex int   `json:"page"`
	PageSize  int   `json:"pageSize"`
	PageCount int   `json:"pageCount"`
}

func GetRealIP(ctx iris.Context) string {
	ipStr := ctx.GetHeader("X-Forwarded-For")
	if ipStr != "" {
		ips := strings.Split(ipStr, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}
	ip := ctx.GetHeader("X-Real-IP")
	if ip != "" {
		return ip
	}
	return ctx.RemoteAddr()
}

// func View(app iris.Party, path, view string, data iris.Map) {
// 	app.Get(path, func(ctx iris.Context) {
// 		ctx.View(view, data)
// 	})
// }
