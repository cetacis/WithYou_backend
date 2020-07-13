package src

import (
	"context"
	"github.com/kataras/iris"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFriendsInfo(ctx iris.Context) {
	email := ctx.FormValue("email")
	filter := bson.M{"email": email}
	var result User
	collection := Client.Database("WithYou").Collection("UserInfo")

	// find email
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		RtData := RtMsg {
			Msg: "Email does not exist",
			Code: -1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}

	// change the password
	result.Password = "*********"
	_, _ = ctx.JSON(result)

}
