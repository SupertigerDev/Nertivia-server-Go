package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	// find in database
	var loginCmd loginCommand
	c.BindJSON(&loginCmd)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login",
		"test":    loginCmd.Email,
	})
}
