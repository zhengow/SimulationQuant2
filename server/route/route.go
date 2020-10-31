package route

import(
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"brown/controller"
)

func InitRouter(app *iris.Application) {
	bathUrl := "/api"
	mvc.New(app.Party(bathUrl + "/register")).Handle(controller.NewRegisterController())
	mvc.New(app.Party(bathUrl + "/login")).Handle(controller.NewLoginController())
	// app.Use(middleware.GetJWT().Serve)  // jwt
	// mvc.New(app.Party(bathUrl +"/book")).Handle(controllers.NewBookController())
}