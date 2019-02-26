package main

import (
	"strconv"
	"github.com/2brokeboys/sheepy-server/common"
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

	// Redirect all angular routes to index page
	r.GET("/create-game", routes.RedirectRoot)
	r.GET("/submit-game", routes.RedirectRoot)

	// Login route
	r.POST("/login", routes.Login)

	// Routes requiring login
	g := r.Group("/", middleware.GetUser)
	/**/ g.POST("/newGame", routes.NewGame)
	/**/ g.POST("/queryUser", routes.QueryUser)
	/**/ g.POST("/queryRecentGames", routes.QueryRecentGames)

	return r
}

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	if gin.Mode() == gin.DebugMode {
		log.Println("Creating 89 test users...")
		for i := 10; i < 100; i++ {
			db.InsertUser(&common.User{
				Name: "Hans Jürgen " + strconv.Itoa(i),
				Username: "x" + strconv.Itoa(i),
			}, "123456")
		}
		log.Println("Test users created.")
	}

	r := setupRouter()
	r.Run()
}
