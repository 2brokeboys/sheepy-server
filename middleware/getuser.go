package middleware

import (
	"github.com/2brokeboys/sheepy-server/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	session := sessions.Default(c)
	user, ok := session.Get("user").(*common.User)
	if !ok {
		c.AbortWithStatusJSON(403, gin.H{
			"error": "not logged in",
		})
		return
	}
	c.Set("user", user)
}
