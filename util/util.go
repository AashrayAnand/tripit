package util

import (
	"net/http"

	"github.com/AashrayAnand/tripit/trip"
	"github.com/AashrayAnand/tripit/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CODE RELATED TO RUNNING SERVER

func RunServer() *gin.Engine {
	router := gin.Default() // initialize gin routing engine

	user.AddUserRoutes(router) // add all routes in the /user groups
	trip.AddTripRoutes(router)
	router.Use(cors.Default())
	return router
}

func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}
