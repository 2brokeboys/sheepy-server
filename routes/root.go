package routes

import (
	"github.com/gin-gonic/gin"
)

// Root handles the / route by serving index.html
func Root(c *gin.Context) {
	if pusher := c.Writer.Pusher(); pusher != nil {
		pusher.Push("/main.js", nil)
	}
	c.File("../sheepy-client/dist/webpack/website/index.html")
}
