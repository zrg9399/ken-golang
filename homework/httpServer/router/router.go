package router

import (
	"httpServer/context"
	"httpServer/controller"
)

func Add(r *context.Router) {
	//r.GET("/getTest", controller.Healthz, controller.getTest)
	r.GET("/healthz", controller.Healthz, controller.Hello)
	r.POST("/hello", controller.Hello)
	r.PUT("/hello", controller.Healthz)
	r.DELETE("/hello", controller.Healthz)
}
