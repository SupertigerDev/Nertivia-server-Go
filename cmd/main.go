package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertiger1234/socketio"

	"github.com/supertiger1234/Nertivia-server-Go/api/user"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Socket events
	server.On("connection", func(so socketio.Socket) {
		fmt.Println("connected")
	})
	server.On("authentication", func(so socketio.Socket) {
		fmt.Println("auth")
		so.Emit("success", gin.H{
			"user": gin.H{
				"friends": []gin.H{},
			},
			"settings": gin.H{},
		})
	})

	// Cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowCredentials: true,
	}))

	// Routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
		})
	})

	// Socket io
	user.Routes(router)
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))

	fmt.Println("Running at http://localhost:80")
	router.Run(":80")

}
