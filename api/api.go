package api

import (
	"Nertivia-server-Go/api/user"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	api := route.Group("/api")
	user.Routes(api)
	//TODO: add message.Routes

}
