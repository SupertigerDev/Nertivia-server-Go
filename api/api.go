package api

import (
	"github.com/gin-gonic/gin"
	"github.com/supertiger1234/Nertivia-server-Go/api/user"
)

func Routes(route *gin.Engine) {
	api := route.Group("/api")
	user.Routes(api)
	//TODO: add message.Routes

}
