package util

import (
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

// CODE RELATED TO RANDOM STRING GENERATION (FOR SESSION TOKEN & TESTING)
