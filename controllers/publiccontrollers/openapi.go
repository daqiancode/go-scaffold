package publiccontrollers

import (
	"go-scaffold/logs"
	"os"

	"github.com/daqiancode/env"
	"github.com/daqiancode/myutils/strs"
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v3"
)

func SetOpenAPI(route iris.Party) {
	route.Get("/openapi/json", func(ctx iris.Context) {
		openAPIBs, err := os.ReadFile("views/openapi/openapi.yaml")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"error": err.Error()})
			return
		}
		openYaml := strs.FormatTpl(string(openAPIBs), map[string]any{
			"IAM_HOST":    env.Get("IAM_HOST"),
			"PATH_PREFIX": env.Get("PATH_PREFIX"),
		})
		var openAPI any
		err = yaml.Unmarshal([]byte(openYaml), &openAPI)
		if err != nil {
			logs.Log.Error().Err(err).Msg("Unmarshal openapi.yaml")
			ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"error": err.Error()})
		}
		ctx.JSON(openAPI)
	})
	route.Get("/openapi/oauth2-redirect", func(ctx iris.Context) {
		ctx.View("openapi/oauth2-redirect")
	})
	route.Get("/openapi", func(ctx iris.Context) {
		data := iris.Map{
			"IAM_HOST":     env.Get("IAM_HOST"),
			"clientId":     "client_id",
			"clientSecret": "client_secret",
		}
		ctx.View("openapi/openapi", data)
	})
}
