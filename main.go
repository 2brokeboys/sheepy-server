package main

import (
	"github.com/2brokeboys/sheepy-server/routes"
	"github.com/2brokeboys/sheepy-server/db"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("_", store))

	r.GET("/", routes.Root)
	r.GET("/newGame", routes.NewGame)
	r.GET("/queryUser", routes.QueryUser)
	r.GET("/login", routes.Login)
	r.Run()
}
