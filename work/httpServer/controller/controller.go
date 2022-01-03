package controller

import "httpServer/context"

func Healthz(ctx *context.Context) {
	ctx.Error("Hi healthz")
}
func Hello(ctx *context.Context) {
	ctx.Normal(map[string]interface{}{
		"data": "你好",
		"msg":  "hello world",
	})
}
