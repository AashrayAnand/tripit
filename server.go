package main

import (
	"fmt"
	"time"

	"github.com/AashrayAnand/tripit/session"
	"github.com/AashrayAnand/tripit/trip"
	"github.com/AashrayAnand/tripit/user"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := session.Client.SetNX("key", "value", 10*time.Second).Result()
	val, err := session.Client.Get("key").Result()

	fmt.Println(val, err)
	router := gin.Default() // initialize gin routing engine

	user.AddUserRoutes(router) // add all routes in the /user groups
	router.POST("/trip", trip.Create)

	//router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run() // listen and server 0.0.0.0:8080
}
