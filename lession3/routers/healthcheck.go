package router

import "github.com/gin-gonic/gin"

func HealthCheckRouter(r *gin.Engine) {
	// router group : healthcheck
	healthcheck := r.Group("/healthcheck")
	{
		healthcheck.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code":    200,
				"message": "ok",
			})
		})

		healthcheck.GET("/db", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code":    200,
				"message": "db is running",
			})
		})
	}
}
