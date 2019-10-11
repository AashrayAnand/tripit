package main

import (
	"github.com/AashrayAnand/Bill-List/user"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // initialize gin routing engine

	user.AddUserRoutes(router) // add all routes in the /user group
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run() // listen and server 0.0.0.0:8080
}
