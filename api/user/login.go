package user

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	// find in database
	c.JSON(200, gin.H{
		"message": "Login",
	})
}
