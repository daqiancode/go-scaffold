package usercontrollers

import "github.com/kataras/iris/v12"

type UserControllers struct{}

type UserPostForm struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"gte=0,lte=130"`
	Password string `json:"password" validate:"required,gte=6,lte=20"`
}

func (u *UserControllers) Post(ctx iris.Context) error {
	var form UserPostForm
	if err := ctx.ReadJSON(&form); err != nil {
		return err
	}

	return ctx.JSON(form)
}
