package main

import (
	"github.com/2brokeboys/sheepy-server/db"
	"github.com/2brokeboys/sheepy-server/middleware"
	"github.com/2brokeboys/sheepy-server/routes"

	"log"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Setup sessions
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("_", store))

	// Setup Gzip compression
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Serve static files
	r.Static("/assets", "../sheepy-client/dist/webpack/website/assets")
	r.StaticFile("/main.js", "../sheepy-client/dist/webpack/website/main.js")

	// HTTP root serves html page
	r.GET("/", routes.Root)
	r.HEAD("/", routes.Root)

	// Login route
	r.POST("/login", routes.Login)

	// Routes requiring login
	g := r.Group("/", middleware.GetUser)
	/**/ g.POST("/newGame", routes.NewGame)
	/**/ g.POST("/queryUser", routes.QueryUser)

	return r
}

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := setupRouter()
	r.Run()
}
