package src

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
)

func Match(ctx iris.Context) {
	TaskId := ctx.FormValue("TaskId")
	Email := ctx.FormValue("Email")
	QueueData := TaskQueue {
		TaskId: TaskId,
		Email: Email,
	}

	var foo TaskQueue
	collection := Client.Database("WithYou").Collection("MatchQueue")
	err := collection.FindOne(context.TODO(), QueueData).Decode(&foo)
	fmt.Println(foo, err)
	if err == nil {
		RtData := RtMsg {
			Msg: "You are queuing!",
			Code: -2,
		}
		_, _ = ctx.JSON(RtData)
		return
	}

	_, err = collection.InsertOne(context.TODO(), QueueData)
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
