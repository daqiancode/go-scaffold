package publiccontrollers

import "github.com/kataras/iris/v12"

type PublicController struct{}

func (c *PublicController) Get(ctx iris.Context) error {
	return ctx.View("index.html", iris.Map{"title": "Index"})
}

func (c *PublicController) GetHealth(ctx iris.Context) string {
	return "OK"
}
