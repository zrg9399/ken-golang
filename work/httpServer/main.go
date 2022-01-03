package main

import (
	"httpServer/context"
	"httpServer/router"
)

func main() {
	r := context.NewRouter()
	router.Add(r)
	r.Run()
}
