package filters

import (
	"go-scaffold/controllers/filters/irisx"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
)

func NewAccessLog() *accesslog.AccessLog {
	ac := accesslog.New(os.Stdout)
	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = true
	ac.IP = false
	ac.BytesReceivedBody = false
	ac.BytesSentBody = false
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = false
	ac.RequestBody = false
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler
	ac.AddFields(func(ctx iris.Context, fields *accesslog.Fields) {
		fields.Set("IP", irisx.GetRealIP(ctx))
	})
	return ac
}
