package main

import (
	"Nertivia-server-Go/api"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertiger1234/socketio"
)

// TODO: You should really check for errors
func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Socket events
	_ = server.On("connection", func(so socketio.Socket) {
		fmt.Println("connected")
	})
	_ = server.On("authentication", func(so socketio.Socket) {
		fmt.Println("auth")
		_ = so.Emit("success", gin.H{
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
	api.Routes(router)
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))

	fmt.Println("Running at http://localhost:80")
	_ = router.Run(":80")
}
