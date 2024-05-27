package main

import (
	"go-scaffold/controllers/filters"
	"go-scaffold/controllers/filters/irisx"
	"go-scaffold/controllers/publiccontrollers"
	"go-scaffold/controllers/usercontrollers"
	"strings"

	"github.com/daqiancode/env"
	"github.com/daqiancode/jwts"
	"github.com/go-playground/validator/v10"
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
	mvcApp.HandleError(func(ctx iris.Context, err error) {
		switch e := err.(type) {
		case validator.ValidationErrors:
			ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{"error": "Validation error", "validationErrors": irisx.WrapValidationErrors(e)})
		case error:
			ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"error": e.Error()})
		}
	})
	mvcApp.Router.UseRouter(cors)
	mvcApp.Handle(new(publiccontrollers.PublicController))

	apiV1 := mvcApp.Party("/v1")
	apiV1.Party("/user").Handle(new(usercontrollers.UserControllers))
}

func newApp() *iris.Application {
	app := iris.New()
	app.Validator = validator.New()
	app.UseRouter(filters.NewAccessLog().Handler)
	tpl := iris.Django("views", ".html")
	tpl.Reload(env.GetBoolMust("IS_DEV", false))
	app.RegisterView(tpl)
	app.HandleDir("/static", "static")
	app.UseGlobal(filters.RecoverFilter)
	tokenSetter := jwts.AccessTokenSetter(jwts.AccessTokenSetterConfig{PublicKey: env.Get("JWT_PUBLIC_KEY")})
	app.Use(tokenSetter)

	if env.GetBoolMust("OPEN_API", true) {
		publiccontrollers.SetOpenAPI(app)
	}

	return app
}

func main() {
	app := newApp()
	setupControllers(app)
	app.Run(iris.Addr(":8080"), iris.WithOptimizations, iris.WithoutServerError(iris.ErrServerClosed), iris.WithRemoteAddrHeader("X-Real-Ip"))
}
