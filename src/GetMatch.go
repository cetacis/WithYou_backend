package src

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMatch(ctx iris.Context) {
	TaskId := ctx.FormValue("TaskId")
	Email := ctx.FormValue("Email")

	// is matched by other
	collection := Client.Database("WithYou").Collection("MatchedMap")
	filter := bson.M{"first": Email}
	var info MatchInfo
	err := collection.FindOne(context.TODO(), filter).Decode(&info)
	if err == nil {
		_, err = collection.DeleteOne(context.TODO(), filter)
		// find then delete
		if err != nil {
			RtData := RtMsg {
				Msg: "Server Error!",
				Code: 2,
			}
			_, _ = ctx.JSON(RtData)
			return
		}
		RtData := RtMsg {
			Msg: info.Second,
			Code: 0,
		}
		_, _ = ctx.JSON(RtData)
		return
	}

	var foo TaskQueue
	// haven't matched by other
	collection = Client.Database("WithYou").Collection("MatchQueue")
	// IsAddQueue ?
	err = collection.FindOne(context.TODO(), TaskQueue {
		Email: Email,
		TaskId: TaskId,
	}).Decode(&foo)

	if err != nil {
		RtData := RtMsg {
			Msg: "No queue info",
			Code: 1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}

	filter = bson.M{"taskid": TaskId}
	opts := options.Find()
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		RtData := RtMsg {
			Msg: "No queue info",
			Code: 1,
		}
		_, _ = ctx.JSON(RtData)
		return
	}

	var results []TaskQueue
	if err = cursor.All(context.TODO(), &results); err != nil {
		RtData := RtMsg {
			Msg: "Server Error!",
			Code: 2,
		}
		_, _ = ctx.JSON(RtData)
		return
	}

	fmt.Println(results)
	for _, result := range results {
		if result.Email != Email {
			// matched
			collection = Client.Database("WithYou").Collection("MatchedMap")
			info := MatchInfo {
				First: result.Email,
				Second: Email,
			}
			_, err = collection.InsertOne(context.TODO(), info)
			// insert matched
			if err != nil {
				RtData := RtMsg {
					Msg: "Server Error!",
					Code: 2,
				}
				_, _ = ctx.JSON(RtData)
				return
			}
			// add success and delete queue
			collection = Client.Database("WithYou").Collection("MatchQueue")
			_, err = collection.DeleteOne(context.TODO(), TaskQueue {
				TaskId: TaskId,
				Email: Email,
			})
			if err != nil {
				RtData := RtMsg {
					Msg: "Server Error!",
					Code: 2,
				}
				_, _ = ctx.JSON(RtData)
				return
			}

			_, err = collection.DeleteOne(context.TODO(), TaskQueue {
				TaskId: TaskId,
				Email: result.Email,
			})
			if err != nil {
				RtData := RtMsg {
					Msg: "Server Error!",
					Code: 2,
				}
				_, _ = ctx.JSON(RtData)
				return
			}

			// add partner info
			collection = Client.Database("WithYou").Collection("UserInfo")
			opts := options.Update().SetUpsert(true)
			filter := bson.M{"email": Email}
			update := bson.D{{"$set", bson.D{{"partner", result.Email}}}}
			_, _ = collection.UpdateOne(context.TODO(), filter, update, opts)
			filter = bson.M{"email": result.Email}
			update = bson.D{{"$set", bson.D{{"partner", Email}}}}
			_, _ = collection.UpdateOne(context.TODO(), filter, update, opts)
			RtData := RtMsg {
				Msg: result.Email,
				Code: 0,
			}
			_, _ = ctx.JSON(RtData)
			return
		}
	}

	RtData := RtMsg {
		Msg: "Haven't find out. Plz wait",
		Code: -1,
	}
	_, _ = ctx.JSON(RtData)
	return
}