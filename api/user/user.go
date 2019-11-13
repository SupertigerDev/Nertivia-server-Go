package user

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {

	user := route.Group("/users")
	{
		// auth
		user.POST("/login", Login)
		user.POST("/register", Register)
	}
}
