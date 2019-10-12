package main

import (
	"fmt"
	"time"

	"github.com/AashrayAnand/Bill-List/session"
	"github.com/AashrayAnand/Bill-List/user"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := session.Client.SetNX("key", "value", 10*time.Second).Result()
	val, err := session.Client.Get("key").Result()

	fmt.Println(val, err)
	router := gin.Default() // initialize gin routing engine

	user.AddUserRoutes(router) // add all routes in the /user groups

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run() // listen and server 0.0.0.0:8080
}
