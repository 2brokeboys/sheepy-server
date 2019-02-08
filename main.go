package main

import (
	"github.com/gin-gonic/gin"
	"github.com/2brokeboys/sheepy-server/routes"
)

func main() {
	r := gin.Default()
	r.GET("/", routes.Root)
	r.Run()
}
