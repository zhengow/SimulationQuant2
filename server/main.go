package main

import (
    "github.com/kataras/iris"
    "github.com/rs/cors"
    "brown/route"
    // "fmt"
    // "brown/model"
)


func main() {
    app := newApp()
    route.InitRouter(app)
    
    // orm, err := model.CreateEngine()
    // if err != nil {
    //     fmt.Println(err)
    // }
    // app.Use(crs)

    // v1 := app.Party("/api")
    // {
    //     v1.Post("/register", func(ctx iris.Context) {
    //     c := &model.User{}
    //     if err := ctx.ReadJSON(c); err != nil{
    //         panic(err.Error())
    //     }else{
    //         _, err := orm.Insert(c)
    //         if err != nil {
    //             fmt.Println(err)
    //         }
    //         ctx.JSON(c)
    //     }
    //     // ctx.WriteString("pong")
    //     })
    // }

    app.Run(iris.Addr(":8082"))
}


func newApp() *iris.Application {
	app := iris.New()
    app.Configure(iris.WithOptimizations)
    app.Logger().SetLevel("debug")
	corsOptions := cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowCredentials: true,
    }
    corsWrapper := cors.New(corsOptions).ServeHTTP
    app.WrapRouter(corsWrapper)
	return app
}