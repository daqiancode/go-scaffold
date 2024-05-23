package main

import (
	"go-scaffold/controllers/filters"
	"go-scaffold/controllers/publiccontrollers"
	"strings"

	"github.com/daqiancode/env"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func setupControllers(app *iris.Application) {
	// publiccontrollers.Setup(app)
	cors := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(env.Get("CORS_ORIGINS_OAUTH", "*"), ","),
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	mvcApp := mvc.New(app.Layout("layouts/main"))
	mvcApp.Router.UseRouter(cors)
	mvcApp.Handle(new(publiccontrollers.PublicController))

	apiV1 := mvcApp.Party("/v1")
	apiV1.Party("/")
}

func newApp() *iris.Application {
	app := iris.New()
	app.UseRouter(filters.NewAccessLog().Handler)
	tpl := iris.Django("views", ".html")
	tpl.Reload(env.GetBoolMust("IS_DEV", false))
	app.RegisterView(tpl)
	app.HandleDir("/static", "static")
	app.UseGlobal(filters.RecoverFilter)

	if env.GetBoolMust("OPEN_API", true) {
		publiccontrollers.SetOpenAPI(app)
	}

	return app
}

func main() {
	app := newApp()
	setupControllers(app)
	app.Run(iris.Addr(":8080"))
}
