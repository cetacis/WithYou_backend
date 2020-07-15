package src

import "github.com/kataras/iris"

func GetImg(ctx iris.Context) {
	filename := ctx.Params().Get("filename")
	print(filename)
	_ = ctx.SendFile(FilePath+filename, filename)
}