package main

import (
	"context"
	"github.com/kataras/iris"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"os"
)

const FilePath = "./img"

func Register(ctx iris.Context) {
	// get info
	Name := ctx.FormValue("name")
	Email := ctx.FormValue("email")
	Pass := ctx.FormValue("pass")
	// get img
	file, info, err := ctx.FormFile("img")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Warnf("Error while uploading: %v", err.Error())
		return
	}
	defer file.Close()
	filename := info.Filename
	out, err := os.OpenFile(FilePath + filename, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Warnf("Error while preparing the new file: %v", err.Error())
		return
	}
	defer out.Close()
	_, _ = io.Copy(out, file)
	ImgPath := FilePath + filename

	// create User info
	UserData := User {
		Username: Name,
		Email: Email,
		Password: Pass,
		ImgPath: ImgPath,
	}
	var result User
	collection := Client.Database("WithYou").Collection("UserInfo")

	// find email
	filter := bson.M{"email": Email}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		RtData := RtMsg {
			Msg: "Email has existed",
			Code: -1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}

<<<<<<< HEAD
	// insert register data
	_, err = collection.InsertOne(context.TODO(), UserData)
	if err != nil {
		RtData := RtMsg {
			Msg: "Register error",
			Code: 1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}
=======
	api.Post("/register", src.Register)

	api.Post("/UserInfo", src.GetUser)

	api.Post("/ChangeProfile", src.ChangeProfile)

	api.Post("/GetFriend", src.GetFriendsInfo)

	api.Get("/img/{filename}", src.GetImg)
	_ = app.Run(iris.Addr(":8080"))
>>>>>>> 4af942fec19a1940e1240c4c9aa692906388cfa8

	RtData := RtMsg {
		Msg: "Register succeed",
		Code: 0,
	}
	_, _ = ctx.JSON(RtData)
}
