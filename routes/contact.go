package routes

import "github.com/gin-gonic/gin"

func Contact(c *gin.Context) {
	c.HTML(200, "contact.tmpl", nil)
}
