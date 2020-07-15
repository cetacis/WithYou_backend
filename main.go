package main

import (
	"./src"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	src.ConnectDB()
	defer src.DisconnectDb()
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	api := app.Party("/api", crs).AllowMethods(iris.MethodOptions)

	api.Post("/login", src.Login)

	api.Post("/register", src.Register)

	api.Post("/UserInfo", src.GetUser)

	api.Post("/ChangeProfile", src.ChangeProfile)

	api.Post("/GetFriend", src.GetFriendsInfo)

	api.Get("/img/{filename}", src.GetImg)
	_ = app.Run(iris.Addr(":8080"))

}
