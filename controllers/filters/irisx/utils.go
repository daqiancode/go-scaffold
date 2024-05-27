package irisx

import (
	"strings"

	"github.com/go-playground/validator/v10"
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

type validationError struct {
	Field  string `json:"field"`
	Should string `json:"should"`
	Param  string `json:"param,omitempty"`
}

var DecaptalizeFieldNames = true

func decaptalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
func WrapValidationErrors(errs validator.ValidationErrors) []validationError {

	validationErrors := make([]validationError, 0, len(errs))
	for _, err := range errs {
		field := err.Field()
		if DecaptalizeFieldNames {
			field = decaptalize(field)
		}

		validationErrors = append(validationErrors, validationError{
			Field:  field,
			Should: err.ActualTag(),
			Param:  err.Param(),
		})
	}

	return validationErrors
}
