package src

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"os"
)

const FilePath = "./img/"
const URLPath = "/img/"

func Register(ctx iris.Context) {
	// get info
	Name := ctx.FormValue("name")
	Email := ctx.FormValue("email")
	Pass := ctx.FormValue("pass")
	// get img
	file, info, err := ctx.FormFile("file")
	fmt.Println(Name, " 1", Email, "2", Pass)
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
	ImgPath := URLPath + filename

	// create User info
	UserData := User {
		Username: Name,
		Email: Email,
		Password: Pass,
		ImgPath: ImgPath,
		Mobile: "",
		Age: "",
		Bio: "",
		Constellation: "",
		Birthday: "",
		Sex: "",
		TogetherTasks: make([]TogetherTask, 0),
		PrivateTasks: make([]PrivateTask, 0),
		Friends: make([]string, 0),
		Messages: make([]Message, 0),
		CurrentTaskId: -1,
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

	RtData := RtMsg {
		Msg: "Register succeed",
		Code: 0,
	}
	_, _ = ctx.JSON(RtData)
}
