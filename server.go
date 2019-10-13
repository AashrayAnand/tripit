package main

import (
	"fmt"
	"time"

	"github.com/AashrayAnand/tripit/session"
	"github.com/AashrayAnand/tripit/user"
	"github.com/gin-gonic/gin"
)

func main() {

	_, _ = session.Client.SetNX("x", "y", 60*time.Second).Result()
	val, _ := session.Client.Get("x").Result()
	fmt.Println(val)

	router := gin.Default() // initialize gin routing engine

	user.AddUserRoutes(router) // add all routes in the /user groups

	//router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run() // listen and server 0.0.0.0:8080
}
