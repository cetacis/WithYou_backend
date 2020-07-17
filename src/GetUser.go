package src

import (
	"context"
	"github.com/kataras/iris"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(ctx iris.Context) {
	Email := ctx.FormValue("email")
	Pass := ctx.FormValue("pass")
	filter := bson.M{"email": Email}
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

	// check password
	if result.Password == Pass {
		_, _ = ctx.JSON(result)
		return
	} else {
		RtData := RtMsg {
			Msg: "Wrong Password",
			Code: 0,
		}
		_, _ = ctx.JSON(RtData)
		return
	}
}
