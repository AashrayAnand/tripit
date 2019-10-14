package main

import (
	"fmt"
	"time"

	"github.com/AashrayAnand/tripit/session"
	"github.com/AashrayAnand/tripit/trip"
	"github.com/AashrayAnand/tripit/user"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	_, _ = session.Client.SetNX("x", "y", 60*time.Second).Result()
	val, err := session.Client.Get("x").Result()
	if err != nil {
	   fmt.Println("error", err.Error())
	   return
	}
	fmt.Println(val)

	router := gin.Default() // initialize gin routing engine

	user.AddUserRoutes(router) // add all routes in the /user groups
	router.POST("/trip", trip.Create)

	router.Use(cors.Default())
	//router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run(":3000") // listen and server 0.0.0.0:8080
}
