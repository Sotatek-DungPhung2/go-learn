package router

import (
	userContrller "lession3/modules/users/controllers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	// router group : healthcheck
	HealthCheckRouter(r)

	// router group : api
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/users", userContrller.GetUsers)
		apiV1.POST("/users", userContrller.CreateUser)
	}

}
