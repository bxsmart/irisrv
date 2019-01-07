package utils

import (
	"github.com/kataras/iris"
	"strings"
)

func GetRemoteAddr(ctx iris.Context) string {
	ip := ctx.GetHeader("X-Forwarded-For")
	if len(ip) <= 0 || strings.ToUpper(ip) == "UNKNOWN" {
		ip = ctx.GetHeader("Proxy-Client-IP")
	}

	if len(ip) <= 0 || strings.ToUpper(ip) == "UNKNOWN" {
		ip = ctx.GetHeader("WL-Proxy-Client-IP")
	}

	if len(ip) <= 0 || strings.ToUpper(ip) == "UNKNOWN" {
		ip = ctx.GetHeader("HTTP_CLIENT_IP")
	}

	if len(ip) <= 0 || strings.ToUpper(ip) == "UNKNOWN" {
		ip = ctx.GetHeader("HTTP_X_FORWARDED_FOR")
	}

	ip = ctx.RemoteAddr()

	return ip
}
