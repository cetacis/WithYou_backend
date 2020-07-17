package src

import (
	"context"
	"github.com/kataras/iris"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"os"
)

func PostImg(ctx iris.Context) {
	Email := ctx.FormValue("email")
	Pass := ctx.FormValue("pass")
	// get img
	file, info, err := ctx.FormFile("file")
	filename := info.Filename
	if err != nil {
		RtData := RtMsg {
			Msg: "Post Failure",
			Code: -1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}

	defer file.Close()

	filter := bson.M{"email": Email, "password": Pass}
	collection := Client.Database("WithYou").Collection("UserInfo")

	var result User
	// find info
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		RtData := RtMsg {
			Msg: "Post Failure",
			Code: -1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}

	out, err := os.OpenFile(FilePath + filename, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Warnf("Error while preparing the new file: %v", err.Error())
		return
	}
	defer out.Close()
	_, _ = io.Copy(out, file)

	result.ImgPath = URLPath + filename
	update := bson.D{{"$set", result}}
	opts := options.Update().SetUpsert(true)
	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		RtData := RtMsg {
			Msg: "Post Failure",
			Code: -1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}
	RtData := RtMsg {
		Msg: URLPath + filename,
		Code: 0,
	}
	_, _ = ctx.JSON(RtData)
}
