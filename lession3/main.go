package main

import (
	env "lession3/config"
	"lession3/middlewares"
	routers "lession3/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(middlewares.LoggingMiddleware())

	// Load environment variables
	env.ConfigEnv()

	// Load all routes
	routers.Router(r)

	// Start the server
	if err := r.Run(":" + env.Port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
