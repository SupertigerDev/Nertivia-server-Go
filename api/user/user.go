package user

import (
	"github.com/gin-gonic/gin"
	"github.com/supertiger1234/Nertivia-server-Go/api/user/auth"
)

func Routes(route *gin.Engine) {

	user := route.Group("/user")
	{
		user.GET("/login", auth.Login)
		user.GET("/register", auth.Register)
	}
}
