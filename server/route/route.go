package route

import (
	"brown/controller"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func InitRouter(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	bathUrl := "/api"
	mvc.New(app.Party(bathUrl + "/register")).Handle(controller.NewRegisterController())
	mvc.New(app.Party(bathUrl + "/login")).Handle(controller.NewLoginController())
	mvc.New(app.Party(bathUrl + "/predict")).Handle(controller.NewPredictController())
	// app.Use(middleware.GetJWT().Serve)  // jwt
	// mvc.New(app.Party(bathUrl +"/book")).Handle(controllers.NewBookController())
}
