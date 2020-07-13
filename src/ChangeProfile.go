package src

import (
	"context"
	"github.com/kataras/iris"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ChangeProfile(ctx iris.Context) {
	var UserInfo User
	err := ctx.ReadJSON(&UserInfo)
	if err != nil {
		RtData := RtMsg {
			Msg: "Server Read Json Error!",
			Code: -1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}
	email := UserInfo.Email
	pass := UserInfo.Password
	collection := Client.Database("WithYou").Collection("UserInfo")
	filter := bson.M{"email": email, "password": pass}
	update := bson.D{{"$set", UserInfo}}
	opts := options.Update().SetUpsert(true)
	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		RtData := RtMsg {
			Msg: "Update Failure",
			Code: -1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}
	RtData := RtMsg {
		Msg: "Update Success",
		Code: 0,
	}
	_, _ = ctx.JSON(RtData)
	return

}
