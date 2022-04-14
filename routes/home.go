package routes

import "github.com/gin-gonic/gin"

func Home(c *gin.Context) {
	c.HTML(200, "home.tmpl", gin.H{
		"title": "Home Page",
	})
}
