package filters

import (
	"fmt"
	"runtime/debug"

	"github.com/kataras/iris/v12"
)

func RecoverFilter(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			fmt.Println(err)
			fmt.Println(stack)
			if e, ok := err.(error); ok {
				if ctx.IsStopped() { // handled by other middleware.
					return
				} else {
					ctx.StopWithError(iris.StatusInternalServerError, e)
					return
				}
			}
		}
	}()

	ctx.Next()
}
