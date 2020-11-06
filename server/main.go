package main

import (
	"brown/route"

	"github.com/kataras/iris"
	"github.com/rs/cors"
)

func main() {
	app := newApp()
	route.InitRouter(app)
	app.Run(iris.Addr(":8082"))
}

func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	app.Logger().SetLevel("debug")
	app.StaticWeb("/", "./static")
	app.RegisterView(iris.HTML("./static", ".html").Reload(true))
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}
	corsWrapper := cors.New(corsOptions).ServeHTTP
	app.WrapRouter(corsWrapper)
	return app
}
