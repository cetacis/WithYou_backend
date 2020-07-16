package src

import (
	"context"
	"github.com/kataras/iris"
)

func Match(ctx iris.Context) {
	TaskId := ctx.FormValue("TaskId")
	Email := ctx.FormValue("Email")
	QueueData := TaskQueue {
		TaskId: TaskId,
		Email: Email,
	}

	collection := Client.Database("WithYou").Collection("MatchQueue")
	_, err := collection.InsertOne(context.TODO(), QueueData)
	if err != nil {
		RtData := RtMsg {
			Msg: "Server Error!",
			Code: 1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}

	RtData := RtMsg {
		Msg: "add queue success",
		Code: 0,
	}
	_, _ = ctx.JSON(RtData)
	return
}
